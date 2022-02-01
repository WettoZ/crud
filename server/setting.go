package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type setting struct {
	GrpcHost     string `json:"GrpcHost"`
	GrpcPort     string `json:"GrpcPort"`
	PgUser       string `json:"PgUser"`
	PgPasswd     string `json:"PgPasswd"`
	PgHost       string `json:"PgHost"`
	PgPort       string `json:"PgPort"`
	PgDB         string `json:"PgDB"`
	CountConnect int    `json:"CountConnect"`
}

var conf setting

func init() {
	file, err := os.Open("setting.json")
	if err != nil {
		fmt.Println("Can't open file ", err)
	}
	defer file.Close()

	buffer, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Can't read file ", err)
	}

	if err = json.Unmarshal(buffer, &conf); err != nil {
		fmt.Println("Can't unmarshal file ", err)
	}

}
