package mongo

import (
	"context"
	"github.com/go-home-admin/home/bootstrap/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// MongoProvider @Bean("mongo")
type MongoProvider struct {
	config *services.Config `inject:"config, database"`
	dbs    map[string]*mongo.Database
}

func (m *MongoProvider) Init() {
	m.dbs = make(map[string]*mongo.Database)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	connections := m.config.GetKey("connections")
	for name, dataT := range connections {
		data, ok := dataT.(map[interface{}]interface{})
		if !ok {
			continue
		}
		driver := data["driver"].(string)
		if driver != "mongo" {
			continue
		}
		config := services.NewConfig(data)
		host := config.GetString("host")
		port := config.GetString("port")
		database := config.GetString("database")
		var client *mongo.Client
		var err error
		if config.GetString("username") != "" {
			auth := options.Credential{
				AuthSource: config.GetString("auth_source"),
				Username:   config.GetString("username"),
				Password:   config.GetString("password"),
			}
			client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+host+":"+port).SetAuth(auth))
		} else {
			client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+host+":"+port))
		}
		if err != nil {
			panic("mongodb 连接格式错误")
		}

		err = client.Ping(ctx, nil)
		if err != nil {
			log.Println(err)
			panic("mongodb无法连接")
		}
		m.dbs[name.(string)] = client.Database(database)
	}
}

func (m *MongoProvider) GetBean(alias string) interface{} {
	return m.dbs[alias]
}
