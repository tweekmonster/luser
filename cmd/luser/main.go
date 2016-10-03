// Command luser is a sample program for looking up users or groups.
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/tweekmonster/luser"
)

var userStr = `
User:     %q
  Uid:    %q
  Gid:    %q
  Name:   %q
  Home:   %q
  Luser:  %t
  Groups: %v
`

var groupStr = `
Group:   %q
  Gid:   %q
  Luser: %t
`

func printuser(u *luser.User) {
	gids, _ := u.GroupIds()
	fmt.Printf(userStr, u.Username, u.Uid, u.Gid, u.Name, u.HomeDir, u.IsLuser,
		strings.Join(gids, ", "))
}

func printgroup(g *luser.Group) {
	fmt.Printf(groupStr, g.Name, g.Gid, g.IsLuser)
}

func main() {
	for _, arg := range os.Args[1:] {
		if _, err := strconv.Atoi(arg); err == nil {
			u, err := luser.LookupId(arg)
			if err == nil {
				printuser(u)
			}

			if g, err := luser.LookupGroupId(arg); err == nil {
				printgroup(g)
			}
		} else {
			if u, err := luser.Lookup(arg); err == nil {
				printuser(u)
			}

			if g, err := luser.LookupGroup(arg); err == nil {
				printgroup(g)
			}
		}
	}
}
