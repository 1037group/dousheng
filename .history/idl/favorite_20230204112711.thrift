namespace go douyin_favorite

include "feed.thrift"

struct FavoriteActionRequest {
    1: required i64 user_id
    2: required i64 video_id
    3: required i32 action_type
}

struct FavoriteActionResponse {
    1: required i32 status_code
    2: optional string status_msg
}

struct FavoriteListRequest {
    1: required i64 user_id
    2: required string token
}

struct FavoriteListResponse {
    1: required i32 status_code
    2: optional string status_msg
    3: list<feed.Video> video_list
}

service FavoriteService {
    FavoriteActionResponse FavoriteAction(1: FavoriteActionRequest req)
    FavoriteListResponse FavoriteList(1: FavoriteListRequest req)

}
