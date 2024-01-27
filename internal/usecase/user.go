package usecase

type UserUsecase struct {
	user User
}

func NewUserUsecase() UserUsecase {
	return UserUsecase{}
}

func (u UserUsecase) Name() string {
	return u.user.Name()
}

func (u UserUsecase) ID() int64 {
	return u.user.ID()
}
