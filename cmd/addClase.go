/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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

	"github.com/galileoluna/clasroom/db"

	"github.com/spf13/cobra"
)

// addClaseCmd represents the addClase command
var addClaseCmd = &cobra.Command{
	Use:   "addClase",
	Short: "Ingresamos el nombre de la clase que queremos cargar",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addClase called")

		var clase db.Clase

		clase = db.NewClase(0, args[0])

		db.InsertClase(clase)
	},
}

func init() {
	rootCmd.AddCommand(addClaseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addClaseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addClaseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
