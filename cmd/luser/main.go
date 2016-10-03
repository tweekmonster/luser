// Command luser is a sample program for looking up users or groups.
package main

import (
	"flag"
	"fmt"
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
	currentUser := flag.Bool("c", false, "Lookup current user")
	lookupGroup := flag.Bool("g", false, "Lookup group")
	flag.Parse()

	if *currentUser {
		if u, err := luser.Current(); err == nil {
			if *lookupGroup {
				if g, err := luser.LookupGroupId(u.Gid); err == nil {
					printgroup(g)
				}
			} else {
				printuser(u)
			}
		}
		return
	}

	for _, arg := range flag.Args() {
		if *lookupGroup {
			if g, err := luser.LookupGroupId(arg); err == nil {
				printgroup(g)
			}

			if g, err := luser.LookupGroup(arg); err == nil {
				printgroup(g)
			}
		} else {
			if u, err := luser.LookupId(arg); err == nil {
				printuser(u)
			}

			if u, err := luser.Lookup(arg); err == nil {
				printuser(u)
			}
		}
	}
}
