package models

import (
	"context"
	"issue-maintainer/db"
)

type Group struct {
	Id          int64
	Title       string `binding:"required"`
	Description string `binding:"required"`
	UserId      int64
}

func (group Group) Save() error {
	query := `INSERT INTO groups (title, description, user_id) VALUES ($1, $2, $3)`
	_, err := db.DB.Exec(context.Background(), query, group.Title, group.Description, group.UserId)
	return err
}

func GetGroupById(id int64) (Group, error) {
	query := `SELECT * FROM groups WHERE id = $1`
	row := db.DB.QueryRow(context.Background(), query, id)

	var group Group
	err := row.Scan(&group.Id, &group.Title, &group.Description, &group.UserId)

	return group, err
}

func DeleteGroupById(id int64) error {
	query := `DELETE FROM groups WHERE id = $1`
	_, err := db.DB.Exec(context.Background(), query, id)
	return err
}
