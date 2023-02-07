package services

type LoginService interface {
	LoginUser(email string, password string) bool
}

type loginInformation struct {
	email string
	password string
}

func (li *loginInformation) LoginUser(email string, password string) bool {
		return li.email == email && li.password == password
}

func StaticLoginService() LoginService {
	return &loginInformation{
		email: "email@email.com",
		password: "12345678",
	}
}