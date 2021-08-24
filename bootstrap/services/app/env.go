package app

type Environment string

var (
	EnvironmentNow        Environment = ""
	EnvironmentLocal      Environment = "local"
	EnvironmentDev        Environment = "dev"
	EnvironmentTesting    Environment = "testing"
	EnvironmentStaging    Environment = "staging"
	EnvironmentProduction Environment = "production"
)

func SetEnvironment(env string) {
	EnvironmentNow = Environment(env)
}

func GetEnvironment() Environment {
	return EnvironmentNow
}

func EnvIsLocal() bool {
	return EnvironmentNow == EnvironmentLocal
}

func EnvIsDev() bool {
	return EnvironmentNow == EnvironmentDev
}

func EnvIsTesting() bool {
	return EnvironmentNow == EnvironmentTesting
}

func EnvIsStaging() bool {
	return EnvironmentNow == EnvironmentStaging
}

func EnvIsProduction() bool {
	return EnvironmentNow == EnvironmentProduction
}
