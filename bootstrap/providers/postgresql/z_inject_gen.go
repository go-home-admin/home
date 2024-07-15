// gen for home toolset
package postgresql

import (
	providers "github.com/go-home-admin/home/bootstrap/providers"
	services "github.com/go-home-admin/home/bootstrap/services"
)

var _PostgresqlProviderSingle *PostgresqlProvider

func GetAllProvider() []interface{} {
	return []interface{}{
		NewPostgresqlProvider(),
	}
}

func NewPostgresqlProvider() *PostgresqlProvider {
	if _PostgresqlProviderSingle == nil {
		_PostgresqlProviderSingle = &PostgresqlProvider{}
		_PostgresqlProviderSingle.config = providers.GetBean("config").(providers.Bean).GetBean("database").(*services.Config)
		providers.AfterProvider(_PostgresqlProviderSingle, "postgresql")
	}
	return _PostgresqlProviderSingle
}
