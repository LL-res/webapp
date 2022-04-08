package logic

import (
	"webapp/dao/mysql"
	"webapp/models"
)

func GetCommunityList() ([]*models.Community, error) {
	return mysql.GetCommunityList()
}
func GetCommunityDetail(id int64) (models.CommunityDetail, error) {
	detail, err := mysql.GetCommunityDetailByID(id)
	if err != nil {
		return models.CommunityDetail{}, err
	} else {
		return detail, nil
	}
}
