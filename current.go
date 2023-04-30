package main

import (
  "fmt"
  "strings"
)

func (c *commands) current() error {
  

  sys, err := c.session.Me()
  if err != nil {
    return fmt.Errorf("getting system: %w", err)
  }

  fronters, err := c.session.Fronters(sys.ID)
  if err != nil {
    return fmt.Errorf("getting fronters: %w", err)
  }

  var f []string
  for _, m := range fronters.Members {
    name := m.Name
    if m.DisplayName != "" {
      name = m.DisplayName
    }

    f = append(f, fmt.Sprintf("%v (%v)", name, m.ID))
  }
  if len(f) == 0 {
    f = []string{"(no fronter)"}
  }

  name := sys.Name
  if name == "" {
    name = sys.ID
  }
  fmt.Printf("Current fronters for %v:\n%v\n", name, strings.Join(f, "\n"))
  return nil
}
