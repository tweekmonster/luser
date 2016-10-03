// +build !windows

package luser

import "strconv"

// Lookup a user by username.
func lookupUser(username string) (*User, error) {
	if _, err := strconv.Atoi(username); err == nil {
		return nil, UnknownUserError(username)
	}

	if dscacheutilExe != "" {
		u, err := dsUser(username)
		if err == nil {
			return luser(u), nil
		}
	}

	u, err := getentUser(username)
	if err == nil {
		return luser(u), nil
	}

	return nil, UnknownUserError(username)
}

// Lookup user by UID.
func lookupId(uid string) (*User, error) {
	id, err := strconv.Atoi(uid)
	if err != nil {
		return nil, err
	}

	if dscacheutilExe != "" {
		u, err := dsUserId(uid)
		if err == nil {
			return luser(u), nil
		}
	}

	u, err := getentUser(uid)
	if err == nil {
		return luser(u), nil
	}

	return nil, UnknownUserIdError(id)
}

func lookupGroup(name string) (*Group, error) {
	if _, err := strconv.Atoi(name); err == nil {
		return nil, UnknownGroupError(name)
	}

	if dscacheutilExe != "" {
		g, err := dsGroup(name)
		if err == nil {
			return lgroup(g), nil
		}
	}

	g, err := getentGroup(name)
	if err == nil {
		return lgroup(g), nil
	}

	return nil, UnknownGroupError(name)
}

func lookupGroupId(gid string) (*Group, error) {
	id, err := strconv.Atoi(gid)
	if err != nil {
		return nil, err
	}

	if dscacheutilExe != "" {
		g, err := dsGroupId(gid)
		if err == nil {
			return lgroup(g), nil
		}
	}

	g, err := getentGroup(gid)
	if err == nil {
		return lgroup(g), nil
	}

	return nil, UnknownGroupIdError(id)
}

func (u *User) lookupUserGroupIds() ([]string, error) {
	if idExe != "" {
		return idGroupList(u)
	}

	return nil, ErrListGroups
}
