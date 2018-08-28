package urlshort

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"net/http"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	serve := func(w http.ResponseWriter, r *http.Request) {
		if val, ok := pathsToUrls[r.RequestURI]; ok {
			http.Redirect(w, r, val, 302)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
	return serve
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var paths []yamlPath
	if err := yaml.Unmarshal(yml, &paths); err != nil {
		return fallback.ServeHTTP, err
	}

	mapPaths := make(map[string]string)
	for _, p := range paths {
		mapPaths[p.Path] = p.URL
	}

	return MapHandler(mapPaths, fallback), nil
}

func serve(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
}

type yamlPath struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
