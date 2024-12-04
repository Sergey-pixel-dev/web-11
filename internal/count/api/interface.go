package api

type Usecase interface {
	FetchHelloMessage(name string) error
	GetHelloMessage() (string, error)
}
