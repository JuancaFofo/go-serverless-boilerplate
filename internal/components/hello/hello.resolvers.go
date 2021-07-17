package hello

type HelloResolver struct {}

func (_ *HelloResolver) SayHello() string {
	return "Hello Golang Serverless!!!"
}
