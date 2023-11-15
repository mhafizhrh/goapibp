package helper

type JWT interface {
	CreateToken() string
}

type jwt struct {
}

func NewJWT() JWT {
	return &jwt{}
}

func (helper jwt) CreateToken() string {
	return "asdasdasd"
}
