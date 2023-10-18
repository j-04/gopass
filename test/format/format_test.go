package format

import (
	"reflect"
	"testing"

	"github.com/j-04/pass-manager/core/format"
	"github.com/j-04/pass-manager/core/model"
)

var (
	cred1 *model.Credential = model.NewCredential("1", "Test1", "testlogin1", "test@test.com", "test1", false)
	cred2 *model.Credential = model.NewCredential("2", "Test2", "testlogin2", "test@test.com", "test2", false)
)

const exp1 string = `{"1":{"id":"1","name":"Test1","login":"testlogin1","email":"test@test.com","password":"test1","is_encoded":false}}`

const exp2 string = `{"1":{"id":"1","name":"Test1","login":"testlogin1","email":"test@test.com","password":"test1","is_encoded":false},"2":{"id":"2","name":"Test2","login":"testlogin2","email":"test@test.com","password":"test2","is_encoded":false}}`

func NewMapWithData(data ...*model.Credential) map[string]*model.Credential {
	m := make(map[string]*model.Credential)
	for _, d := range data {
		m[d.Id] = d
	}
	return m
}

func TestJsonFormat_Marshal(t *testing.T) {
	d1 := NewMapWithData(
		cred1,
	)
	d2 := NewMapWithData(
		cred1,
		cred2,
	)

	type args struct {
		data map[string]*model.Credential
	}
	tests := []struct {
		name    string
		this    *format.JsonSerializer
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Empty map",
			args: args{data: make(map[string]*model.Credential)},
			want: []byte("{}"),
		},
		{
			name: "One element in the map",
			args: args{data: d1},
			want: []byte(exp1),
		},
		{
			name: "Two elements in the map",
			args: args{data: d2},
			want: []byte(exp2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &format.JsonSerializer{}
			got, err := this.Marshal(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonFormat.Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonFormat.Marshal() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func TestJsonFormat_Unmarshal(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		this    *format.JsonSerializer
		args    args
		want    map[string]*model.Credential
		wantErr bool
	}{
		{
			name: "Empty json",
			args: args{data: []byte("{}")},
			want: make(map[string]*model.Credential),
		},
		{
			name: "A map with one element in the json",
			args: args{data: []byte(exp1)},
			want: NewMapWithData(cred1),
		},
		{
			name: "A map with two elements in the json",
			args: args{data: []byte(exp2)},
			want: NewMapWithData(cred1, cred2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &format.JsonSerializer{}
			got, err := this.Unmarshal(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonFormat.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonFormat.Unmarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}
