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
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	platform string
)

var openCmds = map[string]string{
	"windows": "cmd /c start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

var baikeCmd = &cobra.Command{
	Use:     "baike",
	Aliases: []string{"bk", "wk", "wiki"},
	Short:   "find things in baike site",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := findInBaike(args[0], platform)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(baikeCmd)
	baikeCmd.Flags().StringVarP(&platform, "platform", "p", "baidu", "platform to find things")
}

// 百科查找
func findInBaike(keyword, platform string) error {
	var link string
	// 百度百科搜索
	if platform == "baidu" || platform == "bd" {
		link = fmt.Sprintf("https://baike.baidu.com/item/%s", keyword)
	}
	// 互动百科搜索
	if platform == "hudong" || platform == "baike" || platform == "hd" {
		link = fmt.Sprintf("http://www.baike.com/wiki/%s", keyword)
	}
	// 维基百科搜索
	if platform == "wikipedia" || platform == "wiki" || platform == "wp" {
		link = fmt.Sprintf("https://zh.wikipedia.org/wiki/%s", keyword)
	}
	if link == "" {
		return fmt.Errorf("invalid platform")
	}
	goos := runtime.GOOS
	opencmd := "open"
	opencmd, ok := openCmds[goos]
	if !ok {
		return fmt.Errorf("can not open link in %s", goos)
	}
	if err := exec.Command(opencmd, link).Start(); err != nil {
		return err
	}

	return nil
}
