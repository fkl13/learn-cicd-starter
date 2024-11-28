package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		headerName string
		value      string
		want       string
		wantErr    string
	}{
		{
			headerName: "Authorization",
			value:      "ApiKey key",
			want:       "key",
			wantErr:    "no error",
		},
		{
			wantErr: "no authorization header included",
		},
		{
			headerName: "Authorization",
			value:      "No API key",
			wantErr:    "malformed authorization header",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Run test case #%v", i), func(t *testing.T) {
			headers := http.Header{}
			headers.Set(test.headerName, test.value)

			result, err := GetAPIKey(headers)
			if err != nil {
				if err.Error() == test.wantErr {
					return
				}
				t.Errorf("want %s, got: %s", test.wantErr, err.Error())
				return
			}

			if result != test.want {
				t.Errorf("Want: %s, got: %s", test.value, test.want)
				return
			}
		})
	}
}
