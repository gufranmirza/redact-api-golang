package redact

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/gufranmirza/redact-api-golang/src/config"
	"github.com/gufranmirza/redact-api-golang/src/models"
)

func TestNewRedact(t *testing.T) {
	tests := []struct {
		name string
		want Redact
	}{
		{
			name: "Happy Path",
			want: &redact{
				logger: log.New(os.Stdout, "redact :=> ", log.LstdFlags),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRedact(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRedact() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redact_RedactJSON(t *testing.T) {
	config.Config = &models.AppConfig{}
	redact := NewRedact()

	// Invalid Request Body
	req, err := http.NewRequest("POST", "/redact/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := redact.RedactJSON()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Valid Request Body
	var buffer = []byte(`{ "json_to_redact": { "a": { "b": { "c": "va12345l", "d": "val", "e": "val" }, "l": [ {"k": "a"}, {"k": { "p": "va12345l" }}, {"k": "c"} ] } }, "redact_regexes": [ { "path": "a.b.c", "regexes": [ "[0-9]{5}" ] }, { "path": "a.l[1].k.p", "regexes": [ "[0-9]{5}" ] } ], "redact_completely": [ "a.l[0].k" ] }`)
	req1, err := http.NewRequest("POST", "/redact/", bytes.NewBuffer(buffer))
	if err != nil {
		t.Fatal(err)
	}
	req1.Header.Set("Content-Type", "application/json")
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req1)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"a":{"b":{"c":"va*****l","d":"val","e":"val"},"l":[{"k":"*"},{"k":{"p":"va*****l"}},{"k":"c"}]}}`
	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}
