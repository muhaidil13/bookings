package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	UseCache        bool
	AppSecure       bool
	AppSecureCookie bool
	TemplateCache   map[string]*template.Template
	InfoLog         *log.Logger
	ErrorLog        *log.Logger
	InProduct       bool
	Session         *scs.SessionManager
}
