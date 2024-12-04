package usecase

func (u *Usecase) FetchHelloMessage(name string) error {
	err := u.p.UpdateCount(name)
	return err
}
func (u *Usecase) GetHelloMessage() (string, error) {
	count, err := u.p.GetCount()
	return count, err
}
