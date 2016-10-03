// +build !cgo
// +build !windows
// +build go1.7

package luser

import "testing"

func checkLuserGroup(t *testing.T, u *Group) {
	if !u.IsLuser {
		t.Fatalf("luser: User was found without a fallback method")
	}
}

func TestLookupGroup(t *testing.T) {
	user, err := Current()
	if err != nil {
		t.Fatalf("Current(): %v", err)
	}

	checkLuserUser(t, user)

	g1, err := LookupGroupId(user.Gid)
	if err != nil {
		// NOTE(rsc): Maybe the group isn't defined. That's fine.
		// On my OS X laptop, rsc logs in with group 5000 even
		// though there's no name for group 5000. Such is Unix.
		t.Logf("LookupGroupId(%q): %v", user.Gid, err)
		return
	}

	checkLuserGroup(t, g1)

	if g1.Gid != user.Gid {
		t.Errorf("LookupGroupId(%q).Gid = %s; want %s", user.Gid, g1.Gid, user.Gid)
	}

	g2, err := LookupGroup(g1.Name)
	if err != nil {
		t.Fatalf("LookupGroup(%q): %v", g1.Name, err)
	}

	checkLuserGroup(t, g2)

	if g1.Gid != g2.Gid || g1.Name != g2.Name {
		t.Errorf("LookupGroup(%q) = %+v; want %+v", g1.Name, g2, g1)
	}
}

func TestGroupIds(t *testing.T) {
	user, err := Current()
	if err != nil {
		t.Fatalf("Current(): %v", err)
	}

	checkLuserUser(t, user)

	gids, err := user.GroupIds()
	if err != nil {
		t.Fatalf("%+v.GroupIds(): %v", user, err)
	}
	if !containsID(gids, user.Gid) {
		t.Errorf("%+v.GroupIds() = %v; does not contain user GID %s", user, gids, user.Gid)
	}
}

func containsID(ids []string, id string) bool {
	for _, x := range ids {
		if x == id {
			return true
		}
	}
	return false
}
