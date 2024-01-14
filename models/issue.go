package models

import (
	"context"
	"issue-maintainer/db"
)

type Issue struct {
	Id               int64
	Title            string `binding:"required"`
	Description      string `binding:"required"`
	Status           string
	UserId           int64
	AssignedToUserId int64 `binding:"required"`
	GroupId          int64
}

func (issue Issue) Save() error {
	query := `INSERT INTO issues (title, description, user_id, group_id, assigned_to_user_id) VALUES ($1, $2, $3, $4, $5)`
	_, err := db.DB.Exec(context.Background(), query, issue.Title, issue.Description, issue.UserId, issue.GroupId, issue.AssignedToUserId)
	return err
}
