package models

const (
	// DefaultConfigPath is path used when config path is not provided
	DefaultConfigPath = "./app-config.json"
)

// AppConfig  is the overall config that file that our application will use
type AppConfig struct {
	Hostname        string `json:"Hostname"`
	Port            int    `json:"Port"`
	ServiceName     string `json:"ServiceName"`
	ServiceProvider string `json:"ServiceProvider"`
	ServiceVersion  string `json:"ServiceVersion"`
	APIVersionV1    string `json:"APIVersionV1"`
	URLPrefix       string `json:"URLPrefix"`
}
