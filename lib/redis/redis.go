package redis

import ."github.com/go-redis/redis"
import "fmt"


type RedisStore struct {
	client *Client
}

func GetRedisClient()(redisStore RedisStore) {
	options := Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
		}
	clientPtr := NewClient(&options)
	// fmt.Println("returning redis client")
	return RedisStore { clientPtr}
}

func (rs *RedisStore) AddPair(key string, value string) bool {
	err := rs.client.Set(key, value, 0).Err()
	if (err != nil ) {
		fmt.Println(err)
	}
	return err == nil
}


func (rs *RedisStore) GetValue (key string) string {
	value, err := rs.client.Get(key).Result()
	if (err != nil) {
		return value
	} else {
		return "Error"
	}
}


func (rs *RedisStore) RemoveKey(key string) bool{
	existed := rs.client.Del(key).String()
	return existed == ""
}