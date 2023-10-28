package storage

import (
	"context"
	"entities"
)

func (s *sqlserverStore) ListUser(ctx context.Context) ([]entities.UserModel, error) {
	var users []entities.UserModel

	s.db.Select(
		"UserID",
		"FullName",
		"Email",
		"Role").Table(entities.UserModelTable).Find(&users)
	if err := s.db.Table(entities.UserModelTable).Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}
