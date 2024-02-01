package main

import (
	"fmt"
	"log"
	"flag"
	"os"
	"net/http"

	"github.com/fatih/color"
	"github.com/joao406/goney/systems"
)

var ascii = `
 ░▒▓██████▓▒░ ░▒▓██████▓▒░░▒▓███████▓▒░░▒▓████████▓▒░▒▓█▓▒░░▒▓█▓▒░ 
░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░ 
░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░ 
░▒▓█▓▒▒▓███▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓██████▓▒░  ░▒▓██████▓▒░  
░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░         ░▒▓█▓▒░     
░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░         ░▒▓█▓▒░     
 ░▒▓██████▓▒░ ░▒▓██████▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓████████▓▒░  ░▒▓█▓▒░     
Created by github.com/joao406

`

func handler(w http.ResponseWriter, r *http.Request) {
	ipAddress := r.RemoteAddr
	userAgent := r.Header.Get("User-Agent")
	
	log.Printf(color.GreenString("[+] NEW CLIENT: %s %s", ipAddress, userAgent))
}

func main() {
	color.Red(ascii)
	// Args
	var portArg int
	var showHelp bool
	flag.IntVar(&portArg, "p", 8080, "Port to listen server")
	flag.BoolVar(&showHelp, "h", false, "Help screen")
	flag.Parse()

	if showHelp {
		fmt.Println("USE -h to show help screen.")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if portArg == 0 {
		fmt.Println("No port value specified!")
		return
	} else {
		color.Yellow("[*] Checking if html pages exists...")
		if _, err := os.Stat("./html/index.html"); err != nil {
			log.Fatal(err)
			os.Exit(1)	
		}
		color.Green("[+] File check OK!")
	
		fmt.Printf("HONEYPOT RUNNING ON *:%d\n\n", portArg)
		http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, err := systems.GenerateRelatory(r)
			if err != nil {
				log.Println("Error generating relatory:", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			
			handler(w, r)
			http.FileServer(http.Dir("./html")).ServeHTTP(w, r)
		}))
		
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", portArg), nil))
	}
}
