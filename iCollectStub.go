package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var templ = template.Must(template.New("top").Parse(topTemplate))

var headerMap = map[string]string{"a": "1", "b": "2"}

func handler(w http.ResponseWriter, r *http.Request) {
	for hdrKey, _ := range headerMap {
		fmt.Println(r.FormValue(hdrKey))
		if r.FormValue(hdrKey) != "" {
			headerMap[hdrKey] = r.FormValue(hdrKey)
		}
	}
	for hdrKey, hdrValue := range headerMap {
		w.Header().Set(hdrKey, hdrValue)
	}
	templ.Execute(w, headerMap)
}

func main() {
	port := ":8888"
	http.HandleFunc("/", handler)
	http.ListenAndServe(port, nil)
}

const topTemplate = `
<html>
<head>
<title>iCollect Stub</title>
<body>
<form action="/" method="POST">
{{ range $key, $value := . }}
<input name={{$key}}>{{$key}}, {{$value}}<br></input>
{{ end }}
<button type=submit>Change</button>
<form>
</body>
</html>
`
