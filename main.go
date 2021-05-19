package main

import (
	"log"
	"net/http"

	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

func main() {
	images := http.FileServer(http.Dir("images/"))
	http.HandleFunc("/", mainhandler)
	http.Handle("/images/", http.StripPrefix("/images/", images))
	http.HandleFunc("/openTab.js", jshandler)

	//_ = http.ListenAndServe("localhost:8080", http.HandlerFunc(handler))
	log.Fatal(http.ListenAndServe(":8080", nil))

}
func imagehandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path)
}
func jshandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	http.ServeFile(w, r, "./openTab.js")
}
func mainhandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		_ = Page(props{
			title: r.URL.Path,
			path:  r.URL.Path,
		}).Render(w)
	case "POST":
		err := SaveNet(w, r)
		if err != nil {
			_ = ErrorPage(props{
				title: r.URL.Path,
				path:  r.URL.Path,
			}).Render(w)
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

type props struct {
	title string
	path  string
}

// Page is a whole document to output.
func Page(p props) g.Node {
	buttons := []ButtonData{{"Network Details", true}, {"Nodes", true}, {"Access Keys", true}, {"DNS", true}}
	return c.HTML5(c.HTML5Props{
		Title:    p.title,
		Language: "en",
		Head: []g.Node{
			Link(Rel("stylesheet"), Href("https://www.w3schools.com/w3css/4/w3.css")),
			StyleEl(Type("text/css"),
				g.Raw(".center{ margin:auto; width:100%; padding: 10px; text-align: center;}"),
				//g.Raw(".navbar{width:15%}"),
				g.Raw(".navbarbutton{text-align:right}"),
				g.Raw(".maincontent{margin-left:25%}"),
				g.Raw(".form-popup{display:none}"),
				g.Raw(".tabbuttons.label{cursor:pointer}"),
				g.Raw(".block{display:block}"),
				g.Raw(".switch{position:relative; display:block; width:600px; height34px;}"),
				g.Raw(".slider { position: absolute; cursor: pointer; top: 0; left: 0; right: 0; bottom: 0; background-color: #ccc; -webkit-transition: .4s; transition: .4s; }"),
				g.Raw(".slider:before { position: absolute; content: \"\"; height: 26px; width: 26px; left: 4px; bottom: 4px; background-color: white; -webkit-transition: .4s; transition: .4s; }"),
				g.Raw("input:checked + .slider{background-color:#2196F3;}"),
				g.Raw("input:checked + .slider:before { -webkit-transform: translateX(26px); -ms-transform: translateX(26px); transform: translateX(26px); }"),
				g.Raw(".slider.round{border-radius:34px;}"),
				g.Raw(".input[type=checkbox]{visibility: hidden;}"),
				g.Raw(".fieldset{display: inline;}"),
			),
			Link(Rel("javascript"), Href("./openTab.js")),
			Script(g.Attr("src", "openTab.js")),
		},
		Body: []g.Node{
			g.Attr("onLoad", "openTab('All Networks-Network Details')"),
			//Forms(),
			Banner(),
			ButtonGroup(buttons, "tabbutton", "changeColour(this); displayTab();"),
			VertButtonGroup(GetAllNetIDs("All Networks")),
			//Navbar(p.path, GetNetworks()),
			Div(Class("maincontent w3-container"),
				//H1(g.Text(p.title)),
				//P(g.Textf("Welcome to the page at %v.", p.path[1:])),
				//g.If(p.path == "/", All(GetAllNetIDs)),
				All(),
				//Display(p.path),
				//g.If(p.path != "/", Detail(p.path[1:])),
				Detail(),
			),
		},
	})
}

func ErrorPage(p props) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:    "Error",
		Language: "en",
		Head: []g.Node{
			StyleEl(Type("text/css"),
				g.Raw(".center{ margin:auto; width:100%; padding: 10px; text-align: center;}"),
				g.Raw(".maincontent{margin-left:25%}"),
			),
		},
		Body: []g.Node{
			g.Text("An error occured adding new Network"),
		},
	})
}
