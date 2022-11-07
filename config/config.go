package config

import (
	"text/template"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	Session         *scs.SessionManager
	AppSecure       bool
	AppSecureCookie bool
	UseCache        bool
	TemplateCache   map[string]*template.Template
}
