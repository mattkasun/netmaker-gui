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

func GetNetSummary() ([]NetSummary, error) {
	var body []models.Network
	var network NetSummary
	var result []NetSummary
	response, err := API("", http.MethodGet, "/api/networks", "secretkey")
	if err != nil {
		return result, err
	}
	err = json.NewDecoder(response.Body).Decode(&body)
	if err != nil {
		return result, err
	}
	for _, net := range body {
		network.Name = net.DisplayName
		result = append(result, network)
	}

	return result, err
}

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
		result = append(result, node)
	}

	return result, err
}
