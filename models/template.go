package models

import (
	"io"
	"log"
	"text/template"
	"time"
)

/*
@author: shg
@since: 2022/10/14
@desc: //TODO
*/

type TemplateBlog struct {
	*template.Template
}

type HTMLTemplate struct {
	Index      TemplateBlog
	Category   TemplateBlog
	Custom     TemplateBlog
	Detail     TemplateBlog
	Login      TemplateBlog
	Pigeonhole TemplateBlog
	Writing    TemplateBlog
}

func (t *TemplateBlog) WriteData(w io.Writer, data interface{}) {
	err := t.Execute(w, data)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}

func (t *TemplateBlog) WriteError(w io.Writer, err error) {
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}
func InitTemplate(templateDir string) (HTMLTemplate, error) {
	tps, err := parseTemplate(
		[]string{"index", "category", "custom", "detail", "login", "pigeonhole", "writing"},
		templateDir,
	)
	var htmlTemplate HTMLTemplate
	if err != nil {
		log.Println(err)
		return htmlTemplate, err
	}
	htmlTemplate.Index = tps[0]
	htmlTemplate.Category = tps[1]
	htmlTemplate.Custom = tps[2]
	htmlTemplate.Detail = tps[3]
	htmlTemplate.Login = tps[4]
	htmlTemplate.Pigeonhole = tps[5]
	htmlTemplate.Writing = tps[6]
	return htmlTemplate, nil
}

func parseTemplate(templates []string, templateDir string) ([]TemplateBlog, error) {
	var tps []TemplateBlog
	for _, view := range templates {
		viewName := view + ".html"
		t := template.New(viewName)
		home := templateDir + "home.html"
		// layout
		footer := templateDir + "layout/footer.html"
		header := templateDir + "layout/header.html"
		pagination := templateDir + "layout/pagination.html"
		postList := templateDir + "layout/post-list.html"
		personal := templateDir + "layout/personal.html"
		t = t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "dateFormat": DateFormat, "dateDay": DateDay})
		t, err := t.ParseFiles(templateDir+viewName, home, header, footer, pagination, postList, personal)
		if err != nil {
			log.Println("parse file failed:", err)
			return nil, err
		}
		var tp TemplateBlog
		tp.Template = t
		tps = append(tps, tp)
	}
	return tps, nil
}

func IsODD(num int) bool {
	return num&1 == 0
}
func GetNextName(strs []string, index int) string {
	return strs[index+1]
}
func DateFormat(pattern string) string {
	return time.Now().Format(pattern)
}

func DateDay(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}
