package pack

import (
	"github.com/1037group/dousheng/kitex_gen/douyin_comment"
	"github.com/1037group/dousheng/kitex_gen/douyin_user"
	"github.com/1037group/dousheng/pkg/configs/sql"
)

// Comment pack comment info
func Comment(m *sql.Comment, n *douyin_user.User) *douyin_comment.Comment {
	if m == nil {
		return nil
	}

	return &douyin_comment.Comment{
		Id:         m.CommentId,
		User:       n,
		Content:    m.CommentContent,
		CreateDate: m.Utime.String(), //is this right?
	}
}

// Comments pack list of comments info
func Comments(ms []*sql.Comment, ns map[int64]douyin_user.User) []*douyin_comment.Comment {
	comments := make([]*douyin_comment.Comment, 0)
	for _, m := range ms {
		user := ns[m.UserId]
		if comment := Comment(m, &user); comment != nil {
			comments = append(comments, comment)
		}
	}
	return comments
}
