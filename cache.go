package jse_test

var usersCache map[uint64]*User

func init() {
	usersCache = make(map[uint64]*User)
}

func isIdUnique(id uint64) bool {
	if usersCache[id] == nil {
		return true
	}

	return false
}

func addNewUserToBuffer(new_user User) {
	usersCache[new_user.Id] = &new_user
}
