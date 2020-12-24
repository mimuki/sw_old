package main

import (
	"fmt"
	"os"
	"strings"
)

func (c *commands) current() {
	fronters, err := c.session.GetFronters("")
	if err != nil {
		fmt.Println("Error getting fronters:", err)
		os.Exit(1)
	}

	f := make([]string, 0)
	for _, m := range fronters.Members {
		f = append(f, fmt.Sprintf("%v (%v)", m.DisplayedName(), m.ID))
	}
	if len(f) == 0 {
		f = []string{"(no fronter)"}
	}

	name := c.sys.Name
	if name == "" {
		name = c.sys.ID
	}
	fmt.Printf("Current fronters for %v:\n%v\n", name, strings.Join(f, "\n"))
	return
}
