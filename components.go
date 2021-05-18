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
func ButtonGroup(buttons []string, class string) g.Node {
	return Div(Class("center btn-group w3-white"),
		g.Group(g.Map(len(buttons), func(i int) g.Node {
			return Button(Class("w3-bar-item w3-button "+class),
				g.Text(buttons[i]),
				//g.Attr("onClick", "openTab('"+buttons[i]+"'); changeColour(this);"),
				g.Attr("onClick", "changeColour(this); displayTab();"),
				g.If(i != 0, g.Attr("data-selected", "false")),
				g.If(i == 0, g.Attr("data-selected", "true")),
			)
		})),
	)
}

func VertButtonGroup(buttons []string) g.Node {
	return Div(Class("w3-sidebar w3-bar-block"),
		g.Group(g.Map(len(buttons), func(i int) g.Node {
			return Button(Class("w3-button netbutton block"),
				g.Text(buttons[i]),
				//g.Attr("onClick", "navTo('"+buttons[i]+"'); changeColourNets(this);"),
				g.Attr("onClick", "changeColour(this); displayTab();"),
				g.If(i != 0, g.Attr("data-selected", "false")),
				g.If(i == 0, g.Attr("data-selected", "true")),
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
		Br(),
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
	return Div(ID("All Networks-Nodes"), Class("w3-container tab"),
		g.Group(g.Map(len(nodes), func(i int) g.Node {
			return FieldSet(
				Legend(g.Text(nodes[i].Name)),
				Button(ID("edit"+networks[i].NetID), Class("w3-right"),
					Img(g.Attr("src", "/images/edit.png"),
						g.Attr("alt", "edit"),
						g.Attr("width", "24")),
					// ------------this need updating -------
					g.Attr("onClick", "document.getElementById('editnetwork').style.display='block'"),
				),
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
	return Div(ID("All Networks-Access Keys"), Class("w3-container tab"),
		g.Text("Please select a specific network to view it's access keys"),
	)
}
func DNSHolder() g.Node {
	return Div(ID("All Networks-DNS"), Class("w3-container tab"),
		g.Text("Please select a specific network to view it's DNS"),
	)
}

func AllNets() g.Node {
	networks := GetAllNets()
	//make sure a network was returned.
	if networks == nil {
		return g.Text("nothing to see here")
	}
	return Div(ID("All Networks-Network Details"),
		Class("w3-container tab"), g.Attr("display", "inline"),
		P(g.Text("All Networks")),
		g.Group(g.Map(len(networks), func(i int) g.Node {
			return FieldSet(g.Attr("display", "inline"),
				Legend(g.Text(networks[i].DisplayName)),
				Button(ID("edit"+networks[i].NetID), Class("w3-right"),
					Img(g.Attr("src", "/images/edit.png"),
						g.Attr("alt", "edit"),
						g.Attr("width", "24")),
					// ------------this need updating -------
					g.Attr("onClick", "document.getElementById('editnetwork').style.display='block'"),
				),
				Button(ID("refresh"+networks[i].NetID), Class("w3-left"),
					Img(g.Attr("src", "/images/refresh.png"),
						g.Attr("alt", "refresh"),
						g.Attr("width", "24")),
					// ------------this need updating -------
					g.Attr("onClick", "document.getElementById('editnetwork').style.display='block'"),
				),
				Button(ID("addserver"+networks[i].NetID), Class("w3-left"),
					Img(g.Attr("src", "/images/plus.png"),
						g.Attr("alt", "addserver"),
						g.Attr("width", "24")),
					// ------------this need updating -------
					g.Attr("onClick", "document.getElementById('editnetwork').style.display='block'"),
				),
				FieldSet(
					Legend(g.Text("AddressRange")),
					Label(g.Text(networks[i].AddressRange)),
				),
				FieldSet(
					Legend(g.Text("NodesLastModifed")),
					Label(g.Text(time.Unix(networks[i].NodesLastModified, 0).Format("Mon Jan 2 2016 15:04:05 MST"))),
				),
				FieldSet(
					Legend(g.Text("NetworkLastModifed")),
					Label(g.Text(time.Unix(networks[i].NetworkLastModified, 0).Format("Mon Jan 2 2016 15:04:05 MST"))),
				),
			)
		})),
	)
}

func Detail() g.Node {
	networks := GetAllNetIDs("")
	return g.Group(g.Map(len(networks), func(i int) g.Node {
		//for _, netname := range networks {
		return Div(
			Net(networks[i]),
			Nodes(networks[i]),
			Keys(networks[i]),
			DNS(networks[i]),
		)
	}))
}

func Net(netname string) g.Node {
	network := GetNetwork(netname)
	buttons := []string{"Edit", "Save", "Cancel", "Delete"}

	return Div(ID(netname+"-Network Details"), Class("w3-container tab"),
		ButtonGroup(buttons, "netbuttons"),
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
	return Div(ID(netname+"-Nodes"), Class("w3-container tab"),
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
	return Div(ID(netname+"-Access Keys"), Class("w3-container tab"),
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
	return Div(ID(netname+"-DNS"), Class("w3-container tab"),
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
