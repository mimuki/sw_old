package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/Starshine113/pkgo"

	flag "github.com/spf13/pflag"
)

type commands struct {
	token   string
	session *pkgo.Session
	sys     *pkgo.System
}

var c = &commands{}

func init() {
	flag.StringVarP(&c.token, "token", "t", "", "PluralKit API token")
	flag.Parse()
}

func main() {
	if c.token == "" {
		if os.Getenv("PK_TOKEN") != "" {
			c.token = os.Getenv("PK_TOKEN")
		} else {
			home, err := os.UserHomeDir()
			if err != nil {
				fmt.Println("Error: could not get home directory:", err)
				os.Exit(1)
			}
			if _, err := os.Stat(path.Join(home, ".pktoken")); os.IsNotExist(err) {
				fmt.Println(path.Join(home, ".pktoken"), "does not exist.")
				os.Exit(1)
			} else {
				tokenFile, err := ioutil.ReadFile(path.Join(home, ".pktoken"))
				if err != nil {
					fmt.Println("Error reading token file:", err)
				}
				c.token = strings.TrimSpace(string(tokenFile))
			}
		}
	}

	s := pkgo.NewSession(&pkgo.Config{Token: c.token})
	c.session = s

	sys, err := s.GetSystem()
	if err != nil {
		fmt.Println("Error getting system info:", err)
		os.Exit(1)
	}
	c.sys = sys

	if len(os.Args) == 1 {
		c.current()
		return
	}

	switch os.Args[1] {
	case "current":
		c.current()
		return
	case "out":
		c.cmdSwitch()
		return
	default:
		c.cmdSwitch(os.Args[1:]...)
		return
	}
}
