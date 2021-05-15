package main

import (
	"time"

	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

func Banner() g.Node {
	return Div(Class("w3-container w3-blue center"),
		Button(Class("w3-button w3-blue w3-left"),
			g.Text("Add Network"),
		),
		Img(g.Attr("src", "images/netmaker2.png")),
		Button(Class("w3-button w3-blue w3-right"),
			g.Text("Logout"),
		),
	)
}
func ButtonGroup(buttons []string) g.Node {
	return Div(Class("center btn-group w3-white"),
		g.Group(g.Map(len(buttons), func(i int) g.Node {
			return Button(Class("w3-bar-item w3-button"),
				g.Text(buttons[i]),
				g.Attr("onClick", "openTab('"+buttons[i]+"')"),
			)
		})),
	)
}

func Navbar(currentPath string, links []PageLink) g.Node {
	return Div(Class("navbar w3-sidebar w3-light-grey w3-bar-block"),
		NavbarLink("/", "All Networks", currentPath),
		g.Group(g.Map(len(links), func(i int) g.Node {
			return NavbarLink(links[i].Path, links[i].Name, currentPath)
		})),
		Hr(),
	)
}

func NavbarLink(href, name, currentPath string) g.Node {
	return A(
		Class("navbarbutton w3-bar-item w3-button center"),
		Href(href),
		c.Classes{"is-active": currentPath == href},
		g.Text(name),
	)
}

func All() g.Node {
	return Div(
		g.Attr("onload", "openTab('networks')"),
		AllNets(),
		AllNodes(),
		KeyHolder(),
		DNSHolder(),
	)
}

func AllNodes() g.Node {
	return Div(ID("Nodes"), Class("w3-container tab"),
		g.Text("All nodes"),
	)
}

func KeyHolder() g.Node {
	return Div(ID("Access Keys"), Class("w3-container tab"),
		g.Text("Please select a specific network to view it's access keys"),
	)
}
func DNSHolder() g.Node {
	return Div(ID("DNS"), Class("w3-container tab"),
		g.Text("Please select a specific network to view it's DNS"),
	)
}

func AllNets() g.Node {
	//	emptynet := models.Network{}
	networks := GetAllNets()
	//make sure a network was returned.
	if networks == nil {
		return g.Text("nothing to see here")
	}
	return g.Group(g.Map(len(networks), func(i int) g.Node {
		return Div(ID("Network Details"), Class("w3-container tab"),
			FieldSet(
				Legend(g.Text(networks[i].DisplayName)),
				FieldSet(
					Legend(g.Text("AddressRange")),
					Label(g.Text(networks[i].AddressRange)),
				),
				FieldSet(
					Legend(g.Text("NodesLastModifed")),
					Label(g.Text(time.Unix(networks[i].NodesLastModified, 0).Format("Mon Jan 2 2016 15:04:05 MST"))),
				),
			),
		)
	}))
}

func Detail(netname string) g.Node {
	network := GetNetwork(netname)
	buttons := []string{"Edit", "Save", "Cancel", "Delete"}

	return Div(ID("networks"), Class("w3-container tab"),
		ButtonGroup(buttons),
		Input(Class("switch slider round"), Label(g.Text("Allow Node SignUp without Keys"))),
		FieldSet(
			Legend(g.Text("AddressRange")),
			Label(g.Text(network.AddressRange)),
		),
		FieldSet(
			Legend(g.Text("NodesLastModifed")),
			Label(g.Text(time.Unix(network.NodesLastModified, 0).Format("Mon Jan 2 2016 15:04:05 MST"))),
		),
	)
}
