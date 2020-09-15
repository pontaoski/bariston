package templating

import (
	"html/template"
	"io"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/qor/render"
)

var renderer = render.New(&render.Config{
	ViewPaths:     []string{"bot/templates"},
	DefaultLayout: "application",
	FuncMapMaker: func(*render.Render, *http.Request, http.ResponseWriter) template.FuncMap {
		return map[string]interface{}{
			"unescape": func(in string) template.HTML {
				return template.HTML(in)
			},
			"join_strings": strings.Join,
		}
	},
})

type Template struct{}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	d := dummyByte{}
	err := renderer.Execute(name, data, &http.Request{}, &d)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(d.String()))
	return err
}
