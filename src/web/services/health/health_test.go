package health

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/gufranmirza/redact-api-golang/src/config"
	"github.com/gufranmirza/redact-api-golang/src/models"
)

func TestNewHealth(t *testing.T) {
	tests := []struct {
		name string
		want Health
	}{
		{
			name: "Happy Path",
			want: &health{
				logger: log.New(os.Stdout, "health :=> ", log.LstdFlags),
				config: config.Config,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHealth(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHealth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_health_GetHealth(t *testing.T) {
	config.Config = &models.AppConfig{}
	health := NewHealth()
	req, err := http.NewRequest("GET", "/health/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := health.GetHealth()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
