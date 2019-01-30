package account

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/kalledk/go-ynab/ynab"
	"github.com/kalledk/go-ynab/ynab/endpoint"
	"github.com/kalledk/go-ynab/ynab/endpoint/mock_endpoint"
)

var (
	validID      = ID{uuid.MustParse("5d4a990e-193b-4971-9d8a-343607de8a8e")}
	validAccount = Account{
		validID,
		"Fælleskonto",
		"creditCard",
		true,
		50000,
		0,
		50000,
		false,
		false,
		"",
	}

	invalidID      = ID{uuid.MustParse("15360985-bd8f-4018-9354-a7916f30e31d")}
	invalidAccount = Account{}
)

func makeMockEndpoint(ctrl *gomock.Controller) endpoint.Getter {
	mockEndpoint := mock_endpoint.NewMockGetter(ctrl)

	validPath := fmt.Sprintf("accounts/%v", validID)
	var validError error
	validResponse := Response{Wrapper{validAccount}}
	mockEndpoint.EXPECT().Get(validPath, gomock.Any()).SetArg(1, validResponse).Return(validError).AnyTimes()

	invalidPath := fmt.Sprintf("accounts/%v", invalidID)
	invalidError := ynab.ErrorDetail{ID: "1", Name: "2", Detail: "3"}
	invalidResponse := Response{}
	mockEndpoint.EXPECT().Get(invalidPath, gomock.Any()).SetArg(1, invalidResponse).Return(invalidError).AnyTimes()
	return mockEndpoint
}

func TestGetAccount(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		id     ID
		getter endpoint.Getter
	}
	tests := []struct {
		name    string
		args    args
		wantAcc Account
		wantErr bool
	}{
		{
			"ValidID",
			args{validID, makeMockEndpoint(ctrl)},
			validAccount,
			false,
		},
		{
			"InValidID",
			args{invalidID, makeMockEndpoint(ctrl)},
			invalidAccount,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotAcc, err := GetAccount(tt.args.getter, tt.args.id)
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
