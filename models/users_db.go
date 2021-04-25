package models

import (
	"fmt"
)

type User struct {
	Id   int
	Name string
}

func GetAllUsers() ([]User, error) {
	rows, err := db.Query("SELECT * FROm users")
	if err != nil {
		return nil, err
	}
	user_list := make([]User, 0)
	for rows.Next() {
		var user User
		if err = rows.Scan(user.Id, user.Name); err != nil {
			return nil, err
		}
		user_list = append(user_list, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return user_list, err
}

func GetUser(userName string) (*User, error) {
	var user User
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM users WHERE user_name=%s", userName))
	if err != nil {
		return nil, err
	}

	if err = rows.Scan(user.Id, user.Name); err != nil {
		return nil, err
	}
	return &user, err
}

func AddUser(userName string) error {
	_, err := db.Exec(fmt.Sprintf("INSERT INTO users(user_name) VALUE('%s')", userName))
	return err
}

func UpdateUser(oriUserName string, newUserName string) bool {
	_, err := db.Exec(fmt.Sprintf("UPDATE users SET user_name='%s' WHERE user_name=%s", oriUserName, newUserName))
	if err != nil {
		return false
	}
	return true
}
