package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

func ListCommunity() (*[]models.Community, error) {
	return mysql.ListCommunity()
}

func CommunityDetail(id int64) (*models.Community, error) {
	return mysql.GetCommunityDetailById(id)
}
