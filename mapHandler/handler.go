package maphandler

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

func MapHandler(pathTOUrl map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathTOUrl[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func YMLHandler(ymlByte []byte, fallback http.Handler) (http.HandlerFunc, error) {
	fmt.Println(string(ymlByte))
	godotenv.Load()
	var pathUrls []pathURL
	err := yaml.Unmarshal(ymlByte, &pathUrls)
	if err != nil {
		return nil, err
	}
	fmt.Println("pathurls yml", pathUrls)
	pathToUrl := map[string]string{}
	for _, pathUrl := range pathUrls {
		pathToUrl[pathUrl.Path] = pathUrl.Url
	}
	return MapHandler(pathToUrl, fallback), nil
}

type pathURL struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}
