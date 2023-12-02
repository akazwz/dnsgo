package cache

import (
	"context"
	"time"

	"dnsgo/app/dao"
	"dnsgo/app/model"
	"dnsgo/libs/mysql"
	"dnsgo/libs/redis"
)

var Zone = &zoneCache{}

type zoneCache struct{}

func (c *zoneCache) Get(ctx context.Context, id string) (*model.Zone, error) {
	client := redis.GetClient()
	data, err := client.Get(ctx, "zone:"+id).Result()
	if err == nil {
		zone := &model.Zone{}
		err = zone.FromJSON(data)
		if err == nil {
			return zone, nil
		}
	}
	session := mysql.NewSession(ctx)
	zone, err := dao.Zone.Get(session, id)
	if err != nil {
		return nil, err
	}
	jsonStr, err := zone.ToJSON()
	if err == nil {
		client.Set(ctx, "zone:"+id, jsonStr, time.Minute*60)
	}
	return zone, nil
}
