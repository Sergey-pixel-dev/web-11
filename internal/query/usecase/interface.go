package usecase

type Provider interface {
	GetTimeLastVisit(name string) (string, bool, error)
	UpdateTimeLastVisit(name string) error
	SetTimeLastVisit(name string) error
}
