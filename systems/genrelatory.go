package systems

import (
	"html/template"
	"net/http"
	"os"
	"time"
	"log"
	"fmt"
)

var RelatoryStore []*relatoryPage

type relatoryPage struct {
	Time      string
	IPAddr    string
	UserAgent string
	URLPath   string
}

type templateContext struct {
	Relatories []*relatoryPage
}

func GenerateRelatory(r *http.Request) (*relatoryPage, error) {
	relatory := &relatoryPage{
		Time:      time.Now().Format("2006-01-02 15:04:05"),
		IPAddr:    r.RemoteAddr,
		UserAgent: r.Header.Get("User-Agent"),
		URLPath:   r.URL.Path,
	}

	RelatoryStore = append(RelatoryStore, relatory)

	fileName := fmt.Sprintf("./html/relatory_%s.html", time.Now().Format("2006-01-02"))	

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	context := &templateContext{
		Relatories: RelatoryStore,
	}
	
	tmpl := template.Must(template.New("relatorypage").Parse(`
<!DOCTYPE html>
<html>
<head>
    <title>Reports</title>
</head>
<body>
    <h1>Reports</h1>
    {{range .}}
    <p>Date: {{.Time}}</p>
    <p>IP Address: {{.IPAddr}}</p>
    <p>User-Agent: {{.UserAgent}}</p>
    <p>Path: {{.URLPath}}</p>
    <br>
    {{end}}
</body>
</html>
`))

	err = tmpl.Execute(file, relatory)
	if err != nil {
		return nil, err
	}

	return relatory, nil
}

