package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/6233/jhcoin/blockchain"
)

const (
	port 		string = ":4000"
	templateDir string = "templates/"
)
var templates *template.Template

type homeData struct {
	PageTitle	string
	blocks	[]*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	data := homeData{
		PageTitle: "home",
		blocks : blockchain.GetBlockchain().AllBlock(),
	}
	templates.ExecuteTemplate(rw, "home", data)
}

func add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")

		blockchain.GetBlockchain().AddBlock(data)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}
}

func main() {
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}