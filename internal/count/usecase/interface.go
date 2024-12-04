package usecase

type Provider interface {
	UpdateCount(count string) error
	GetCount() (string, error)
}
