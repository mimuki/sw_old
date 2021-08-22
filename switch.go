package main

import (
	"fmt"
	"strings"
)

func (c *commands) cmdSwitch(args ...string) error {
	var in string

	var toSwitch, names []string
	if len(args) == 0 {
		toSwitch = []string{}
		names = []string{"(no fronter)"}
	} else {
		sys, err := c.session.Me(false)
		if err != nil {
			return fmt.Errorf("getting system: %w", err)
		}

		members, err := c.session.Members(sys.ID)
		if err != nil {
			return fmt.Errorf("getting members: %w", err)
		}

		for _, name := range args {
			found := false
			for _, m := range members {
				if strings.EqualFold(name, m.Name) {
					found = true
					toSwitch = append(toSwitch, m.ID)
					names = append(names, m.Name)
					break
				} else if strings.EqualFold(name, m.ID) {
					found = true
					toSwitch = append(toSwitch, m.ID)
					names = append(names, m.Name)
					break
				}
			}
			if !found {
				return fmt.Errorf("no member named \"%v\" found. Note that a member ID is 5 characters long", name)
			}
		}
	}

	fmt.Printf("Are you sure you want to switch in the following members?\n%v\n", strings.Join(names, ", "))
	_, err := fmt.Scanln(&in)
	if err != nil {
		return fmt.Errorf("error getting input: %w", err)
	}

	switch in {
	case "yes", "y":
		err = c.session.RegisterSwitch(toSwitch...)
		if err != nil {
			return fmt.Errorf("error switching [%v] in: %w", strings.Join(names, ", "), err)
		}
		fmt.Printf("Switch registered. Current fronters are now %v.\n", strings.Join(names, ", "))
		return nil
	default:
		fmt.Println("Switch aborted.")
		return nil
	}
}
