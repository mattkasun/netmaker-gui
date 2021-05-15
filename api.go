package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gravitl/netmaker/models"
)

type PageLink struct {
	Path string
	Name string
}

func GetAllNodes() (nodes []models.Node) {
	response, err := API("", http.MethodGet, "/api/nodes", "secretkey")
	if err != nil {
		return []models.Node{}
	}
	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&nodes)
	return nodes
}

func GetAllNets() (networks []models.Network) {
	response, err := API("", http.MethodGet, "/api/networks", "secretkey")
	if err != nil {
		return []models.Network{}
	}
	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&networks)
	return networks
}

func GetNetworks() (pagelinks []PageLink) {
	var pagelink PageLink
	networks := GetAllNets()
	for _, network := range networks {
		pagelink.Path = "/" + network.NetID
		pagelink.Name = network.NetID
		pagelinks = append(pagelinks, pagelink)
	}
	return pagelinks
}

func GetNetwork(name string) (network models.Network) {
	response, err := API("", http.MethodGet, "/api/networks/"+name, "secretkey")
	if err != nil {
		return models.Network{}
	}
	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&network)
	return network
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
