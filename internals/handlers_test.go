package internals

import (
	"net/http"
	"testing"
)

func Test_application_handler1(t *testing.T) {
	type fields struct {
		auth struct {
			username string
			password string
		}
	}
	type args struct {
		rw http.ResponseWriter
		r  *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &application{
				auth: tt.fields.auth,
			}
			app.handler1(tt.args.rw, tt.args.r)
		})
	}
}
