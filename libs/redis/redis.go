package redis

import "github.com/redis/go-redis/v9"

var kv *redis.Client

func InstallWithURL(url string) {
	opt, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}
	kv = redis.NewClient(opt)
}

func GetClient() *redis.Client {
	return kv
}
