package account

/*
import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"testing"

	"io/ioutil"
	filepath "path"

	"github.com/golang/mock/gomock"
	"github.com/kalledk/go-ynab/ynab/api"
	"github.com/kalledk/go-ynab/ynab/endpoint"
	"github.com/kalledk/go-ynab/ynab/endpoint/mock_endpoint"
)

type EndpointReply struct {
	Response Response
	Error    *api.ErrorDetail
}

type EndpointListReply struct {
	Response ResponseList
	Error    *api.ErrorDetail
}

func loadAccount(t *testing.T, name string) (account Account) {
	path := filepath.Join("testdata", name) // relative path
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	if err := json.Unmarshal(bytes, &account); err != nil {
		t.Fatal(err)
	}

	return
}

func loadGolden(t *testing.T, name string, model interface{}) {
	path := filepath.Join("testdata", name)

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	if err := json.Unmarshal(bytes, model); err != nil {
		t.Fatal(err)
	}
}

func loadAccountList(t *testing.T, name string) (account AccountList) {
	path := filepath.Join("testdata", name) // relative path
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	if err := json.Unmarshal(bytes, &account); err != nil {
		t.Fatal(err)
	}

	return
}

type EndpointMocker struct {
	t    *testing.T
	ctrl *gomock.Controller
}

func (e *EndpointMocker) Finish() {
	e.ctrl.Finish()
}

func (e *EndpointMocker) LoadMockResponse(name string, model interface{}) {
	path := filepath.Join("testdata", name) // relative path
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		e.t.Fatal(err)
	}

	if err := json.Unmarshal(bytes, model); err != nil {
		e.t.Fatal(err)
	}

	return
}

func (e *EndpointMocker) Make(resp interface{}, errorDetail *api.ErrorDetail) endpoint.Getter {

	var err error
	if errorDetail != nil {
		err = *errorDetail
	}

	mockEndpoint := mock_endpoint.NewMockGetter(e.ctrl)
	mockEndpoint.EXPECT().Get(gomock.Any()).SetArg(0, resp).Return(err)
	return mockEndpoint
}

func MakeAccount(e *EndpointMocker, path string) endpoint.Getter {
	var args EndpointReply
	e.LoadMockResponse(path, &args)
	return e.Make(args.Response, args.Error)
}

func MakeAccountList(e *EndpointMocker, path string) endpoint.Getter {
	var args EndpointListReply
	e.LoadMockResponse(path, &args)
	return e.Make(args.Response, args.Error)
}

func NewMocker(t *testing.T) *EndpointMocker {
	return &EndpointMocker{
		t,
		gomock.NewController(t),
	}
}

func TestGet(t *testing.T) {

	ctrl := NewMocker(t)
	defer ctrl.Finish()

	tests := []struct {
		name    string
		mock    endpoint.Getter
		wantAcc string
		wantErr bool
	}{
		{
			name:    "ValidID",
			mock:    MakeAccount(ctrl, "account.json"),
			wantAcc: "account.golden",
			wantErr: false,
		},
		{
			name:    "InValidID",
			mock:    MakeAccount(ctrl, "invalidToken.json"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotAcc, err := Get(tt.mock)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var wantAcc Account

			if tt.wantAcc != "" {
				loadGolden(t, tt.wantAcc, &wantAcc)
			}

			if !reflect.DeepEqual(gotAcc, wantAcc) {
				t.Errorf("GetAccount() = %v, want %v", gotAcc, wantAcc)
			}

			json, err := json.MarshalIndent(gotAcc, "", "  ")
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(string(json))
		})
	}
}

func TestGetList(t *testing.T) {

	ctrl := NewMocker(t)
	defer ctrl.Finish()

	tests := []struct {
		name    string
		mock    endpoint.Getter
		wantAcc AccountList
		wantErr bool
	}{
		{
			name:    "ValidID",
			mock:    MakeAccountList(ctrl, "accounts.json"),
			wantAcc: loadAccountList(t, "accounts.golden"),
			wantErr: false,
		},
		{
			name:    "InValidID",
			mock:    MakeAccountList(ctrl, "invalidToken.json"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotAcc, err := GetList(tt.mock)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAcc, tt.wantAcc) {
				t.Errorf("GetAccount() = %v, want %v", gotAcc, tt.wantAcc)
			}
		})
	}
}
*/
