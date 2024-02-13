package usecase

type UserCase struct {
}

func NewUserCase() *UserCase {
	return &UserCase{}
}

func (usecase *UserCase) SendToGraph() error {
	return nil
}
