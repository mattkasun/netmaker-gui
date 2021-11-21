package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	controller "github.com/gravitl/netmaker/controllers"
	"github.com/gravitl/netmaker/logic"
	"github.com/gravitl/netmaker/models"
	"github.com/skip2/go-qrcode"
)

func CreateIngressClient(c *gin.Context) {
	var client models.ExtClient
	client.Network = c.Param("net")
	client.IngressGatewayID = c.Param("mac")
	node, err := logic.GetNodeByMacAddress(client.Network, client.IngressGatewayID)
	if err != nil {
		fmt.Println(err)
		ReturnError(c, http.StatusBadRequest, err, "Nodes")
		return
	}
	client.IngressGatewayEndpoint = node.Endpoint + ":" + strconv.FormatInt(int64(node.ListenPort), 10)

	err = controller.CreateExtClient(client)
	if err != nil {
		fmt.Println(err)
		ReturnError(c, http.StatusBadRequest, err, "Nodes")
		return
	}
	ReturnSuccess(c, "ExtClients", "external client has been created")
}

func DeleteIngressClient(c *gin.Context) {
	net := c.Param("net")
	id := c.Param("id")
	err := controller.DeleteExtClient(net, id)
	if err != nil {
		fmt.Println(err)
		ReturnError(c, http.StatusBadRequest, err, "Nodes")
		return
	}
	ReturnSuccess(c, "ExtClients", "external client "+id+" @ "+net+" has been deleted")
}

//EditIngressClient displays a form to update name of external client
func EditIngressClient(c *gin.Context) {
	net := c.Param("net")
	id := c.Param("id")
	client, err := controller.GetExtClient(id, net)
	if err != nil {
		fmt.Println(err)
		ReturnError(c, http.StatusBadRequest, err, "Nodes")
		return
	}
	c.HTML(http.StatusOK, "EditExtClient", client)
}

func GetQR(c *gin.Context) {
	net := c.Param("net")
	id := c.Param("id")
	config, err := GetConf(net, id)
	if err != nil {
		fmt.Println(err)
		ReturnError(c, http.StatusBadRequest, err, "ExtClient")
		return
	}
	b, err := qrcode.Encode(config, qrcode.Medium, 220)
	if err != nil {
		fmt.Println(err)
		ReturnError(c, http.StatusBadRequest, err, "ExtClient")
		return
	}
	c.Header("Content-Type", "image/png")
	c.Data(http.StatusOK, "application/octet-strean", b)
}

func GetConf(net, id string) (string, error) {
	client, err := controller.GetExtClient(id, net)
	if err != nil {
		return "", err
	}
	gwnode, err := logic.GetNodeByMacAddress(client.Network, client.IngressGatewayID)
	if err != nil {
		return "", err
	}
	network, err := logic.GetParentNetwork(client.Network)
	if err != nil {
		return "", err
	}
	keepalive := ""
	if network.DefaultKeepalive != 0 {
		keepalive = "PersistentKeepalive = " + strconv.Itoa(int(network.DefaultKeepalive))
	}
	gwendpoint := gwnode.Endpoint + ":" + strconv.Itoa(int(gwnode.ListenPort))
	newAllowedIPs := network.AddressRange
	if egressGatewayRanges, err := logic.GetEgressRangesOnNetwork(&client); err == nil {
		for _, egressGatewayRange := range egressGatewayRanges {
			newAllowedIPs += "," + egressGatewayRange
		}
	}
	defaultDNS := ""
	if network.DefaultExtClientDNS != "" {
		defaultDNS = "DNS = " + network.DefaultExtClientDNS
	}

	config := fmt.Sprintf(`[Interface]
Address = %s
PrivateKey = %s
%s

[Peer]
PublicKey = %s
AllowedIPs = %s
Endpoint = %s
%s

`, client.Address+"/32",
		client.PrivateKey,
		defaultDNS,
		gwnode.PublicKey,
		newAllowedIPs,
		gwendpoint,
		keepalive)

	return config, nil
}

func GetClientConfig(c *gin.Context) {
	net := c.Param("net")
	id := c.Param("id")
	config, err := GetConf(net, id)
	b := []byte(config)
	if err != nil {
		fmt.Println(err)
		ReturnError(c, http.StatusBadRequest, err, "ExtClient")
		return
	}
	filename := id + ".conf"
	//c.FileAttachment(filepath, filename)
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment: filename="+filename)
	c.Data(http.StatusOK, "application/octet-stream", b)
}

//UpdateClient updates name of external Client
func UpdateClient(c *gin.Context) {
	net := c.Param("net")
	id := c.Param("id")
	newid := c.PostForm("newid")
	client, err := controller.GetExtClient(id, net)
	if err != nil {
		fmt.Println(err)
		ReturnError(c, http.StatusBadRequest, err, "ExtClient")
		return
	}
	_, err = controller.UpdateExtClient(newid, net, client)
	if err != nil {
		fmt.Println(err)
		ReturnError(c, http.StatusBadRequest, err, "ExtClient")
		return
	}
	ReturnSuccess(c, "ExtClients", "external client has been updated")
}
