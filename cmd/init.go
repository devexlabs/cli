/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/devexlabs/cli/internal/config"
	"github.com/devexlabs/cli/internal/docker"
	"github.com/devexlabs/cli/internal/utils"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var port string

var toolsQuestion = []*survey.Question{
	{
		Name: "tools",
		Prompt: &survey.MultiSelect{
			Message: "Choose your tools:",
			Options: []string{"awscli", "terraform", "kubectl"},
		},
		Validate: survey.Required,
	},
}

func versionAskQuestion(tool string, versions []string) []*survey.Question {
	return []*survey.Question{
		{
			Name: "version",
			Prompt: &survey.Select{
				Message: "Choose " + tool + " cli versions:",
				Options: versions,
			},
			Validate: survey.Required,
		},
	}
}

type VersionAnswer struct {
	Version string
}

type VersionHash struct {
	Version string
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init cli configuration",
	Long:  `Choose and build your tools.`,
	Run: func(cmd *cobra.Command, args []string) {
		toolsHash := map[string][]string{
			"awscli":    {"1", "2"},
			"terraform": {"1.0.6", "1.0.5", "1.0.4", "1.0.3", "1.0.2", "1.0.1", "1.0.0", "0.15.5", "0.15.4", "0.15.3", "0.15.2", "0.15.1", "0.15.0"},
			"kubectl":   {"1.22.1", "1.22.0", "1.21.4", "1.21.3", "1.21.2", "1.21.1", "1.21.0"},
		}

		toolAnswers := struct {
			Tools []string
		}{}

		tools := make(map[string]struct {
			Version string
		})

		err := survey.Ask(toolsQuestion, &toolAnswers)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		for tool, versions := range toolsHash {
			if utils.Contains(toolAnswers.Tools, tool) {
				var answer VersionAnswer

				err := survey.Ask(versionAskQuestion(tool, versions), &answer)

				tools[tool] = answer

				if err != nil {
					fmt.Println(err.Error())
					return
				}
			}
		}

		config.Write(tools)

		docker.WriteDockerfile(tools)
		docker.Build()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&port, "port", "p", ":4040", "Port to be used on HTTP Server")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
