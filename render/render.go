package render

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/MRP/config"
	"github.com/MRP/model"
	"github.com/justinas/nosurf"
)

var path = "../view/"
var funcs = template.FuncMap{}
var App *config.AppConfig

func Init(app *config.AppConfig) {
	App = app
}

func setDefaultData(data *model.TemplateModel, r *http.Request) *model.TemplateModel {
	data.Error = App.Session.PopString(r.Context(), "error")
	data.Info = App.Session.PopString(r.Context(), "info")
	data.Warning = App.Session.PopString(r.Context(), "waring")
	data.CsrfToken = nosurf.Token(r)

	return data
}

func SetTemplate(w http.ResponseWriter, r *http.Request, file string, data *model.TemplateModel) error {
	var temp map[string]*template.Template

	if App.UseCache {
		temp = App.TemplateCache
	} else {
		temp, _ = GetTemplate()

	}
	tc, ok := temp[file]
	if !ok {
		return errors.New("Cant get Template")
	}
	buf := new(bytes.Buffer)
	datamodel := setDefaultData(data, r)

	_ = tc.Execute(buf, datamodel)
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func GetTemplate() (map[string]*template.Template, error) {
	templates := map[string]*template.Template{}
	match, err := filepath.Glob(fmt.Sprintf("%s*.page.html", path))
	if err != nil {
		log.Println(err)
	}
	for _, v := range match {
		base := filepath.Base(v)
		tc, err := template.New(base).Funcs(funcs).ParseFiles(v)
		if err != nil {
			log.Println(err)
		}
		match, err := filepath.Glob(fmt.Sprintf("%s*.layout.html", path))
		if err != nil {
			log.Println(err)
		}
		if len(match) > 0 {
			tc, err = tc.ParseGlob(fmt.Sprintf("%s*.layout.html", path))
			if err != nil {
				log.Println(err)
			}
			templates[base] = tc
		}
	}
	return templates, nil

}
