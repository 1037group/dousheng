namespace go douyin_relation

include "user.thrift"

struct RelationActionRequest {
    1: required string token
    2: required i64 to_user_id
    3: required i32 action_type
}

struct RelationActionResponse {
    1: required i32 status_code
    2: optional string status_msg
}

struct RelationFollowListRequest {
    1: required i64 user_id
    2: required string token
}

struct RelationFollowListResponse {
    1: required i32 status_code
    2: optional string status_msg
    3: list<user.User> user_list
}

struct RelationFriendListRequest {
    1: required i64 user_id
    2: required string token
}

struct RelationFriendListResponse {
    1: required i32 status_code
    2: optional string status_msg
    3: list<user.User> user_list
}

service RelationService {
    RelationActionResponse RelationAction(1: RelationActionRequest req)
    RelationFollowListResponse RelationFollowList(1: RelationFollowListRequest req)
    RelationFriendListResponse RelationFriendList(1: RelationFriendListRequest req)

}
