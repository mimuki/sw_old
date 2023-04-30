package main

import (
  "fmt"
  "sort"
  "strings"
)

func (c *commands) list() error {
  sys, err := c.session.Me()
  if err != nil {
    return fmt.Errorf("getting system: %w", err)
  }

  members, err := c.session.Members(sys.ID)
  if err != nil {
    return fmt.Errorf("getting member list: %w", err)
  }

  name := sys.Name
  if name == "" {
    name = sys.ID
  }

  sort.Slice(members, func(i, j int) bool {
    return members[i].Name < members[j].Name
  })

  fmt.Printf("Members of %v:\n", name)
  var b strings.Builder
  for _, m := range members {
    b.WriteString(fmt.Sprintf("%v [%v], ", m.Name, m.ID))
    if b.Len() >= 80 {
      fmt.Println(b.String())
      b.Reset()
    }
  }
  return nil
}
