package main

import (
	"fmt"

	"github.com/gravitl/netmaker/models"
)

type User struct {
	UserName string
	Password string
	IsAdmin  bool
}

type ErrorResponse struct {
	Code    int
	Message string
}

type Success struct {
	Code     int
	Message  string
	Response Auth
}

type Auth struct {
	UserName  string
	AuthToken string
}

//PageData -contains data for html template
type PageData struct {
	Page     string
	Networks []NetSummary
	Nodes    []NodeSummary
}

//Initializes (fetches) page data from backend
func (data *PageData) Init(page string) {
	data.Page = page
	networks, err := GetNetSummary()
	if err != nil {
		//panic(err)
		fmt.Println("error geting network data", err)
	}
	data.Networks = networks
	nodes, err := GetNodeSummary()
	if err != nil {
		fmt.Println("error getting node data", err)
	}
	data.Nodes = nodes

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
