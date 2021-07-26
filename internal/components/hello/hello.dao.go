package hello

type HelloDao struct {}

func NewHelloDao() *HelloDao {
	return &HelloDao{}
}

func(d *HelloDao) SayHello() string {
	return "Hello from DAO"
}
