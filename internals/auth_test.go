package internals

import (
	"net/http"
	"reflect"
	"testing"
)

func Test_application_basicAuth(t *testing.T) {
	type fields struct {
		auth struct {
			username string
			password string
		}
	}
	type args struct {
		next http.HandlerFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &application{
				auth: tt.fields.auth,
			}
			if got := app.basicAuth(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("application.basicAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}
