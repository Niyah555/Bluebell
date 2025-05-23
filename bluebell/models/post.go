package models

import "time"

type Post struct {
	ID          int64     `json:"id,string" db:"post_id"`                            // 帖子ID
	AuthorID    int64     `json:"author_id" db:"author_id" binding:"required"`       // 作者ID
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"` // 社区ID
	Status      int32     `json:"status" db:"status"`                                // 帖子状态
	Title       string    `json:"title" db:"post_title" binding:"required"`          // 帖子标题
	Content     string    `json:"content" db:"content" binding:"required"`           // 帖子内容
	CreateTime  time.Time `json:"create_time" db:"create_time"`                      // 创建时间
}

//ApiPostDetail 帖子详情接口的结构体

type ApiPostDetail struct {
	AuthorName       string `json:"author_name"` // 作者名称
	VoteNum          int64  `json:"vote_num"`    // 点赞数
	*Post                                        // 嵌入帖子结构体
	*CommunityDetail `json:"community"`          // 嵌入社区信息
}
