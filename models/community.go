package models

type Community struct {
	ID            int    `json:"id" db:"id"`
	CommunityID   int    `json:"community_id" db:"community_id"`
	CommunityName string `json:"community_name" db:"community_name"`
	Introduction  string `json:"introduction" db:"introduction"`
	//CreateTime    time.Time `json:"create_time"`
	//UpdateTime    time.Time `json:"update_time"`
}
