package main

import (
	"reflect"
	"time"
)

func CheckChanges() {
	for true {
		users := make(map[uint64]*User)
		deepCopyMap(users, GetUsersCache())

		time.Sleep(10 * time.Second)

		new_users := make(map[uint64]*User)
		deepCopyMap(new_users, GetUsersCache())

		if !reflect.DeepEqual(users, new_users) {
			for _, one_user := range new_users {
				updateUser(one_user)
			}
		}
	}
}

func deepCopyMap(new_container map[uint64]*User, users_data map[uint64]*User) {
	for key, value := range users_data {
		new_container[key] = value
	}
}
