package endpoint

import (
	"path"
	"testing"
)

type Base struct {
	Path string
}

func (e *Base) Get(subpath string, responseModel interface{}) (err error) {
	*responseModel.(*string) = path.Join(e.Path, subpath)
	return nil
}

func TestEndpoint_Get(t *testing.T) {
	type args struct {
		subpath       string
		responseModel string
	}
	tests := []struct {
		name    string
		e       *Endpoint
		args    args
		want    string
		wantErr bool
	}{
		{"Simple", &Endpoint{&Base{"/"}, "sub"}, args{"flaf", "delete"}, "/sub/flaf", false},
		{"RelRoot", &Endpoint{&Base{"/"}, "sub"}, args{"", "delete"}, "/sub", false},
		{"AbRoot", &Endpoint{&Base{"/"}, "sub"}, args{"/", "delete"}, "/sub", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.Get(tt.args.subpath, &tt.args.responseModel); (err != nil) != tt.wantErr {
				t.Errorf("Endpoint.Get() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.args.responseModel != tt.want {
				t.Errorf("Endpoint.Get() get = %v, wantErr %v", tt.args.responseModel, tt.want)
			}
		})
	}
}
