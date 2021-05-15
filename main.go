package main

import (
	"log"
	"net/http"

	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

func main() {
	http.HandleFunc("/", mainhandler)
	http.HandleFunc("/images/netmaker2.png", imagehandler)
	http.HandleFunc("/openTab.js", jshandler)

	//_ = http.ListenAndServe("localhost:8080", http.HandlerFunc(handler))
	log.Fatal(http.ListenAndServe(":8080", nil))

}
func imagehandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./images/netmaker2.png")
}
func jshandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	http.ServeFile(w, r, "./openTab.js")
}
func mainhandler(w http.ResponseWriter, r *http.Request) {
	_ = Page(props{
		title: r.URL.Path,
		path:  r.URL.Path,
	}).Render(w)
}

type props struct {
	title string
	path  string
}

// Page is a whole document to output.
func Page(p props) g.Node {
	buttons := []string{"Network Details", "Nodes", "Access Keys", "DNS"}
	return c.HTML5(c.HTML5Props{
		Title:    p.title,
		Language: "en",
		Head: []g.Node{
			Link(Rel("stylesheet"), Href("https://www.w3schools.com/w3css/4/w3.css")),
			StyleEl(Type("text/css"),
				g.Raw(".center{ margin:auto; width:100%; padding: 10px; text-align: center;}"),
				g.Raw(".navbar{width:15%}"),
				g.Raw(".navbarbutton{text-align:right}"),
				g.Raw(".maincontent{margin-left:25%}"),
			),
			//Link(Rel("javascript"), Href("./openTab.js")),
			Script(g.Attr("src", "openTab.js")),
		},
		Body: []g.Node{
			g.Attr("onload", "openTab('Network Details')"),
			Banner(),
			ButtonGroup(buttons),
			Navbar(p.path, GetNetworks()),
			Div(Class("maincontent w3-container"),
				H1(g.Text(p.title)),
				P(g.Textf("Welcome to the page at %v.", p.path[1:])),
				g.If(p.path == "/", All()),
				//Display(p.path),
				g.If(p.path != "/", Detail(p.path[1:])),
			),
		},
	})
}
