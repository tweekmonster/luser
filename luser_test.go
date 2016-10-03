// This is a copy of os/user/user_test.go with some modifications.  These tests
// will only run if CGO_ENABLED=0 and the test target isn't Windows.

// +build !cgo
// +build !windows

package luser

import (
	"runtime"
	"testing"
)

func checkLuserUser(t *testing.T, u *User) {
	if !u.IsLuser {
		t.Fatalf("luser: User was found without a fallback method")
	}
}

func TestFallback(t *testing.T) {
	if !fallbackEnabled {
		t.Fatalf("fallbackEnabled is not true")
	}
}

func TestCurrent(t *testing.T) {
	u, err := Current()
	if err != nil {
		t.Fatalf("Current: %v (got %#v)", err, u)
	}

	checkLuserUser(t, u)

	if u.HomeDir == "" {
		t.Errorf("didn't get a HomeDir")
	}
	if u.Username == "" {
		t.Errorf("didn't get a username")
	}
}

func compare(t *testing.T, want, got *User) {
	if want.Uid != got.Uid {
		t.Errorf("got Uid=%q; want %q", got.Uid, want.Uid)
	}
	if want.Username != got.Username {
		t.Errorf("got Username=%q; want %q", got.Username, want.Username)
	}
	if want.Name != got.Name {
		t.Errorf("got Name=%q; want %q", got.Name, want.Name)
	}
	// TODO(brainman): fix it once we know how.
	if runtime.GOOS == "windows" {
		t.Skip("skipping Gid and HomeDir comparisons")
	}
	if want.Gid != got.Gid {
		t.Errorf("got Gid=%q; want %q", got.Gid, want.Gid)
	}
	if want.HomeDir != got.HomeDir {
		t.Errorf("got HomeDir=%q; want %q", got.HomeDir, want.HomeDir)
	}
}

func TestLookup(t *testing.T) {
	want, err := Current()
	if err != nil {
		t.Fatalf("Current: %v", err)
	}

	checkLuserUser(t, want)

	got, err := Lookup(want.Username)
	if err != nil {
		t.Fatalf("Lookup: %v", err)
	}

	checkLuserUser(t, got)

	compare(t, want, got)
}

func TestLookupId(t *testing.T) {
	want, err := Current()
	if err != nil {
		t.Fatalf("Current: %v", err)
	}

	checkLuserUser(t, want)

	got, err := LookupId(want.Uid)
	if err != nil {
		t.Fatalf("LookupId: %v", err)
	}

	checkLuserUser(t, got)

	compare(t, want, got)
}
