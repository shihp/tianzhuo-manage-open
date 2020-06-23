package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

type Database struct {
	Redis *redis.Client
}

func (d *Database) Init() {
	DB = &Database{
		Redis: getSelfDb(),
	}
}

var DB *Database

func getSelfDb() *redis.Client {
	return InitSelfDB()
}

func InitSelfDB() *redis.Client {
	return openDB(viper.GetString("redis.ip"),
		viper.GetString("redis.port"),
		viper.GetInt("redis.db"),
		viper.GetString("redis.password"))
}

// 连接设置

func openDB(ip string, port string, db int, password string) *redis.Client {
	uri := fmt.Sprintf("%s:%s", ip, port)
	//"localhost:6379",
	client := redis.NewClient(&redis.Options{
		Addr:     uri,
		Password: password, // no password set
		DB:       db,       // use default DB
	})
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	pong, err := client.Ping(ctx).Result()

	if err != nil {
		log.Fatal(err)
	}
	log.Println(pong)

	setupDb(client)
	return client
}

func LuaUnLockEval(ctx context.Context, code []string, val []string) interface{} {
	luaUnlockScript := "if redis.call('get', KEYS[1]) == ARGV[1] then return redis.call('del',KEYS[1]) else return 0 end"
	rs := DB.Redis.Eval(ctx, luaUnlockScript, code, val)
	if rs.Val().(int64) != 1 {
		errStr, err := jsoniter.MarshalToString(rs)
		if err != nil {
			log.Errorf("%s 解锁异常 ! 错误格式化异常 !", code[0])
		}
		log.Errorf("%s 解锁异常 ! , %s", code[0], errStr)
	}
	log.Infof("%s 解锁success ! ", code[0])
	return rs.Val()
}

func setupDb(client *redis.Client) {
}
