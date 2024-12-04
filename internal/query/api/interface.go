package api

type Usecase interface {
	FetchHelloMessage(name string) (string, error)
}
