package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gravitl/netmaker/config"
	controller "github.com/gravitl/netmaker/controllers"
	"github.com/gravitl/netmaker/dnslogic"
	"github.com/gravitl/netmaker/functions"
	"github.com/gravitl/netmaker/models"
)

var VERSION = "v0.3"
var backend = "https://api.nusak.ca/"
var authorization = "secretkey"

type NodeStatus struct {
	Mac     string
	Network string
	Status  string
	Color   string
}

type Relay struct {
	Node  models.Node
	Nodes []models.Node
}

type Version struct {
	Backend string
	Mine    string
}

//PageData -contains data for html template
type PageData struct {
	Page       string
	Message    string
	Admin      bool
	Networks   []models.Network
	Nodes      []models.Node
	Users      []models.ReturnUser
	ExtClients []models.ExtClient
	DNS        []models.DNSEntry
	CustomDNS  []models.DNSEntry
	Version    Version
}

// Init fetches page data from backend
func (data *PageData) Init(c *gin.Context, page, message string) {
	data.Page = page
	data.Message = message
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
	extclients, err := functions.GetAllExtClients()
	if err != nil {
		fmt.Println("error getting external client data", err)
	}
	nodes, err := models.GetAllNodes()
	if err != nil {
		fmt.Println("error getting node data", err)
	}
	users, err := controller.GetUsers()
	if err != nil {
		fmt.Println("error getting user data", err)
	}
	var dnsEntries []models.DNSEntry
	var customDnsEntries []models.DNSEntry
	for _, net := range networks {
		entries, err := controller.GetNodeDNS(net.NetID)
		if err != nil {
			fmt.Println("error getting dns data", err)
		}
		dnsEntries = append(dnsEntries, entries...)
		entries, err = dnslogic.GetCustomDNS(net.NetID)
		if err != nil {
			fmt.Println("error getting custom dns data", err)
		}
		customDnsEntries = append(customDnsEntries, entries...)
	}
	if isAdmin {
		data.Networks = networks
		data.Nodes = nodes
		data.Users = users
		data.ExtClients = extclients
		data.DNS = dnsEntries
		data.CustomDNS = customDnsEntries
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

		var clients []models.ExtClient
		for _, client := range extclients {
			if SliceContains(allowedNets, client.Network) {
				clients = append(clients, client)
			}
			data.ExtClients = clients
		}
		var userdns []models.DNSEntry
		for _, dns := range dnsEntries {
			if SliceContains(allowedNets, dns.Network) {
				userdns = append(userdns, dns)
			}
			data.DNS = userdns
		}
		var customdns []models.DNSEntry
		for _, dns := range customDnsEntries {
			if SliceContains(allowedNets, dns.Network) {
				customdns = append(customdns, dns)
			}
			data.CustomDNS = customdns
		}

	}
	data.Version.Backend = GetBackendVersion()
	data.Version.Mine = VERSION
}

func GetBackendVersion() string {
	request, err := http.NewRequest(http.MethodGet, backend+"api/server/getconfig", nil)
	if err != nil {
		fmt.Println("error creating http request ", err)
		return ""
	}
	request.Header.Set("Authorization", "Bearer "+authorization)
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("error from backend ", err)
		return ""
	}
	var config config.ServerConfig
	json.NewDecoder(response.Body).Decode(&config)
	return config.Version
}

func GetAllExtClients() []models.ExtClient {
	var clients []models.ExtClient
	var client models.ExtClient
	client.ClientID = "clientid"
	client.Description = "description"
	client.PrivateKey = "private key"
	client.PublicKey = "my public key"
	client.Network = "net"
	client.Address = "10.2.2.23"
	client.IngressGatewayID = "tbd"

	clients = append(clients, client)
	return clients
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
