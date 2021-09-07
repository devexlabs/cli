package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ToolJSONDefinition struct {
	DefaultVersion string   `json:"default_version"`
	Versions       []string `json:"versions"`
}

type JSONStructure struct {
	AwsCli    ToolJSONDefinition `json:"awscli"`
	Terraform ToolJSONDefinition `json:"terraform"`
	Kubectl   ToolJSONDefinition `json:"kubectl"`
}

func ParseData(tools map[string]struct{ Version string }) JSONStructure {
	return JSONStructure{
		AwsCli: ToolJSONDefinition{
			DefaultVersion: tools["awscli"].Version,
			Versions:       []string{tools["awscli"].Version},
		},
		Terraform: ToolJSONDefinition{
			DefaultVersion: tools["terraform"].Version,
			Versions:       []string{tools["terraform"].Version},
		},
		Kubectl: ToolJSONDefinition{
			DefaultVersion: tools["kubectl"].Version,
			Versions:       []string{tools["kubectl"].Version},
		},
	}
}

func Write(tools map[string]struct{ Version string }) {
	data := ParseData(tools)

	file, _ := json.MarshalIndent(data, "", "  ")

	_ = ioutil.WriteFile("cli.conf.json", file, 0644)

	fmt.Println("Write cli.conf.json")
}
