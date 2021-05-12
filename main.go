package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gravitl/netmaker/models"
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

func main() {
	http.HandleFunc("/", mainhandler)
	http.HandleFunc("/images/netmaker2.png", imagehandler)

	//_ = http.ListenAndServe("localhost:8080", http.HandlerFunc(handler))
	log.Fatal(http.ListenAndServe(":8080", nil))

}
func imagehandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./images/netmaker2.png")
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
	return c.HTML5(c.HTML5Props{
		Title:    p.title,
		Language: "en",
		Head: []g.Node{
			Link(Rel("stylesheet"), Href("https://www.w3schools.com/w3css/4/w3.css")),
			StyleEl(Type("text/css"),
				g.Raw(".center{ margin:auto; width:100%; padding: 10px; text-align: center;}"),
				g.Raw(".navbar{width:25%}"),
				g.Raw(".maincontent{margin-left:25%"),
			),
		},
		Body: []g.Node{
			Div(Class("w3-container w3-blue center"),
				//H5(g.Text("Hello World")),
				g.Raw("<img src=images/netmaker2.png>"),
				//H5(Img(Source("netmaker2.png"))),
			),
			ButtonGroup(),
			Navbar(p.path, GetNetworks()),
			//	[]PageLink{
			//		{Path: "/foo", Name: "Foo"},
			//		{Path: "/bar", Name: "Bar"},
			//	}),
			Div(Class("maincontent w3-container"),
				H1(g.Text(p.title)),
				P(g.Textf("Welcome to the page at %v.", p.path)),
				//				c.Classes{"margin-left:25%"},
			),
		},
	})
}

func GetNetworks() (pagelinks []PageLink) {
	var networks []models.Network
	var pagelink PageLink
	response, err := API("", http.MethodGet, "/api/networks", "secretkey")
	if err != nil {
		return []PageLink{}
	}
	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&networks)
	for _, network := range networks {
		pagelink.Path = "/" + network.NetID
		pagelink.Name = network.NetID
		pagelinks = append(pagelinks, pagelink)
	}
	return pagelinks
}

func ButtonGroup() g.Node {
	return Div(
		Button(Class("w3-button w3-white w3-left"),
			g.Text("Add Network"),
		),
		Class("center btn-group w3-white"),
		Button(Class("w3-bar-item w3-button"),
			g.Text("Network Details"),
		),
		Button(Class("w3-bar-item w3-button"),
			g.Text("Nodes"),
		),
		Button(Class("w3-bar-item w3-button"),
			g.Text("Access Key"),
		),
		Button(Class("w3-bar-item w3-button"),
			g.Text("DNS"),
		),
		Button(Class("w3-button w3-white w3-right"),
			g.Text("Logout"),
		),
	)
}

type PageLink struct {
	Path string
	Name string
}

func Navbar(currentPath string, links []PageLink) g.Node {
	return Div(Class("navbar w3-sidebar w3-light-grey w3-bar-block"),
		//Ul(
		NavbarLink("/", "All Networks", currentPath),

		g.Group(g.Map(len(links), func(i int) g.Node {
			return NavbarLink(links[i].Path, links[i].Name, currentPath)
		})),
		//),

		Hr(),
	)
}

func NavbarLink(href, name, currentPath string) g.Node {
	return A(
		Class("w3-bar-item w3-button center"),
		Href(href),
		c.Classes{"is-active": currentPath == href},
		g.Text(name),
		//		c.Classes{"width:25%"},
	)
}

func API(data interface{}, method, url, authorization string) (*http.Response, error) {
	backendURL := "http://localhost:8081"
	var request *http.Request
	var err error
	if data != "" {
		payload, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		request, err = http.NewRequest(method, backendURL+url, bytes.NewBuffer(payload))
		if err != nil {
			return nil, err
		}
		request.Header.Set("Content-Type", "application/json")
	} else {
		request, err = http.NewRequest(method, backendURL+url, nil)
		if err != nil {
			return nil, err
		}
	}
	if authorization != "" {
		request.Header.Set("Authorization", "Bearer "+authorization)
	}
	client := http.Client{}
	return client.Do(request)
}
