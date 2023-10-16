package core

type Credential struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Login     string `json:"login"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsEncoded bool   `json:"is_encoded"`
}

func NewCredential(id, name, login, email, password string, isEncoded bool) *Credential {
	credential := &Credential{
		Id:        id,
		Name:      name,
		Login:     login,
		Email:     email,
		Password:  password,
		IsEncoded: isEncoded,
	}
	return credential
}
