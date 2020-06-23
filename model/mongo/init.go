package mongo

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

import (
	"fmt"
)

type Database struct {
	Mongo *mongo.Client
}

func (d *Database) Init() {
	DB = &Database{
		Mongo: getSelfDb(),
	}
}

var DB *Database

func getSelfDb() *mongo.Client {
	return InitSelfDB()
}

func InitSelfDB() *mongo.Client {
	return openDB(viper.GetString("mongo.ip"),
		viper.GetString("mongo.port"))
}

// 连接设置

func openDB(ip string, port string) *mongo.Client {

	uri := fmt.Sprintf("mongodb://%s:%s", ip, port)
	//uri := "mongodb://127.0.0.1:27017"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetMaxPoolSize(20)) // 连接池
	if err != nil {
		fmt.Println(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic("mongo init err ")
	}

	setupDb(client)
	return client
}

func setupDb(client *mongo.Client) {
	//client.
}

func CloseCursor(cursor *mongo.Cursor) {
	err := cursor.Close(context.Background())
	if err != nil {
		log.Fatalf("%s mongo count查询失败", err)
	}
}
