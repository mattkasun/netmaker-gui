package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gravitl/netmaker/models"
)

var backendURL = "https://api.netmaker.nusak.ca"

func API(data interface{}, method, url, authorization string) (*http.Response, error) {
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

//func GetNetSummary() ([]NetSummary, error) {
//	var network NetSummary
//	var result []NetSummary
//	//response, err := API("", http.MethodGet, "/api/networks", "secretkey")
//	networks, err := models.GetNetworks()
//	if err != nil {
//		return result, err
//	}
//	//err = json.NewDecoder(response.Body).Decode(&body)
//	//if err != nil {
//	//	return result, err
//	//}
//	for _, net := range networks {
//		fmt.Println(net.NodesLastModified, net.NetworkLastModified)
//		network.ID = net.NetID
//		network.Name = net.DisplayName
//		network.Address = net.AddressRange
//		network.Keys = net.AccessKeys
//		network.NodeModified = time.Unix(net.NodesLastModified, 0).Format(time.UnixDate)
//		network.NetModified = time.Unix(net.NetworkLastModified, 0).Format(time.UnixDate)
//		result = append(result, network)
//	}
//	return result, err
//}
//
func GetNodeSummary() ([]NodeSummary, error) {
	var body []models.Node
	var node NodeSummary
	var result []NodeSummary
	response, err := API("", http.MethodGet, "/api/nodes", "secretkey")
	if err != nil {
		return result, err
	}
	err = json.NewDecoder(response.Body).Decode(&body)
	if err != nil {
		return result, err
	}
	for _, net := range body {
		node.Name = net.Name
		node.Network = net.Network
		node.PublicIP = net.Endpoint
		node.SubNet = net.Address
		result = append(result, node)
	}

	return result, err
}

//func GetNetDetails(net string) (models.Network, error) {
//	var body models.Network
//	response, err := API("", http.MethodGet, "/api/networks/"+net, "secretkey")
//
//	if err != nil {
//		return body, err
//	}
//	err = json.NewDecoder(response.Body).Decode(&body)
//	return body, err
//}
