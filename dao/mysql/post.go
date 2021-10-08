package mysql

import (
	"bluebell/models"
	"database/sql"
	"strings"

	"go.uber.org/zap"

	"github.com/jmoiron/sqlx"
)

func CreatePost(p *models.Post) (err error) {
	sqlStr := "insert into post (post_id,title,content,author_id,community_id) values (?,?,?,?,?)"
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	if err != nil {
		return err
	}
	return
}

// GetPostByID 根据pid获取帖子的详情
func GetPostByID(pid int64) (p *models.Post, err error) {
	p = new(models.Post)
	//使用select * 不行,部分字段没有，会报missing destination name id in
	//sqlStr := `select * from post where post_id = ?`
	sqlStr := `select
	post_id, title, content, author_id, community_id, create_time
	from post
	where post_id = ?
	`
	err = db.Get(p, sqlStr, pid)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrorCommunityIsNull
		}
		return nil, err
	}
	return p, nil
}

func GetPostList(page, size int64) (ps []*models.Post, err error) {
	//ps = new([]models.Post) 使用指针后面要转来转去  []*models.Post为指针数组， *[]model.Post为数组指针 Community使用的是*[]model.Community为数组指针
	ps = make([]*models.Post, 0, 2)
	sqlStr := `select
	post_id, title, content, author_id, community_id, create_time
	from post
	    order by create_time
	    DESC 
	limit ?,?
	`
	err = db.Select(&ps, sqlStr, (page-1)*size, size)
	//zap.L().Info("post集合", zap.Any("ps", ps))
	if err != nil {
		return nil, err
	}
	return
}

// GetPostList 查询帖子列表函数
//func GetPostList(page, size int64) (posts []*models.Post, err error) {
//	sqlStr := `select
//	post_id, title, content, author_id, community_id, create_time
//	from post
//	ORDER BY create_time
//	DESC
//	limit ?,?
//	`
//	posts = make([]*models.Post, 0, 2) // 不要写成make([]*models.Post, 2)
//	err = db.Select(&posts, sqlStr, (page-1)*size, size)
//	return
//}
//跟

// GetPostListByIDs 根据给定的pid列表查询贴详情
func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id ,community_id, create_time
				from post
				where post_id in (?)
				order by FIND_IN_SET(post_id,?)
				`
	// 每一条post_id使用","隔开
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	zap.L().Debug("query", zap.Any("query", query))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)
	return
}
