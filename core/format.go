package core

import (
	"encoding/json"
	"fmt"
)

type Serializer interface {
	Marshal(data map[string]*Credential) ([]byte, error)
	Unmarshal(data []byte) (map[string]*Credential, error)
}

type CsvSerializer struct{}

func NewCsvSerializer() *CsvSerializer {
	return &CsvSerializer{}
}

func (serializer *CsvSerializer) Marshal(data map[string]*Credential) ([]byte, error) {
	return nil, nil
}

func (serializer *CsvSerializer) Unmarshal(data []byte) (map[string]*Credential, error) {
	return nil, nil
}

type JsonSerializer struct{}

func NewJsonSerializer() *JsonSerializer {
	return &JsonSerializer{}
}

func (serializer *JsonSerializer) Marshal(data map[string]*Credential) ([]byte, error) {
	m, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("couldn't marshal the data %v. Err: %w", data, err)
	}
	return m, nil
}

func (serializer *JsonSerializer) Unmarshal(data []byte) (map[string]*Credential, error) {
	m := make(map[string]*Credential)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return nil, fmt.Errorf("couldn't unmarshal data %v. Err: %w", string(data), err)
	}
	return m, nil
}
