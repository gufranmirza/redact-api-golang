// Package api configures an http server for administration and application resources.

package router

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gufranmirza/redact-api-golang/src/config"
	"github.com/gufranmirza/redact-api-golang/src/models"
	"github.com/gufranmirza/redact-api-golang/src/web/services/health"
	"github.com/gufranmirza/redact-api-golang/src/web/services/health/healthmock"
	"github.com/gufranmirza/redact-api-golang/src/web/services/redact"
	"github.com/gufranmirza/redact-api-golang/src/web/services/redact/redactmock"
)

func Test_router_Router(t *testing.T) {
	config.Config = &models.AppConfig{
		URLPrefix: "",
	}
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()

	mockHealth := healthmock.NewMockHealth(ctrl)
	mockRedact := redactmock.NewMockRedact(ctrl)
	r := http.NewServeMux()
	urlPrefix := config.Config.URLPrefix

	type fields struct {
		config *models.AppConfig
		health health.Health
		redact redact.Redact
	}
	tests := []struct {
		name   string
		fields fields
		want   *http.ServeMux
	}{
		{
			name: "Happy Path",
			fields: fields{
				config: &models.AppConfig{},
				health: func() *healthmock.MockHealth {
					mockHealth.EXPECT().GetHealth().Return(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { return })).AnyTimes()
					r.Handle(urlPrefix+"/health/", mockHealth.GetHealth())
					return mockHealth
				}(),
				redact: func() *redactmock.MockRedact {
					mockRedact.EXPECT().RedactJSON().Return(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { return })).AnyTimes()
					r.Handle(urlPrefix+"/redact/", mockRedact.RedactJSON())
					return mockRedact
				}(),
			},
			want: r,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := &router{
				config: tt.fields.config,
				health: tt.fields.health,
				redact: tt.fields.redact,
			}
			if got := router.Router(); got != nil && tt.want == nil {
				t.Errorf("router.Router() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRouter(t *testing.T) {
	tests := []struct {
		name string
		want Router
	}{
		{
			name: "Happy Path",
			want: &router{
				config: config.Config,
				health: health.NewHealth(),
				redact: redact.NewRedact(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRouter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}
