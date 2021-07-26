package hello

type HelloResolver struct{
	service *HelloService
}

func NewHelloResolver(service *HelloService) *HelloResolver {
	return &HelloResolver{
		service,
	}
}

func (r *HelloResolver) SayHello() string {
	// container.DBOrm.Create(&todo.TodoItem{Description: "test1", Completed: false})
	return r.service.SayHello()
}
