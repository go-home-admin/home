// gen for home toolset
package mongo

import (
	providers "github.com/go-home-admin/home/bootstrap/providers"
	services "github.com/go-home-admin/home/bootstrap/services"
)

var _MongoProviderSingle *MongoProvider

func GetAllProvider() []interface{} {
	return []interface{}{
		NewMongoProvider(),
	}
}

func NewMongoProvider() *MongoProvider {
	if _MongoProviderSingle == nil {
		_MongoProviderSingle = &MongoProvider{}
		_MongoProviderSingle.config = providers.GetBean("config").(providers.Bean).GetBean("database").(*services.Config)
		providers.AfterProvider(_MongoProviderSingle, "mongo")
	}
	return _MongoProviderSingle
}
