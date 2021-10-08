package models

const (
	OrderTime  = "time"
	OrderScore = "score"
)

type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamVoteData 投票数据
//type ParamVoteData struct {
//	//UserID中 从请求中获取当前的用户
//	PostID    string `json:"post_id,string" binding:"required"`       //帖子ID
//	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` //赞成票1 反对票-1 取消投票0
//}

// ParamVoteData 投票数据
type ParamVoteData struct {
	// UserID 从请求中获取当前的用户
	PostID    string `json:"post_id" binding:"required"`               // 贴子id
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1" ` // 赞成票(1)还是反对票(-1)取消投票(0)
}

// bluebell/models/params.go

// ParamPostList 获取帖子列表query string参数
type ParamPostList struct {
	CommunityID int64  `json:"community_id" form:"community_id"`   // 可以为空
	Page        int64  `json:"page" form:"page" example:"1"`       // 页码
	Size        int64  `json:"size" form:"size" example:"10"`      // 每页数据量
	Order       string `json:"order" form:"order" example:"score"` // 排序依据
}

//type ParamPostListByCommunity struct {
//	Page        int64  `json:"page" form:"page" example:"1"`
//	Size        int64  `json:"size" form:"size" example:"10"`
//	Order       string `json:"order" form:"order" `
//	CommunityID int64  `json:"community_id" form:"community_id"`
//}
