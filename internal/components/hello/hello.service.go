package hello_comp

type HelloService struct {
	dao *HelloDao
}

func NewHelloService(dao *HelloDao) *HelloService {
	return &HelloService{
		dao,
	}
}

func (s *HelloService) SayHello() string {
	return s.dao.SayHello()
}
