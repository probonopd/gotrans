// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/skitta/gotrans/api/baidu"
	"github.com/spf13/cobra"
)

var toLang string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "gotrans",
	Short: "Comandline tool for translating",
	Long:  `A simple translator using Baidu fanyi API`,
	RunE:  translate,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	// cobra.OnInitialize(initConfig)

	RootCmd.Flags().StringVarP(&toLang, "to", "t", "zh", "translate to")
}

// initConfig reads in config file and ENV variables if set.
// func initConfig() {
// 	viper.SetConfigName("config")
// 	viper.SetConfigType("yaml")
// 	viper.AddConfigPath(".")	// TODO: cross platform support

// 	if err := viper.ReadInConfig(); err != nil {
// 		panic(fmt.Errorf("Fatal error config file: %s \n", err))
// 	}
// }

func translate(cmd *cobra.Command, args []string) error {
	result, err := baidu.Translator(strings.Join(args, " "), "auto", toLang)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
