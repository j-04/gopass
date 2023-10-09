package core

type Credential struct {
	Id        string
	Name      string
	Login     string
	Email     string
	Password  string
	IsEncoded bool
}

func NewCredential(name, login, email, password string, isEncoded bool) *Credential {
	credential := &Credential{
		Name:      name,
		Login:     login,
		Email:     email,
		Password:  password,
		IsEncoded: isEncoded,
	}
	return credential
}
