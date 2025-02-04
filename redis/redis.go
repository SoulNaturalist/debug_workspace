package redis

import (
	"fmt"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func GetValue(key string) (string, error) {
	val, err := rdb.Get(key).Result()
	switch {
	case err == redis.Nil:
		fmt.Println("key does not exist")
	case err != nil:
		fmt.Println("Get failed", err)
	case val == "":
		fmt.Println("value is empty")
	}
	return val, err
}

func SetValue(key string, value string) error {
	err := rdb.Set(key, value, 0).Err()
	return err
}

func main() {
	InitRedis()

	value, err := GetValue("test")
	if err == nil {
		fmt.Printf("Value: %s\n", value)
	}
}
