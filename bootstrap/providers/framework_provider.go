package providers

// FrameworkProvider @Bean
type FrameworkProvider struct {
	config   *ConfigProvider   `inject:""`
	database *DatabaseProvider `inject:""`
}
