package hello_comp

type HelloResolver struct {
	service *HelloService
}

func NewHelloResolver(service *HelloService) *HelloResolver {
	return &HelloResolver{
		service,
	}
}

func (r *HelloResolver) SayHello() string {
	return r.service.SayHello()
}
