package bootstrap

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/miekg/dns"
	"github.com/oschwald/geoip2-golang"

	"dnsgo/app/cache"
	"dnsgo/app/model"
	"dnsgo/app/service"
	"dnsgo/libs/mysql"
	"dnsgo/libs/redis"
	"dnsgo/utils"
)

var (
	serverIP string
	geoIP    *geoip2.Reader
)

func getMyIP() (string, error) {
	resp, err := http.Get(os.Getenv("MY_IP_URL"))
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(ip), nil
}

func initGeoIP() {
	db, err := geoip2.Open("GeoLite2-Country.mmdb")
	if err != nil {
		panic(err)
	}
	geoIP = db
}

func preDNSServer() {
	err := godotenv.Load("config/.env.dns")
	if err != nil {
		panic(err)
	}

	serverIP, err = getMyIP()
	if err != nil {
		panic(err)
	}

	initGeoIP()

	mysql.InitWithDsn(os.Getenv("MYSQL_DSN"))
	redis.InstallWithURL(os.Getenv("REDIS_URL"))
}

func DNSServer() {
	preDNSServer()
	fmt.Println("dns server")

	dns.HandleFunc(".", func(w dns.ResponseWriter, r *dns.Msg) {
		start := time.Now()

		var msg = new(dns.Msg)
		msg.SetReply(r)
		msg.Authoritative = true

		var zoneID string

		question := r.Question[0]
		name := strings.ToLower(question.Name)
		zoneName := strings.TrimSuffix(name, ".")
		qType := question.Qtype
		recordType := dns.TypeToString[qType]

		remoteIp, _, _ := net.SplitHostPort(w.RemoteAddr().String())

		country, _ := geoIP.Country(net.ParseIP(remoteIp))
		var remoteCountry, remoteCountryCode string
		if country != nil {
			remoteCountry = country.Country.Names["en"]
			remoteCountryCode = country.Country.IsoCode
		}

		log.Println("country: ", remoteCountry, remoteCountryCode)

		defer func() {
			// record query log
			cost := time.Since(start).Milliseconds()
			// record request and response
			var content string
			if msg.Rcode == dns.RcodeSuccess {
				content = strings.Join(extractContentFromAnswer(msg.Answer), ",")
			}
			id, _ := utils.ID.Generate(21)
			queryLog := &model.QueryLog{
				ID:                id,
				ZoneID:            zoneID,
				QueryType:         recordType,
				QueryName:         name,
				RemoteIP:          remoteIp,
				RemoteCountry:     remoteCountry,
				RemoteCountryCode: remoteCountryCode,
				ServerIP:          serverIP,
				Rcode:             dns.RcodeToString[msg.Rcode],
				Content:           content,
				Cost:              cost,
				CreatedAt:         start,
			}
			go func() {
				_, err := service.QueryLog.Create(context.Background(), queryLog)
				if err != nil {
					log.Println(err)
				}
			}()
		}()

		if recordType == "ANY" {
			msg.SetRcode(r, dns.RcodeRefused)
			_ = w.WriteMsg(msg)
			return
		}

		ctx := context.Background()
		records, err := cache.DNSRecord.FindByZoneNameAndType(ctx, zoneName, recordType)
		if err != nil || len(records) == 0 {
			msg.SetRcode(r, dns.RcodeNameError)
			_ = w.WriteMsg(msg)
			return
		}

		for _, record := range records {
			zoneID = record.ZoneID
			switch record.Type {
			case "A":
				msg.Answer = append(msg.Answer, &dns.A{
					Hdr: dns.RR_Header{
						Name:   name,
						Rrtype: dns.TypeA,
						Class:  dns.ClassINET,
						Ttl:    uint32(record.TTL),
					},
					A: net.ParseIP(record.Content),
				})
			case "CNAME":
				msg.Answer = append(msg.Answer, &dns.CNAME{
					Hdr: dns.RR_Header{
						Name:   name,
						Rrtype: dns.TypeCNAME,
						Class:  dns.ClassINET,
						Ttl:    uint32(record.TTL),
					},
					Target: record.Content,
				})
			}
		}
		_ = w.WriteMsg(msg)
	})

	srv := &dns.Server{
		Addr: ":53",
		Net:  "udp",
	}

	log.Fatal(srv.ListenAndServe())
}

func extractContentFromAnswer(answer []dns.RR) []string {
	var contentList []string

	for _, rr := range answer {
		switch rr := rr.(type) {
		case *dns.A:
			contentList = append(contentList, rr.A.String())
		case *dns.CNAME:
			contentList = append(contentList, rr.Target)
		}
	}

	return contentList
}
