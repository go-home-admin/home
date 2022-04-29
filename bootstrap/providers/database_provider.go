package providers

// DatabaseProvider @Bean("database")
type DatabaseProvider struct{}

func (m *DatabaseProvider) GetBean(alias string) interface{} {
	switch alias {
	case "postgresql":
	case "mysql":
		return NewMysqlProvider()
	}
	return nil
}
