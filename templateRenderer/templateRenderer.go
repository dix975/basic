package templateRenderer
import (
	"net/http"
	"html/template"
	"path/filepath"
	"os"
	"fmt"
	"strings"
	"errors"
	"dix975.com/basic/logger"
)

type TemplateConfig struct {
	RootPath  string    `json:"rootPath"`

}

var (
	templates = make(map[string]*template.Template)
	config TemplateConfig
	funcMap template.FuncMap
)

func dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values) % 2 != 0 {
		return nil, errors.New("invalid dict call")
	}
	dict := make(map[string]interface{}, len(values) / 2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, errors.New("dict keys must be strings")
		}
		dict[key] = values[i + 1]
	}
	return dict, nil
}

func jsSafe(value string) template.JS{

	return  template.JS(value)
}

func visit(path string, f os.FileInfo, err error) error {

	if (!f.IsDir()) {

		if strings.HasSuffix(path, ".html") {
			templatePath := config.RootPath

			name := strings.TrimPrefix(path, fmt.Sprintf("%v/templates/", templatePath))
			name = strings.TrimSuffix(name, ".html")

			t := template.New("")
			logger.Info.Printf("Adding [%d] functions\n", len(funcMap))
			t.Funcs(funcMap)


			check := func(err error) {
				if err != nil {
					panic(err)
				}
			}

			var err error
			_, err = t.ParseFiles(path)
			check(err)

			_, err = t.ParseGlob(fmt.Sprintf("%v/layout/*.html", templatePath))
			check(err)

			templates[name] = t

			fmt.Printf("Added templates: %s for path %s \n", name, path)

		}
	}

	return nil
}

func Init(templateConfig TemplateConfig, functionMap template.FuncMap) {


	funcMap = functionMap
	funcMap["dict"] = dict
	funcMap["jsSafe"] = jsSafe

	logger.Info.Printf("Init with [%d] functions\n", len(funcMap))


	config = templateConfig
	templatePath := config.RootPath

	logger.Info.Printf("Wild load template from [%v]\n", templatePath)

	filepath.Walk(fmt.Sprintf("%v/templates", templatePath), visit)

}

func RenderTemplate(name string, w http.ResponseWriter, data interface{}) {


	t := templates[name];

	err := t.ExecuteTemplate(w, "layout", data)
	if(err != nil){
		panic(err)
	}


}

func RenderNoLayoutTemplate(name string, w http.ResponseWriter, data interface{}) {


	t := templates[name];

	err := t.ExecuteTemplate(w, "sub-template", data)
	if(err != nil){
		panic(err)
	}


}
