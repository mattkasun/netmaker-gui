package main

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	controller "github.com/gravitl/netmaker/controllers"
	"github.com/gravitl/netmaker/models"
)

type NodeUpdate struct {
	Oldmac  string
	Oldnet  string
	NewNode *models.Node
}

type NodeStatus struct {
	Mac    string
	Status string
	Color  string
}

//type User struct {
//	UserName string
//	Password string
//	IsAdmin  bool
//}
//
//type ErrorResponse struct {
//	Code    int
//	Message string
//}
//
//type Success struct {
//	Code     int
//	Message  string
//	Response Auth
//}
//
//type Auth struct {
//	UserName  string
//	AuthToken string
//}

//PageData -contains data for html template
type PageData struct {
	Page     string
	Admin    bool
	Networks []models.Network
	Nodes    []models.Node
	Users    []models.ReturnUser
}

//Initializes (fetches) page data from backend
func (data *PageData) Init(page string, c *gin.Context) {
	data.Page = page
	session := sessions.Default(c)
	user := session.Get("username").(string)
	isAdmin := session.Get("isAdmin").(bool)
	data.Admin = isAdmin
	allowedNets := session.Get("networks").([]string)
	networks, err := models.GetNetworks()
	if err != nil {
		//panic(err)
		fmt.Println("error geting network data", err)
	}
	nodes, err := models.GetAllNodes()
	if err != nil {
		fmt.Println("error getting node data", err)
	}
	users, err := controller.GetUsers()
	if err != nil {
		fmt.Println("error getting user data", err)
	}
	if isAdmin {
		data.Networks = networks
		data.Nodes = nodes
		data.Users = users
	} else {
		var nets []models.Network
		for _, network := range networks {
			if SliceContains(allowedNets, network.NetID) {
				nets = append(nets, network)
			}
			data.Networks = nets
		}
		var hosts []models.Node
		for _, node := range nodes {
			if SliceContains(allowedNets, node.Network) {
				hosts = append(hosts, node)
			}
			data.Nodes = hosts
		}
		user := models.ReturnUser{user, allowedNets, isAdmin}

		data.Users = append([]models.ReturnUser{}, user)
	}

}

func SliceContains(s []string, x string) bool {
	if len(s) == 0 {
		return false
	}
	for i := range s {
		if s[i] == x {
			return true
		}
	}
	return false
}

//NetSummary - contains summary network data for html template
type NetSummary struct {
	Name         string
	ID           string
	Address      string
	NodeModified string
	NetModified  string
	Keys         []models.AccessKey
	//AddressRange        string
	//NodeLastModified    string
	//NetworkLastModified string
}

//NodeSummary - contains summary node data for html template
type NodeSummary struct {
	Name     string
	Network  string
	PublicIP string
	SubNet   string
	//	PublicIP    string
	//	SubNetIP    string
	//	Status      string
	//	PublicKey   string
	//	ListenPort  string
	//	LastCheckin string
}

//NetDetail - contains detailed network data for html template
//type NetDetail struct {
//	Name                string
//	AddressRange        string
//	IP6Address          string
//	LocalRange          string
//	DisplayName         string
//	NodeLastModified    string
//	NetworkLastModified string
//	DefaultInterface    string
//	DefaultPort         string
//	PostUp              string
//	PostDown            string
//	KeepAlive           string
//	CheckinInterval     string
//	DualStack           bool
//	SaveConfig          bool
//	UDPHolePunch        bool
//	KeyRequired         bool
//}

//NodeDetail - contains detailed node data for html template
type NodeDetail struct {
	Name               string
	Interface          string
	Network            string
	AddressRange       string
	IP6Address         string
	ListenPort         string
	PublicKey          string
	Endpoint           string
	Expires            string
	PostUp             string
	PostDown           string
	KeepAlive          string
	CheckinInterval    string
	EgressGatewayRange string
	AllowedIPs         string
	DualStack          bool
	SaveConfig         bool
	UDPHolePunch       bool
	KeyRequired        bool
	Static             bool
}

type NewNet struct {
	Name    string
	Address string
	Dual    string
	Local   string
	UDP     string
}
