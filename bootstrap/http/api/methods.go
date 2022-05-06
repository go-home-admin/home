package api

type Config = map[string]string

func Get(str string) *Config {
	config := make(Config)
	config["method"] = "get"
	config["url"] = str

	return &config
}

func Post(str string) *Config {
	config := make(Config)
	config["method"] = "post"
	config["url"] = str

	return &config
}

func Delete(str string) *Config {
	config := make(Config)
	config["method"] = "delete"
	config["url"] = str

	return &config
}

func Put(str string) *Config {
	config := make(Config)
	config["method"] = "put"
	config["url"] = str

	return &config
}

func Options(str string) *Config {
	config := make(Config)
	config["method"] = "options"
	config["url"] = str

	return &config
}
func Any(str string) *Config {
	config := make(Config)
	config["method"] = "any"
	config["url"] = str

	return &config
}
