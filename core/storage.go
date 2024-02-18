package core

import (
	"fmt"
	"os"
)

type Storage interface {
	AddCredential(credential *Credential) error
	GetCredentials() []*Credential
	GetCredential(id string) (*Credential, error)
	UpdateCredential(credential *Credential) error
	DeleteCredential(credential *Credential) error
}

/*
	There is no need for a smmodelproject to keep all implementations
	in dedicated files, because only one person is working on storage project
	and in perspective there will be no other implementations except In-memory
	and File.
*/

type InMemoryStorage struct {
	storage []*Credential
	encoder Encoder
}

func NewInMemoryStorage(
	encoder Encoder,
) *InMemoryStorage {
	c := generateTestCredentials()
	storage := &InMemoryStorage{
		storage: c,
	}
	return storage
}

func generateTestCredentials() []*Credential {
	c := make([]*Credential, 0, 10)
	c = append(c,
		&Credential{
			Id:        "1",
			Name:      "Test 1",
			Login:     "testlogin",
			Email:     "test1@test.com",
			Password:  "test1",
			IsEncoded: false,
		},
		&Credential{
			Id:        "2",
			Name:      "Test 2",
			Login:     "testlogin2",
			Email:     "test2@test.com",
			Password:  "test2",
			IsEncoded: false,
		},
		&Credential{
			Id:        "3",
			Name:      "Test 3",
			Login:     "testlogin3",
			Email:     "test3@test.com",
			Password:  "test3",
			IsEncoded: false,
		},
		&Credential{
			Id:        "4",
			Name:      "Test 4",
			Login:     "testlogin4",
			Email:     "test4@test.com",
			Password:  "test4",
			IsEncoded: false,
		},
		&Credential{
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

func (storage *InMemoryStorage) AddCredential(credential *Credential) error {
	storage.storage = append(storage.storage, credential)
	return nil
}

func (storage *InMemoryStorage) GetCredentials() []*Credential {
	return storage.storage
}

func (storage *InMemoryStorage) GetCredential(id string) (*Credential, error) {
	for _, c := range storage.storage {
		if c.Id == id {
			return c, nil
		}
	}
	return nil, fmt.Errorf("couldn't find a credential by id: %s", id)
}

func (storage *InMemoryStorage) UpdateCredential(credential *Credential) error {
	for _, c := range storage.storage {
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

func (storage *InMemoryStorage) DeleteCredential(credential *Credential) error {
	for i, c := range storage.storage {
		if c.Id == credential.Id {
			storage.storage[i] = nil
			break
		}
	}
	return nil
}

type FileStorage struct {
	file    *os.File
	storage map[string]*Credential
	encoder Encoder
	format  Serializer
}

func NewFileStorage(file *os.File, encoder Encoder, format Serializer) *FileStorage {
	return &FileStorage{
		file:    file,
		storage: map[string]*Credential{},
		encoder: encoder,
		format:  format,
	}
}

func (storage *FileStorage) AddCredential(credential *Credential) error {
	storage.storage[credential.Id] = credential
	// TODO Save updated data to the file
	return nil
}

func (storage *FileStorage) GetCredentials() []*Credential {
	var data []*Credential = make([]*Credential, 0, len(storage.storage))
	for _, v := range storage.storage {
		data = append(data, v)
	}
	return data
}

func (storage *FileStorage) GetCredential(id string) (*Credential, error) {
	c, ok := storage.storage[id]
	if !ok {
		return nil, fmt.Errorf("couldn't find credentila by id %s", id)
	}
	return c, nil
}

func (storage *FileStorage) UpdateCredential(credential *Credential) error {
	// TODO Save updated data to the file
	return nil
}

func (storage *FileStorage) DeleteCredential(credential *Credential) error {
	delete(storage.storage, credential.Id)
	// TODO Save updated data to the file
	return nil
}
