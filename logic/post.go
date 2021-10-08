package logic

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/models"
	"bluebell/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	//使用雪花算法生成ID
	p.ID = snowflake.GenID()

	err = mysql.CreatePost(p)
	if err != nil {
		zap.L().Error("mysql.CreatePost fail", zap.Error(err))
		return err
	}
	err = redis.CreatePost(p.ID, p.CommunityID)
	return
}

func GetPost(pid int64) (p *models.ApiPostDetail, err error) {

	post, err := mysql.GetPostByID(pid)
	if err != nil {
		return nil, err
	}
	//传入作者名称
	user, err := mysql.GetUserById(post.AuthorID)
	if err != nil {
		return nil, err
	}
	//传入社区名称
	community, err := mysql.GetCommunityDetailById(post.CommunityID)
	if err != nil {
		return nil, err
	}
	//数据拼接
	data := &models.ApiPostDetail{
		AuthorName: user.Username,
		Post:       post,
		Community:  community,
	}
	return data, nil
}

// GetPostList 分页获取帖子
func GetPostList(page int64, size int64) (aps []*models.ApiPostDetail, err error) {
	//查询返回全部帖子
	postList, err := mysql.GetPostList(page, size)
	if err != nil {
		return nil, err
	}
	dataList := make([]*models.ApiPostDetail, 0, len(postList))
	for _, post := range postList {

		//传入作者名称
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			continue
		}
		//传入社区名称
		community, err := mysql.GetCommunityDetailById(post.CommunityID)
		if err != nil {
			continue
		}
		//数据拼接
		data := &models.ApiPostDetail{
			AuthorName: user.Username,
			Community:  community,
			Post:       post,
		}
		dataList = append(dataList, data)
	}
	//对每一个帖子进行数据拼接
	//声明切片 循环处理数据 进行拼接
	//return
	return dataList, nil
}

// GetPostList2  根据传入order返回
func GetPostList2(p *models.ParamPostList) (aps []*models.ApiPostDetail, err error) {

	//2.去redis查询id列表
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		return
	}
	//如何返回post_list为data为0行直接return
	if len(ids) == 0 {
		zap.L().Warn("postlist data is 0")
		return
	}

	//3.根据id查询帖子详细信息
	//sql中使用 FIND_IN_SET函数 ，所以返回的数据根据传入顺序从redis中的返回
	postList, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}

	//提前查询好pids投票数据避免重复查询
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}

	//重复 GetPostList
	dataList := make([]*models.ApiPostDetail, 0, len(postList))
	for idx, post := range postList {
		//zap.L().Debug("voteData_idx", zap.Int64("idx", voteData[idx]))
		//传入作者名称
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			continue
		}
		//传入社区名称
		community, err := mysql.GetCommunityDetailById(post.CommunityID)
		if err != nil {
			continue
		}
		//数据拼接
		data := &models.ApiPostDetail{
			AuthorName: user.Username,
			VoteNum:    voteData[idx],
			Community:  community,
			Post:       post,
		}
		dataList = append(dataList, data)
	}
	//对每一个帖子进行数据拼接
	//声明切片 循环处理数据 进行拼接
	//return
	return dataList, nil
}

// GetPostListByCommunity  根据传入order和community返回
func GetPostListByCommunity(p *models.ParamPostList) (aps []*models.ApiPostDetail, err error) {
	//2.去redis查询id列表
	ids, err := redis.GetCommunityPostIDsInOrder(p)
	if err != nil {
		return
	}
	//如何返回post_list为data为0行直接return
	if len(ids) == 0 {
		zap.L().Warn("postlist data is 0")
		return
	}

	//3.根据id查询帖子详细信息
	//sql中使用 FIND_IN_SET函数 ，所以返回的数据根据传入顺序从redis中的返回
	postList, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}

	//提前查询好pids投票数据避免重复查询
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}

	//重复 GetPostList
	dataList := make([]*models.ApiPostDetail, 0, len(postList))
	for idx, post := range postList {
		//zap.L().Debug("voteData_idx", zap.Int64("idx", voteData[idx]))
		//传入作者名称
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			continue
		}
		//传入社区名称
		community, err := mysql.GetCommunityDetailById(post.CommunityID)
		if err != nil {
			continue
		}
		//数据拼接
		data := &models.ApiPostDetail{
			AuthorName: user.Username,
			VoteNum:    voteData[idx],
			Community:  community,
			Post:       post,
		}
		dataList = append(dataList, data)
	}
	//对每一个帖子进行数据拼接
	//声明切片 循环处理数据 进行拼接
	//return
	return dataList, nil
}

// GetPostListNew 两个查询接口合二为一
func GetPostListNew(p *models.ParamPostList) (aps []*models.ApiPostDetail, err error) {

	if p.CommunityID == 0 {
		aps, err = GetPostList2(p)
	} else {
		aps, err = GetPostListByCommunity(p)
	}
	if err != nil {
		zap.L().Error("GetPostListNew fail", zap.Error(err))
		return nil, err
	}
	return
}
