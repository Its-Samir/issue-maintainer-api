package models

type Group struct {
	Id          int64
	Title       string `binding:"required"`
	Description string `binding:"required"`
	UserId      int64
}
