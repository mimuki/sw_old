package main

import (
	"fmt"
	"os"
	"strings"
)

func (c *commands) cmdSwitch(args ...string) {
	var in string

	members := make([]string, 0)
	memberNames := make([]string, 0)
	if len(args) == 0 {
		members = []string{}
		memberNames = []string{"no fronter"}
	} else {
		m, err := c.session.GetMembers("")
		if err != nil {
			fmt.Println("Error getting members:", err)
			os.Exit(1)
		}

		names, _ := m.ToMaps()

		for _, name := range args {
			for n, id := range names {
				if strings.ToLower(name) == n {
					members = append(members, id)
					memberNames = append(memberNames, n)
					break
				} else if strings.ToLower(name) == id {
					members = append(members, id)
					memberNames = append(memberNames, n)
					break
				}
			}
		}
	}

	fmt.Printf("Are you sure you want to switch in the following members?\n%v\n", strings.Join(memberNames, ", "))
	_, err := fmt.Scanln(&in)
	if err != nil {
		fmt.Println("Error getting input:", err)
		os.Exit(1)
	}

	switch in {
	case "yes", "y":
		err = c.session.RegisterSwitch(members...)
		if err != nil {
			fmt.Printf("Error switching [%v] in: %v\n", strings.Join(memberNames, ", "), err)
			os.Exit(1)
		}
		fmt.Printf("Switch registered. Current fronters are now %v.\n", strings.Join(memberNames, ", "))
		return
	default:
		fmt.Println("Switch aborted.")
		return
	}
}
