package model

type TemplateModel struct {
	MapString map[string]string
	MapInt    map[int]int
	Data      interface{}
	Info      string
	Error     string
	Warning   string
	CsrfToken string
}
