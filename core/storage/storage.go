package storage

import (
	"fmt"
	"os"

	"github.com/j-04/pass-manager/core"
	"github.com/j-04/pass-manager/core/crypto"
	"github.com/j-04/pass-manager/core/format"
)

type Storage interface {
	AddCredential(credential *core.Credential) error
	GetCredentials() []*core.Credential
	GetCredential(id string) (*core.Credential, error)
	UpdateCredential(credential *core.Credential) error
	DeleteCredential(credential *core.Credential) error
}

/*
	There is no need for a small project to keep all implementations
	in dedicated files, because only one person is working on this project
	and in perspective there will be no other implementations except In-memory
	and File.
*/

type InMemoryStorage struct {
	storage []*core.Credential
	encoder crypto.Encoder
}

func NewInMemoryStorage(
	encoder crypto.Encoder,
) *InMemoryStorage {
	storage := &InMemoryStorage{
		storage: make([]*core.Credential, 0, 10),
	}
	return storage
}

func (this *InMemoryStorage) AddCredential(credential *core.Credential) error {
	this.storage = append(this.storage, credential)
	return nil
}

func (this *InMemoryStorage) GetCredentials() []*core.Credential {
	return this.storage
}

func (this *InMemoryStorage) GetCredential(id string) (*core.Credential, error) {
	for _, c := range this.storage {
		if c.Id == id {
			return c, nil
		}
	}
	return nil, fmt.Errorf("Couldn't find a credential by id: %s", id)
}

func (this *InMemoryStorage) UpdateCredential(credential *core.Credential) error {
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

func (this *InMemoryStorage) DeleteCredential(credential *core.Credential) error {
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
	storage map[string]*core.Credential
	encoder crypto.Encoder
	format  format.Format
}

func NewFileStorage(file *os.File, encoder crypto.Encoder, format format.Format) *FileStorage {
	return &FileStorage{
		file:    file,
		storage: map[string]*core.Credential{},
		encoder: encoder,
		format:  format,
	}
}

func (this *FileStorage) AddCredential(credential *core.Credential) error {
	return nil
}

func (this *FileStorage) GetCredentials() []*core.Credential {
	return nil
}

func (this *FileStorage) UpdateCredential(credential *core.Credential) error {
	return nil
}

func (this *FileStorage) DeleteCredential(credential *core.Credential) error {
	return nil
}
