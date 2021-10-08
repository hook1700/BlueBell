package mysql

import (
	"bluebell/models"
	"database/sql"
	"errors"

	"go.uber.org/zap"
)

var ErrorCommunityIsNull = errors.New("种类类清单为空")

func ListCommunity() (cs *[]models.Community, err error) {
	cs = new([]models.Community)
	sqlStr := `select id,community_id,community_name,introduction from community `
	if err = db.Select(cs, sqlStr); err != nil {
		zap.L().Warn("数据库没有数据")
		if err == sql.ErrNoRows {
			return nil, ErrorCommunityIsNull
		}
		return nil, err
	}
	return cs, nil
}

func GetCommunityDetailById(id int64) (c *models.Community, err error) {
	c = new(models.Community)
	sqlStr := `select id,community_id,community_name,introduction from community where id =? `
	if err = db.Get(c, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrorCommunityIsNull
		}
		return nil, err
	}
	return c, nil
}
