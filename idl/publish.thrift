namespace go douyin_publish

include "feed.thrift"

struct PublishActionRequest {
    1: required i64 req_user_id
    2: required string title
    3: required string video_play_url
    4: required string video_cover_url
}

struct PublishActionResponse {
    1: required i32 status_code
    2: optional string status_msg
}

struct PublishListRequest {
    1: required i64 user_id
    2: required i64 req_user_id
}

struct PublishListResponse {
    1: required i32 status_code
    2: optional string status_msg
    3: list<feed.Video> video_list
}

service PublishService {
    PublishActionResponse PublishAction(1: PublishActionRequest req)
    PublishListResponse PublishList(1: PublishListRequest req)
}