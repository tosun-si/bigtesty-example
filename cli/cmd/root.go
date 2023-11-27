/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var projectId string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bigtesty run tests",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		print("Running tests with BigTesty for the GCP project : " + projectId)

		fmt.Println("##################### Args ##########################")
		fmt.Println(args)

		//argsGcloud := args[3:]

		//out, err := exec.Command("gcloud", argsGcloud...).Output()

		out, err := exec.Command("bash", "-c", "docker run -it -e PROJECT_ID=gb-poc-373711 -e SA_EMAIL=sa-dataflow-dev@gb-poc-373711.iam.gserviceaccount.com -e LOCATION=europe-west1 -e IAC_BACKEND_URL=gs://gb-poc-pulumi-state/bigtesty -e ROOT_TEST_FOLDER=tests -v /Users/mazlum/my-projects/custom/bigtesty-example/tests:/app/tests -v /Users/mazlum/my-projects/custom/bigtesty-example/tests/tables:/app/bigtesty/infra/resource/tables -v /Users/mazlum/.config/gcloud:/root/.config/gcloud -v /var/run/docker.sock:/var/run/docker.sock bigtesty").Output()
		//out, err := exec.Command("bash", "-c", "docker ps").Output()

		//out, err := exec.Command("bash", "-c", "ls -R").Output()

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(out))

	},
	//BashCompletionFunction: bash_completion_func,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&projectId, "project_id", "p", "", "GCP Project ID")
	rootCmd.MarkFlagRequired("project_id")
}
