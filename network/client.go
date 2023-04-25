package network

import (
	"app/linuxsystem"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var ServAddr = "http://0.0.0.0:8080"

func GetAuthToken(name string, password string) (string, error) {
	data := struct {
		Name          string
		HashPassoword string
	}{Name: name, HashPassoword: password}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	resp, err := http.Post(ServAddr+"/api/connection", "json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	var res struct {
		Message string
		Token   string
	}
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	return res.Token, nil
}

func AddSoftwares(name string, token string, softwares []linuxsystem.Software) error {
	client := &http.Client{}

	data := struct {
		Softwares []linuxsystem.Software
	}{Softwares: softwares}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", ServAddr+"/api/software", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("name", name)
	req.Header.Set("token", token)

	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return nil
}
