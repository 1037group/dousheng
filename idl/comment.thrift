namespace go douyin_comment

include "user.thrift"

struct CommentActionRequest {
    1: required string token
    2: required i64 video_id
    3: required i32 action_type
    4: optional string comment_text
    5: optional i64 comment_id
}

struct CommentActionResponse {
    1: required i32 status_code
    2: optional string status_msg
    3: optional Comment comment
}

struct Comment {
    1: required i64 id
    2: required user.User user
    3: required string content
    4: required string create_date
}

struct CommentListRequest {
    1: required string token
    2: required i64 video_id
}

struct CommentListResponse {
    1: required i32 status_code
    2: required string status_msg
    3: list<Comment> comment_list
}

service CommentService {
    CommentActionResponse CommentAction(1: CommentActionRequest req)
    CommentListResponse CommentList(1: CommentListRequest req)
}