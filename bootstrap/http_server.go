package bootstrap

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"dnsgo/app/api"
	"dnsgo/app/middleware"
	"dnsgo/app/model"
	"dnsgo/libs/mysql"
	"dnsgo/libs/redis"
)

func preHttpServer() {
	err := godotenv.Load("config/.env.http")
	if err != nil {
		panic(err)
	}

	mysql.InitWithDsn(os.Getenv("MYSQL_DSN"))
	redis.InstallWithURL(os.Getenv("REDIS_URL"))

	session := mysql.NewSession(nil)

	err = session.AutoMigrate(&model.Zone{}, &model.DNSRecord{}, &model.QueryLog{})
	if err != nil {
		panic(err)
	}
}

func HttpServer() {
	preHttpServer()
	fmt.Println("http server")

	r := mux.NewRouter()

	r.Use(middleware.Auth.AuthToken)

	r.HandleFunc("/zones", api.Zone.List).Methods("GET")
	r.HandleFunc("/zones", api.Zone.Create).Methods("POST")

	withZoneRouter := r.PathPrefix("/zones/{zone_id}").Subrouter()
	withZoneRouter.Use(middleware.Zone.GetZone)
	withZoneRouter.HandleFunc("", api.Zone.Get).Methods("GET")
	withZoneRouter.HandleFunc("", api.Zone.Update).Methods("PUT")
	withZoneRouter.HandleFunc("", api.Zone.Delete).Methods("DELETE")

	withZoneRouter.HandleFunc("/dns_records", api.DNSRecord.List).Methods("GET")
	withZoneRouter.HandleFunc("/dns_records", api.DNSRecord.Create).Methods("POST")

	withDNSRecordRouter := withZoneRouter.PathPrefix("/dns_records/{record_id}").Subrouter()
	withDNSRecordRouter.Use(middleware.DNSRecord.GetDNSRecord)
	withDNSRecordRouter.HandleFunc("", api.DNSRecord.Get).Methods("GET")
	withDNSRecordRouter.HandleFunc("", api.DNSRecord.Update).Methods("PUT")
	withDNSRecordRouter.HandleFunc("", api.DNSRecord.Delete).Methods("DELETE")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Fatal(srv.ListenAndServe())
}
