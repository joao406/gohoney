package systems

import (
	"html/template"
	"io/ioutil"
	"log"
)

type relatoryPage struct {
	Time      string
	IPAddr    string
	UserAgent string
}

func GenerateRelatory(r *http.Request) (*relatoryPage, error) {
	relatory := &relatoryPage{
		Time:      time.Now().Format("2006-01-02 15:04:05"),
		IPAddr:    r.RemoteAddr,
		UserAgent: r.Header.Get("User-Agent"),
	}

	fileName := fmt.Sprintf("relatory_%s.html", time.Now().Format("2006-01-02"))	

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	tmpl := template.Must(template.New("relatorypage").Parse(`
<!DOCTYPE html>
<html>
<head>
    <title>Relatório</title>
</head>
<body>
    <h1>Relatório</h1>
    <p>Data e Hora: {{.Time}}</p>
    <p>Endereço IP: {{.IPAddr}}</p>
    <p>User-Agent: {{.UserAgent}}</p>
</body>
</html>
`))

	err = tmpl.Execute(file, relatory)
	if err != nil {
		return nil, err
	}

	return relatory, nil
}
