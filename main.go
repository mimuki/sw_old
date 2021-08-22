package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/starshine-sys/pkgo"

	flag "github.com/spf13/pflag"
)

type commands struct {
	token   string
	session *pkgo.Session
}

var c = &commands{}

func init() {
	flag.StringVarP(&c.token, "token", "t", "", "PluralKit API token")
	flag.Parse()
	os.Args = flag.Args()
}

func main() {
	var err error

	if c.token == "" {
		if os.Getenv("PK_TOKEN") != "" {
			c.token = os.Getenv("PK_TOKEN")
		} else {
			home, err := os.UserHomeDir()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: could not get home directory: %v\n", err)
				os.Exit(1)
			}
			if _, err := os.Stat(path.Join(home, ".pktoken")); os.IsNotExist(err) {
				fmt.Fprintln(os.Stderr, path.Join(home, ".pktoken"), "does not exist.")
				os.Exit(1)
			} else {
				tokenFile, err := ioutil.ReadFile(path.Join(home, ".pktoken"))
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error reading token file: %v\n", err)
					os.Exit(1)
				}
				c.token = strings.TrimSpace(string(tokenFile))
			}
		}
	}

	c.session = pkgo.New(c.token)

	if len(os.Args) == 0 {
		err = c.current()
	} else {
		switch os.Args[0] {
		case "current":
			err = c.current()
		case "list":
			err = c.list()
		case "out":
			err = c.cmdSwitch()
		default:
			err = c.cmdSwitch(os.Args...)
		}
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
