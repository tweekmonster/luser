// Package luser is a drop-in replacement for 'os/user' which allows you to
// lookup users and groups in cross-compiled builds without 'cgo'.
//
// 'os/user' requires 'cgo' to lookup users using the target OS's API.  This is
// the most reliable way to look up user and group information.  However,
// cross-compiling means that 'os/user' will only work for the OS you're using.
// 'user.Current()' is usable when building without 'cgo', but doesn't always
// work.  The '$USER' and '$HOME' variables could be different from what you
// expect or not even exist.
//
// If you want to cross-compile a relatively simple program that needs to write
// a config file somewhere in the user's directory, the last thing you want to
// do is figure out some elaborate build scheme involving virtual machines.
//
// When cgo is not available for a build, one of the following methods will be
// used to lookup user and group information:
//
//  | Method        | Used for                                                       |
//  |---------------|----------------------------------------------------------------|
//  | `/etc/passwd` | Parsed to lookup user information. (Unix, Linux)               |
//  | `/etc/group`  | Parsed to lookup group information. (Unix, Linux)              |
//  | `getent`      | Optional. Find user/group information. (Unix, Linux)           |
//  | `dscacheutil` | Lookup user/group information via Directory Services. (Darwin) |
//  | `id`          | Finding a user's groups when using `GroupIds()`.               |
//
// You should be able to simply replace 'user.' with 'luser.' (in most cases).
package luser

import (
	"os/user"
	"strings"
)

// Tests if the error was created because cgo wasn't enabled.
func isCgoErr(err error) bool {
	return strings.HasSuffix(err.Error(), "requires cgo")
}

// Group represents a grouping of users.  Embedded *user.Group reference:
// https://golang.org/pkg/os/user/#Group
type Group struct {
	*user.Group

	// IsLuser is a flag indicating if the user was found without cgo.
	IsLuser bool
}

// LookupGroup looks up a group by name. If the group cannot be found, the
// returned error is of type UnknownGroupError.
func LookupGroup(name string) (*Group, error) {
	g, err := user.LookupGroup(name)
	if err == nil {
		return &Group{Group: g}, err
	}

	if isCgoErr(err) {
		return lookupGroup(name)
	}

	return nil, err
}

// LookupGroupId looks up a group by groupid. If the group cannot be found, the
// returned error is of type UnknownGroupIdError.
func LookupGroupId(gid string) (*Group, error) {
	g, err := user.LookupGroupId(gid)
	if err == nil {
		return &Group{Group: g}, err
	}

	if isCgoErr(err) {
		return lookupGroupId(gid)
	}

	return nil, err
}

// User represents a user account.  Embedded *user.User reference:
// https://golang.org/pkg/os/user/#User
type User struct {
	*user.User

	IsLuser bool // flag indicating if the user was found without cgo.
}

// GroupIds returns the list of group IDs that the user is a member of.
func (u *User) GroupIds() ([]string, error) {
	if u.IsLuser {
		return u.lookupUserGroupIds()
	}
	return u.User.GroupIds()
}

// Current returns the current user.  On builds where cgo is available, this
// returns the result from user.Current().  Otherwise, alternate lookup methods
// are used before falling back to the built-in stub.
func Current() (*User, error) {
	return currentUser()
}

// Lookup looks up a user by username. If the user cannot be found, the
// returned error is of type UnknownUserError.
func Lookup(username string) (*User, error) {
	u, err := user.Lookup(username)
	if err == nil {
		return &User{User: u}, nil
	}

	if isCgoErr(err) {
		return lookupUser(username)
	}

	return nil, err
}

// LookupId looks up a user by userid. If the user cannot be found, the
// returned error is of type UnknownUserIdError.
func LookupId(uid string) (*User, error) {
	u, err := user.LookupId(uid)
	if err == nil {
		return &User{User: u}, nil
	}

	if isCgoErr(err) {
		return lookupId(uid)
	}

	return nil, err
}
