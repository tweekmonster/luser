// Command luser is a sample program for looking up users or groups.

package main

import (
	"flag"
	"fmt"

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

var currentUser = flag.Bool("c", false, "Lookup current user")

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: luser [options] args...")
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
	}

	if *currentUser {
		if u, err := luser.Current(); err == nil {
			if lookupGroup != nil && *lookupGroup {
				checkgroup(u.Gid)
			} else {
				printuser(u)
			}
		}
		return
	}

	for _, arg := range flag.Args() {
		if lookupGroup != nil && *lookupGroup {
			checkgroup(arg)
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
