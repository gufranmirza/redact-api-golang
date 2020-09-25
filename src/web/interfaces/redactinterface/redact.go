package redactinterface

// RedactInterface represents redact request
type RedactInterface struct {
	JSONToRedact     map[string]interface{} `json:"json_to_redact,omitempty"`
	RedactRegexes    []RedactRegexes        `json:"redact_regexes,omitempty"`
	RedactCompletely []string               `json:"redact_completely,omitempty"`
}

// RedactRegexes represents regx to be applied
type RedactRegexes struct {
	Path    string   `json:"path,omitempty"`
	Regexes []string `json:"regexes,omitempty"`
}
