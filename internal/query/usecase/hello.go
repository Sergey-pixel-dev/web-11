package usecase

func (u *Usecase) FetchHelloMessage(name string) (string, error) {
	time, flag, err := u.p.GetTimeLastVisit(name)
	if err == nil {
		if flag {
			err1 := u.p.UpdateTimeLastVisit(name)
			return time, err1
		} else {
			err2 := u.p.SetTimeLastVisit(name)
			return "", err2
		}
	}
	return "", err
}
