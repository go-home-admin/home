package constraint

type ServiceProvider interface {
	register()
	boot()
}
