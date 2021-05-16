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
			//g.Attr("onClick", "openTab('addnetwork')"),
			g.Attr("onClick", "document.getElementById('addnetwork').style.display='block'"),
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
			return Button(Class("w3-bar-item w3-button tabbuttons"),
				g.Text(buttons[i]),
				g.Attr("onClick", "openTab(event, '"+buttons[i]+"')"),
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
		AllNets(),
		AllNodes(),
		KeyHolder(),
		DNSHolder(),
		Forms(),
	)
}

func AllNodes() g.Node {
	nodes := GetAllNodes()
	if nodes == nil {
		g.Text("There are no nodes")
	}
	return Div(ID("Nodes"), Class("w3-container tab"),
		g.Group(g.Map(len(nodes), func(i int) g.Node {
			return FieldSet(
				Legend(g.Text(nodes[i].Name)),
				FieldSet(
					Legend(g.Text("Public Key")),
					Label(g.Text(nodes[i].PublicKey)),
				),
				FieldSet(
					Legend(g.Text("Listen Port")),
					Label(g.Textf("%v", nodes[i].ListenPort)),
				),
				FieldSet(
					Legend(g.Text("Last CheckIn")),
					Label(g.Textf("%v", nodes[i].LastCheckIn)),
				),
			)
		})),
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
	networks := GetAllNets()
	//make sure a network was returned.
	if networks == nil {
		return g.Text("nothing to see here")
	}
	return Div(ID("Network Details"), Class("w3-container tab"),
		g.Group(g.Map(len(networks), func(i int) g.Node {
			return FieldSet(
				Legend(g.Text(networks[i].DisplayName)),
				FieldSet(
					Legend(g.Text("AddressRange")),
					Label(g.Text(networks[i].AddressRange)),
				),
				FieldSet(
					Legend(g.Text("NodesLastModifed")),
					Label(g.Text(time.Unix(networks[i].NodesLastModified, 0).Format("Mon Jan 2 2016 15:04:05 MST"))),
				),
			)
		})),
	)
}

func Detail(netname string) g.Node {
	return Div(
		Net(netname),
		Nodes(netname),
		Keys(netname),
		DNS(netname),
	)
}

func Net(netname string) g.Node {
	network := GetNetwork(netname)
	buttons := []string{"Edit", "Save", "Cancel", "Delete"}

	return Div(ID("Network Details"), Class("w3-container tab"),
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

func Nodes(netname string) g.Node {
	nodes := GetNodes(netname)
	if nodes == nil {
		g.Text("There are no nodes")
	}
	return Div(ID("Nodes"), Class("w3-container tab"),
		g.Group(g.Map(len(nodes), func(i int) g.Node {
			return FieldSet(
				Legend(g.Text(nodes[i].Name)),
				FieldSet(
					Legend(g.Text("Public Key")),
					Label(g.Text(nodes[i].PublicKey)),
				),
				FieldSet(
					Legend(g.Text("Listen Port")),
					Label(g.Textf("%v", nodes[i].ListenPort)),
				),
				FieldSet(
					Legend(g.Text("Last CheckIn")),
					Label(g.Textf("%v", nodes[i].LastCheckIn)),
				),
			)
		})),
	)
}

func Keys(netname string) g.Node {
	keys := GetKeys(netname)
	return Div(ID("Access Keys"), Class("w3-container tab"),
		Button(Class("w3-button w3-white w3-center"),
			g.Text("Add New Access Key"),
		),
		g.Group(g.Map(len(keys), func(i int) g.Node {
			return FieldSet(
				Legend(g.Text(keys[i].Name)),
				Label(g.Textf("Uses: %v", keys[i].Uses)),
				Button(Class("w3-button w3-white w3-right"),
					g.Text("Delete Key"),
				),
			)
		})),
	)
}

func DNS(netname string) g.Node {
	dns := GetDNS(netname)
	return Div(ID("DNS"), Class("w3-container tab"),
		Button(Class("w3-button w3-white w3-center"),
			g.Text("Add Entry"),
		),
		g.Group(g.Map(len(dns), func(i int) g.Node {
			return Div(
				Table(Tr(Th(g.Text("Name")), Th(g.Text("Address"))),
					Tr(Td(g.Text(dns[i].Name)), Td(g.Text(dns[i].Address))),
				),
			)
		})),
	)
}

func Forms() g.Node {
	return (AddNetwork())
}

func AddNetwork() g.Node {
	return Div(ID("addnetwork"), Class("w3-modal"),
		Div(Class("w3-modal-content"),
			Div(Class("w3-container"),
				FormEl(g.Attr("action", "/"),
					g.Attr("method", "post"),
					H1(g.Text("New Network")),

					Input(
						g.Attr("type", "text"),
						g.Attr("placeholder", "Network Name*"),
						g.Attr("name", "netname"),
						g.Attr("required"),
					),
					Br(),
					Input(
						g.Attr("type", "text"),
						g.Attr("placeholder", "Address Range*"),
						g.Attr("name", "addressrange"),
						g.Attr("required"),
					),
					Br(),
					Input(
						g.Attr("type", "checkbox"),
						g.Attr("name", "dualstack"),
						g.Attr("value", "true"),
					),
					Label(
						g.Attr("for", "dualstack"),
						g.Text("Use Dual Stack (IPv6)?"),
					),
					Br(),
					Input(
						g.Attr("type", "checkbox"),
						g.Attr("name", "islocal"),
						g.Attr("value", "true"),
					),
					Label(
						g.Attr("for", "islocal"),
						g.Text("Is Local?"),
					),
					Br(),
					Button(
						Class("w3-white"),
						g.Attr("onClick", "document.getElementById('addnetwork').style.display='none'"),
						g.Text("Cancel"),
					),
					Input(
						//Class("w3-blue"),
						g.Attr("type", "submit"),
						g.Attr("value", "Submit"),
						g.Text("Create Network"),
					),
				),
			),
		),
	)
}
