package systems

import (
	"html/template"
	"net/http"
	"os"
	"time"
	"log"
	"fmt"
)

type relatoryPage struct {
	Time      string
	IPAddr    string
	UserAgent string
	URLPath   string
}

func GenerateRelatory(r *http.Request) (*relatoryPage, error) {
	relatory := &relatoryPage{
		Time:      time.Now().Format("2006-01-02 15:04:05"),
		IPAddr:    r.RemoteAddr,
		UserAgent: r.Header.Get("User-Agent"),
		URLPath:   r.URL.Path,
	}

	fileName := fmt.Sprintf("./html/relatory_%s.html", time.Now().Format("2006-01-02"))	

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	tmpl := template.Must(template.New("relatorypage").Parse(`
<!DOCTYPE html>
<html>
<head>
    <title>Reports</title>
</head>
<body>
    <h1>Reports</h1>
    <p>Date: {{.Time}}</p>
    <p>IP Address: {{.IPAddr}}</p>
    <p>User-Agent: {{.UserAgent}}</p>
    <p>Path: {{.URLPath}}</p>
    <br>
</body>
</html>
`))

	err = tmpl.Execute(file, relatory)
	if err != nil {
		return nil, err
	}

	return relatory, nil
}
