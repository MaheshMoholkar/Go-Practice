package util

var data map[int]string = map[int]string{
	1: "bar",
}

type Storage interface {
	Get(id int) (any, error)
	Put(id int, val any) error
}

type FooStorage struct{}

func (s *FooStorage) Get(id int) (any, error) {
	return data[id], nil
}

func (s *FooStorage) Put(id int, val any) error {
	data[id] = val.(string)
	return nil
}

type Server struct {
	Store Storage
}
