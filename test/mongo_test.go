package test

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
)

func init() {

}

func TestMongo(t *testing.T) {
	//client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	//if err != nil {
	//	log.Infof("mongo启动失败")
	//}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatalf("启动mongo失败")
		panic(err)
	}

	//err = client.Ping(ctx, readpref.Primary())
	//
	//if err != nil {
	//	log.Fatalf("ping mongo失败")
	//	panic(err)
	//}

	collection := client.Database("testing").Collection("numbers")
	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.1415926})
	if err != nil {
		panic(err)
	}
	id := res.InsertedID
	fmt.Println(id)

}

func BenchmarkMongoInsert(b *testing.B) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatalf("启动mongo失败")
		panic(err)
	}
	collection := client.Database("testing").Collection("numbers")
	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.1415926})
	if err != nil {
		panic(err)
	}
	id := res.InsertedID
	fmt.Println(id)
}
