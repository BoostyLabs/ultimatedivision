// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package users

import (
	gouuid "github.com/google/uuid"
	"time"
)

type User struct {
	ID        gouuid.UUID `json:"id"`
	Email     string      `json:"email"`
	Password  string      `json:"password"`
	NickName  string      `json:"nick_name"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	LastLogin string      `json:"last_login"`
	Status    string      `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
}

func NewUser(email, password, nickName, firstName, lastName, lastLogin, status string) *User {
	u, _ := gouuid.NewUUID()
	return &User{
		ID:        u,
		Email:     email,
		Password:  password,
		NickName:  nickName,
		FirstName: firstName,
		LastName:  lastName,
		LastLogin: lastLogin,
		Status:    status,
	}
}

//func (user *User) EncodePass() error {
//	hash, err :=
//}
