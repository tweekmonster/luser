// +build go1.7

package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/tweekmonster/luser"
)

var lookupGroup = flag.Bool("g", false, "Lookup group")

func printuser(u *luser.User) {
	gids, _ := u.GroupIds()
	fmt.Printf(userStr, u.Username, u.Uid, u.Gid, u.Name, u.HomeDir, u.IsLuser,
		strings.Join(gids, ", "))
}

func printgroup(g *luser.Group) {
	fmt.Printf(groupStr, g.Name, g.Gid, g.IsLuser)
}

func checkgroup(arg string) {
	if g, err := luser.LookupGroupId(arg); err == nil {
		printgroup(g)
	}

	if g, err := luser.LookupGroup(arg); err == nil {
		printgroup(g)
	}
}
