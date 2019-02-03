package user

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kalledk/go-ynab/ynab/api"
	"github.com/kalledk/go-ynab/ynab/endpoint"
	"github.com/kalledk/go-ynab/ynab/endpoint/mock_endpoint"
)

type EndpointMocker struct {
	t    *testing.T
	ctrl *gomock.Controller
}

func (e *EndpointMocker) Finish() {
	e.ctrl.Finish()
}

func (e *EndpointMocker) Make(path string) endpoint.API {
	mockEndpoint := mock_endpoint.NewMockAPI(e.ctrl)
	mockEndpoint.EXPECT().Get(gomock.Any()).DoAndReturn(
		func(model interface{}) error {
			return LoadReply(e.t, path, model)
		})
	return mockEndpoint
}

func SprintJSON(model interface{}) string {
	data, err := json.MarshalIndent(model, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func LoadJSON(t *testing.T, name string, model interface{}) {
	path := filepath.Join("testdata", name) // relative path

	rawbytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	decoder := json.NewDecoder(bytes.NewBuffer(rawbytes))
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(model); err != nil {
		t.Fatal(err)
	}

}

func LoadUser(t *testing.T, name string) (user User) {
	LoadJSON(t, name, &user)
	return user
}

func LoadReply(t *testing.T, name string, model interface{}) error {

	reply := struct {
		Response    interface{}
		ErrorDetail *api.ErrorDetail
	}{
		Response:    model,
		ErrorDetail: nil,
	}

	LoadJSON(t, name, &reply)

	if reply.ErrorDetail != nil {
		return *reply.ErrorDetail
	}

	return nil
}

func NewMocker(t *testing.T) *EndpointMocker {
	return &EndpointMocker{
		t,
		gomock.NewController(t),
	}
}

func TestGet(t *testing.T) {

	mocker := NewMocker(t)
	defer mocker.Finish()

	tests := []struct {
		name     string
		mock     endpoint.API
		wantUser User
		wantErr  bool
	}{
		{"Simple", mocker.Make("user.json"), LoadUser(t, "user.golden"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotUser, err := Get(tt.mock)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("Get() = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}
