// +build !go1.7

package main

import (
	"fmt"

	"github.com/tweekmonster/luser"
)

var lookupGroup *bool

func checkgroup(arg string) {}

func printuser(u *luser.User) {
	fmt.Printf(userStr, u.Username, u.Uid, u.Gid, u.Name, u.HomeDir, u.IsLuser)
}
