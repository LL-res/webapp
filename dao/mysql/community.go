package mysql

import (
	"database/sql"
	"webapp/models"

	"go.uber.org/zap"
)

func GetCommunityList() (list []*models.Community, err error) {
	sqlStr := "select community_id,community_name from community"
	if err = Db.Select(&list, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("no rows in db")
			err = nil
		}
	}
	return
}
func GetCommunityDetailByID(id int64) (detail models.CommunityDetail, err error) {
	sqlStr := "select community_id,community_name,introduction,create_time from community where id =? "
	if err = Db.Get(&detail, sqlStr, id); err != nil {
		return models.CommunityDetail{}, err
	} else {
		return
	}
}
