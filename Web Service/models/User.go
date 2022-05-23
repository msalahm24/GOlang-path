package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextid = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("the id of the new user must be empty")
	}
	u.ID = nextid
	nextid++
	users = append(users, &u)
	return u, nil
}

func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("The user of id '%v' not found",id)
}

func UpdateUser(u User)(User,error){
	for ind,candinate := range users{
		if candinate.ID == u.ID{
			users[ind] = &u
			return u , nil 
		}
	}
	return User{},fmt.Errorf("this user of id '%v' not found ",u.ID)
}

func RemoveUser(id int) error{
	for i , u := range users {
		if u.ID == id {
			users = append(users[:i],users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("the User with id '%v' not found",id)
}
