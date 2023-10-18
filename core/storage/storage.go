package storage

import (
	"fmt"
	"os"

	"github.com/j-04/pass-manager/core/crypto"
	"github.com/j-04/pass-manager/core/format"
	"github.com/j-04/pass-manager/core/model"
)

type Storage interface {
	AddCredential(credential *model.Credential) error
	GetCredentials() []*model.Credential
	GetCredential(id string) (*model.Credential, error)
	UpdateCredential(credential *model.Credential) error
	DeleteCredential(credential *model.Credential) error
}

/*
	There is no need for a smmodelproject to keep all implementations
	in dedicated files, because only one person is working on this project
	and in perspective there will be no other implementations except In-memory
	and File.
*/

type InMemoryStorage struct {
	storage []*model.Credential
	encoder crypto.Encoder
}

func NewInMemoryStorage(
	encoder crypto.Encoder,
) *InMemoryStorage {
	c := generateTestCredentials()
	storage := &InMemoryStorage{
		storage: c,
	}
	return storage
}

func generateTestCredentials() []*model.Credential {
	c := make([]*model.Credential, 0, 10)
	c = append(c,
		&model.Credential{
			Id:        "1",
			Name:      "Test 1",
			Login:     "testlogin",
			Email:     "test1@test.com",
			Password:  "test1",
			IsEncoded: false,
		},
		&model.Credential{
			Id:        "2",
			Name:      "Test 2",
			Login:     "testlogin2",
			Email:     "test2@test.com",
			Password:  "test2",
			IsEncoded: false,
		},
		&model.Credential{
			Id:        "3",
			Name:      "Test 3",
			Login:     "testlogin3",
			Email:     "test3@test.com",
			Password:  "test3",
			IsEncoded: false,
		},
		&model.Credential{
			Id:        "4",
			Name:      "Test 4",
			Login:     "testlogin4",
			Email:     "test4@test.com",
			Password:  "test4",
			IsEncoded: false,
		},
		&model.Credential{
			Id:        "5",
			Name:      "Test 5",
			Login:     "testlogin5",
			Email:     "test5@test.com",
			Password:  "test5",
			IsEncoded: false,
		},
	)
	return c
}

func (this *InMemoryStorage) AddCredential(credential *model.Credential) error {
	this.storage = append(this.storage, credential)
	return nil
}

func (this *InMemoryStorage) GetCredentials() []*model.Credential {
	return this.storage
}

func (this *InMemoryStorage) GetCredential(id string) (*model.Credential, error) {
	for _, c := range this.storage {
		if c.Id == id {
			return c, nil
		}
	}
	return nil, fmt.Errorf("Couldn't find a credential by id: %s", id)
}

func (this *InMemoryStorage) UpdateCredential(credential *model.Credential) error {
	for _, c := range this.storage {
		if c.Id == credential.Id {
			c.Email = credential.Email
			c.Login = credential.Login
			c.Name = credential.Name
			c.Password = credential.Password
			c.IsEncoded = credential.IsEncoded
		}
	}
	return nil
}

func (this *InMemoryStorage) DeleteCredential(credential *model.Credential) error {
	for i, c := range this.storage {
		if c.Id == credential.Id {
			this.storage[i] = nil
			break
		}
	}
	return nil
}

type FileStorage struct {
	file    *os.File
	storage map[string]*model.Credential
	encoder crypto.Encoder
	format  format.Serializer
}

func NewFileStorage(file *os.File, encoder crypto.Encoder, format format.Serializer) *FileStorage {
	return &FileStorage{
		file:    file,
		storage: map[string]*model.Credential{},
		encoder: encoder,
		format:  format,
	}
}

func (this *FileStorage) AddCredential(credential *model.Credential) error {
	this.storage[credential.Id] = credential
	// TODO Save updated data to the file
	return nil
}

func (this *FileStorage) GetCredentials() []*model.Credential {
	var data []*model.Credential = make([]*model.Credential, 0, len(this.storage))
	for _, v := range this.storage {
		data = append(data, v)
	}
	return data
}

func (this *FileStorage) GetCredential(id string) (*model.Credential, error) {
	c, ok := this.storage[id]
	if !ok {
		return nil, fmt.Errorf("Couldn't find credentila by id %s", id)
	}
	return c, nil
}

func (this *FileStorage) UpdateCredential(credential *model.Credential) error {
	// TODO Save updated data to the file
	return nil
}

func (this *FileStorage) DeleteCredential(credential *model.Credential) error {
	delete(this.storage, credential.Id)
	// TODO Save updated data to the file
	return nil
}
