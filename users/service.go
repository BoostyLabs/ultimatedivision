// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package users

import "ultimatedivision/database"

type Service struct {
	db *database.UsersRepository
}

func(service *Service) GetUserByID(userID string) (*User, error){
	user, err := service.db.GetById(userID)

	if err != nil {
		return nil, err
	}

	return user, nil
}