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

/* get all issues of a certain group */
func GetAllIssues(groupId int64) ([]Issue, error) {
	query := `SELECT * FROM issues WHERE group_id = $1`
	row, err := db.DB.Query(context.Background(), query, groupId)

	if err != nil {
		return nil, err
	}

	issues := []Issue{}

	for row.Next() {
		var issue Issue

		err = row.Scan(&issue.Id, &issue.Title, &issue.Description, &issue.Status, &issue.UserId, &issue.AssignedToUserId, &issue.GroupId)

		if err != nil {
			return nil, err
		}

		issues = append(issues, issue)
	}

	return issues, nil
}

func GetIssueById(id, groupId int64) (Issue, error) {
	query := `SELECT * FROM issues WHERE id = $1 AND group_id = $2`
	row := db.DB.QueryRow(context.Background(), query, id, groupId)

	var issue Issue
	err := row.Scan(&issue.Id, &issue.Title, &issue.Description, &issue.Status, &issue.UserId, &issue.AssignedToUserId, &issue.GroupId)

	if err != nil {
		return Issue{}, err
	}

	return issue, nil
}

/* get 5 new issues as featured */
func GetFeaturedIssues() ([]Issue, error) {
	query := `SELECT * FROM issues ORDER BY id DESC LIMIT $1`
	row, err := db.DB.Query(context.Background(), query, 5)

	if err != nil {
		return nil, err
	}

	issues := []Issue{}

	for row.Next() {
		var issue Issue

		err = row.Scan(&issue.Id, &issue.Title, &issue.Description, &issue.Status, &issue.UserId, &issue.AssignedToUserId, &issue.GroupId)

		if err != nil {
			return nil, err
		}

		issues = append(issues, issue)
	}

	return issues, nil
}

func UpdateIssueById(id int64, updatedIssue Issue) error {
	query := `UPDATE issues SET title = $1, description = $2, assigned_to_user_id = $3 WHERE id = $4`
	_, err := db.DB.Exec(context.Background(), query, updatedIssue.Title, updatedIssue.Description, updatedIssue.AssignedToUserId, id)
	return err
}

func DeleteIssueById(id int64) error {
	query := `DELETE FROM issues WHERE id = $1`
	_, err := db.DB.Exec(context.Background(), query, id)
	return err
}
