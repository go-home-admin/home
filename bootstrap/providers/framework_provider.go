package providers

// @Bean
type FrameworkProvider struct {
	config   *ConfigProvider   `inject:""`
	database *DatabaseProvider `inject:""`
}
