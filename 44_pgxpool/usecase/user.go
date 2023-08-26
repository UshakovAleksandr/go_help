package usecase

import (
	"fmt"

	ruauka "github.com/ruauka/attrs-go"

	"44_pgxpool/model"
)

func (u *Usecases) AddUsers() {
	users := []model.User{
		{
			Username: "user1",
			Password: "password1",
		},
		{
			Username: "user2",
			Password: "password2",
		},
	}
	for _, user := range users {
		id, err := u.usecases.User.Add(&user)
		if err != nil {
			fmt.Println("func - AddUsers. ERROR by adding new user: ", err)
			return
		}

		fmt.Printf("added new user: id=%d\n", id)
	}
}

func (u *Usecases) GetUser() {
	user, err := u.usecases.User.GetById(1)
	if err != nil {
		fmt.Println("func - GetUserById. ERROR by getting user:", err)
		return
	}

	fmt.Printf("get user: %#v\n", user)
}

func (u *Usecases) GetAllUsers() {
	users, err := u.usecases.User.GetAll()
	if err != nil {
		fmt.Println("func - GetAllUsers. ERROR by getting users:", err)
		return
	}

	fmt.Printf("get users: %#v\n", users)
}

func (u *Usecases) UpdateUser() {
	var (
		userID       = 2
		userName     = "new_user2"
		userPassword = "new_password2"
	)

	newUser := model.UserUpdate{
		Username: &userName,
		Password: &userPassword,
	}

	if IsEmpty(newUser, model.UserUpdate{}) {
		fmt.Println("no fields to update")
		return
	}

	user, err := u.usecases.User.GetById(userID)
	if err != nil {
		fmt.Println("func - UpdateUser. ERROR by updating user:", err)
	}

	if err := ruauka.SetStructAttrs(&user, newUser); err != nil {
		fmt.Println("func - UpdateUser. ERROR by updating user:", err)
		return
	}

	if err := u.usecases.User.Update(user); err != nil {
		fmt.Println("func - UpdateUser. ERROR by updating user:", err)
		return
	}

	fmt.Printf("update user: %#v\n", user)
}

func (u *Usecases) DeleteUser() {
	userID := 1
	err := u.usecases.User.DeleteById(userID)
	if err != nil {
		fmt.Println("func - DeleteUser. ERROR by deleting user:", err)
		return
	}

	fmt.Printf("deleted user: id=%d\n", userID)
}
