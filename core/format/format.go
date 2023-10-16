package format

import (
	"encoding/json"
	"fmt"

	"github.com/j-04/pass-manager/core"
)

type Serializer interface {
	Marshal(data map[string]*core.Credential) ([]byte, error)
	Unmarshal(data []byte) (map[string]*core.Credential, error)
}

type JsonSerializer struct{}

func (this *JsonSerializer) Marshal(data map[string]*core.Credential) ([]byte, error) {
	m, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("Couldn't marshal the data %v. Err: %w", data, err)
	}
	return m, nil
}

func (this *JsonSerializer) Unmarshal(data []byte) (map[string]*core.Credential, error) {
	m := make(map[string]*core.Credential)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return nil, fmt.Errorf("Couldn't unmarshal data %v. Err: %w", string(data), err)
	}
	return m, nil
}
