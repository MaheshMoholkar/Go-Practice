package util

import "fmt"

type NumberStorer interface {
	GetAll() ([]int, error)
	Put(int) error
}

type PostgresNumberStorer struct {
	// values
}

func (p PostgresNumberStorer) GetAll() ([]int, error) {
	return []int{4, 5, 6}, nil
}

func (p PostgresNumberStorer) Put(int) error {
	return nil
}

type MongoDBNumberStorer struct {
	// values
}

func (m MongoDBNumberStorer) GetAll() ([]int, error) {
	return []int{1, 2, 3}, nil
}

func (m MongoDBNumberStorer) Put(int) error {
	return nil
}

type ApiServer struct {
	numberStorer NumberStorer
}

func main() {
	var apiServer ApiServer = ApiServer{
		numberStorer: PostgresNumberStorer{},
	}

	if err := apiServer.numberStorer.Put(1); err != nil {
		panic(err)
	}

	numbers, err := apiServer.numberStorer.GetAll()

	if err != nil {
		panic(err)
	}
	fmt.Println(numbers)
}
