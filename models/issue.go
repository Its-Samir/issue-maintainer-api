package models

type Issue struct {
	Id               int64
	Title            string `binding:"required"`
	Description      string `binding:"required"`
	Status           string
	UserId           int64
	AssignedToUserId int64
	GroupId          int64
}
