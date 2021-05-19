package main

import (
	"fmt"
	"time"

	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

type ButtonData struct {
	Name    string
	Enabled bool
}

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
func ButtonGroup(buttons []ButtonData, class, onclick string) g.Node {
	//func ButtonGroup(buttons []string, class string) g.Node {
	return Div(Class("center btn-group w3-white"),
		g.Group(g.Map(len(buttons), func(i int) g.Node {
			return Button(Class("w3-bar-item w3-button "+class),
				g.Text(buttons[i].Name),
				//g.Attr("onClick", "openTab('"+buttons[i]+"'); changeColour(this);"),
				g.If(!buttons[i].Enabled, g.Attr("disabled")),
				g.Attr("onClick", onclick),
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
		P(g.Text("All Nodes")),
		g.Group(g.Map(len(nodes), func(i int) g.Node {
			return FieldSet(
				Legend(g.Text(nodes[i].Name)),
				g.Text("public ip:"),
				g.Text(nodes[i].Endpoint),
				g.Text(" | subnet ip:"),
				g.Text(nodes[i].Address),
				g.Text(" | status:"),
				g.Text("TODO"),
				Button(ID("edit"+nodes[i].Name), Class("w3-button w3-right"),
					Img(g.Attr("src", "/images/edit.png"),
						g.Attr("alt", "edit"),
						g.Attr("width", "24")),
					// ------------this need updating -------
					g.Attr("onClick", "document.getElementById('editnetwork').style.display='block'"),
				),
				Button(ID("gateway"+nodes[i].Name), Class("w3-button w3-left"),
					Img(g.Attr("src", "/images/network.png"),
						g.Attr("alt", "add gateway"),
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
					Label(g.Text(time.Unix(nodes[i].LastCheckIn, 0).Format("Mon Jan 2 2016 15:04:05 MST"))),
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
				Button(ID("edit"+networks[i].NetID), Class("w3-button w3-right"),
					Img(g.Attr("src", "/images/edit.png"),
						g.Attr("alt", "edit"),
						g.Attr("width", "24")),
					// ------------this need updating -------
					g.Attr("onClick", "document.getElementById('editnetwork').style.display='block'"),
				),
				Br(), Br(),
				Button(ID("refresh"+networks[i].NetID), Class("w3-button w3-left"),
					Img(g.Attr("src", "/images/refresh.png"),
						g.Attr("alt", "refresh"),
						g.Attr("width", "24")),
					// ------------this need updating -------
					g.Attr("onClick", "document.getElementById('editnetwork').style.display='block'"),
				),
				Button(ID("addserver"+networks[i].NetID), Class("w3-button w3-left"),
					Img(g.Attr("src", "/images/plus.png"),
						g.Attr("alt", "addserver"),
						g.Attr("width", "24")),
					// ------------this need updating -------
					g.Attr("onClick", "document.getElementById('editnetwork').style.display='block'"),
				),
				FieldSet(g.Attr("display", "inline"),
					Legend(g.Text("AddressRange")),
					Label(g.Text(networks[i].AddressRange)),
				),
				FieldSet(g.Attr("display", "inline"),
					Legend(g.Text("NodesLastModifed")),
					Label(g.Text(time.Unix(networks[i].NodesLastModified, 0).Format("Mon Jan 2 2016 15:04:05 MST"))),
				),
				FieldSet(g.Attr("display", "inline"),
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
	buttons := []ButtonData{{"Edit", true}, {"Save", false}, {"Cancel", false}, {"Delete", true}}

	return Div(ID(netname+"-Network Details"), Class("w3-container tab"),
		c.Classes{"width:600px;": true},
		ButtonGroup(buttons, "netbuttons", ""),
		Label(Class("switch"),
			Input(g.Attr("type", "checkbox")),
			//g.Raw("<span class=\"slider round\"></span>"),
			(g.Text("Allow Node SignUp without Keys")),
		),
		//Editable
		FieldSet(
			Legend(g.Text("Address Range")),
			Input(
				ID(netname+"."+network.AddressRange),
				g.Attr("type", "text"),
				g.Attr("name", "AddressRange"),
				g.Attr("placeholder", network.AddressRange),
				g.Attr("disabled"),
			),
		),
		FieldSet(
			Legend(g.Text("Address Range (IPv6)")),
			Label(g.Text(network.AddressRange6)),
		),
		//Editable
		FieldSet(
			Legend(g.Text("Local Range")),
			Input(
				ID(netname+"."+network.LocalRange),
				g.Attr("type", "text"),
				g.Attr("name", "LocalRange"),
				g.Attr("placeholder", network.LocalRange),
				g.Attr("disabled"),
			),
		),
		//Editable
		FieldSet(
			Legend(g.Text("Display Name")),
			Input(
				ID(netname+"."+network.DisplayName),
				g.Attr("type", "text"),
				g.Attr("name", "DefaultDisplayName"),
				g.Attr("placeholder", network.DisplayName),
				g.Attr("disabled"),
			),
		),
		FieldSet(
			Legend(g.Text("NodesLastModifed")),
			Label(g.Text(time.Unix(network.NodesLastModified, 0).Format("Mon Jan 2 2016 15:04:05 MST"))),
		),
		FieldSet(
			Legend(g.Text("Network Last Modified")),
			Label(g.Text(time.Unix(network.NetworkLastModified, 0).Format("Mon Jan 2 2016 15:04:05 MST"))),
		),
		//Editable
		FieldSet(
			Legend(g.Text("Interface")),
			Input(
				ID(netname+"."+network.DefaultInterface),
				g.Attr("type", "text"),
				g.Attr("name", "DefaultInterface"),
				g.Attr("placeholder", network.DefaultInterface),
				g.Attr("disabled"),
			),
		),
		//Editable
		FieldSet(
			Legend(g.Text("Listen Port")),
			Input(
				ID(netname+"."+fmt.Sprintf("%v", network.DefaultListenPort)),
				g.Attr("type", "text"),
				g.Attr("name", "DefaultListenPort"),
				g.Attr("placeholder", fmt.Sprintf("%v", network.DefaultListenPort)),
				g.Attr("disabled"),
			),
		),
		//Editable
		FieldSet(
			Legend(g.Text("Post Up")),
			Input(
				ID(netname+"."+network.DefaultPostUp),
				g.Attr("type", "text"),
				g.Attr("name", "DefaultPostUp"),
				g.Attr("placeholder", network.DefaultPostUp),
				g.Attr("disabled"),
			),
		),
		//Editable
		FieldSet(
			Legend(g.Text("Post Down")),
			Input(
				ID(netname+"."+network.DefaultPostDown),
				g.Attr("type", "text"),
				g.Attr("name", "DefaultPostDown"),
				g.Attr("placeholder", network.DefaultPostDown),
				g.Attr("disabled"),
			),
		),
		//Editable
		FieldSet(
			Legend(g.Text("KeepAlive")),
			Input(
				ID(netname+"."+fmt.Sprintf("%v", network.DefaultKeepalive)),
				g.Attr("type", "text"),
				g.Attr("name", "DefaultKeepalive"),
				g.Attr("placeholder", fmt.Sprintf("%v", network.DefaultKeepalive)),
				g.Attr("disabled"),
			),
		),
		//Editable
		FieldSet(
			Legend(g.Text("Check In Interval")),
			Input(
				ID(netname+"."+fmt.Sprintf("%v", network.DefaultCheckInInterval)),
				g.Attr("type", "text"),
				g.Attr("name", "DefaultCheckInInterval"),
				g.Attr("placeholder", fmt.Sprintf("%v", network.DefaultCheckInInterval)),
				g.Attr("disabled"),
			),
		),
		Span(
			Label(Class("switch"), Input(g.Attr("type", "checkbox")), g.Text("Dual Stack")),
			Label(Class("switch"), Input(g.Attr("type", "checkbox")), g.Text("Save Config")),
		),
	)
}

func Nodes(netname string) g.Node {
	nodes := GetNodes(netname)
	return Div(ID(netname+"-Nodes"), Class("w3-container tab"),
		g.Text(netname+" nodes"),
		g.If(nodes == nil, P(g.Text("No nodes present in network"))),
		g.Group(g.Map(len(nodes), func(i int) g.Node {
			return FieldSet(
				Legend(g.Text(nodes[i].Name)),
				g.Text("public ip:"),
				g.Text(nodes[i].Endpoint),
				g.Text(" | subnet ip:"),
				g.Text(nodes[i].Address),
				g.Text(" | status:"),
				g.Text("TODO"),
				Button(ID("edit"+nodes[i].Name), Class("w3-button w3-right"),
					Img(g.Attr("src", "/images/edit.png"),
						g.Attr("alt", "edit"),
						g.Attr("width", "24")),
					// ------------this need updating -------
					g.Attr("onClick", "document.getElementById('editnetwork').style.display='block'"),
				),
				Button(ID("gateway"+nodes[i].Name), Class("w3-button w3-left"),
					Img(g.Attr("src", "/images/network.png"),
						g.Attr("alt", "add gateway"),
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
					Label(g.Text(time.Unix(nodes[i].LastCheckIn, 0).Format("Mon Jan 2 2016 15:04:05 MST"))),
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
	dns := GetDNS(netname, false)
	custom := GetDNS(netname, true)

	return Div(ID(netname+"-DNS"), Class("w3-container tab"),
		StyleEl(Type("text/css"),
			g.Raw("table{ border: 1px solid black}"),
			g.Raw("th { background: blue; color: white}"),
		),
		Button(Class("w3-button w3-white w3-center"),
			g.Text("Add Entry"),
		),
		Table(Tr(Th(g.Text("Node DNS (default)"))),
			g.Group(g.Map(len(dns), func(i int) g.Node {
				return Tr(Td(g.Text(dns[i].Name)), Td(g.Text(dns[i].Address)))
			})),
		),
		Table(Tr(Th(g.Text("Custom DNS"))),
			g.Group(g.Map(len(custom), func(i int) g.Node {
				return Tr(Td(g.Text(custom[i].Name)), Td(g.Text(custom[i].Address)),
					Td(Img(g.Attr("src", "/images/delete.png"),
						g.Attr("alt", "delete entry"),
						g.Attr("width", "24"),
					)),
				)
			})),
		),
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
