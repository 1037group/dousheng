namespace go douyin_api

/*
*
* 调用逻辑：
* 客户端 --(api idl)-> http服务(api) --(services idl)->
* rpc服务(feed user publish favorite comment relation message) -> db
*
* */


/*
*
* feed
*
* */

struct FeedRequest {
    1: optional i64 latest_time
    2: optional string token
}

struct Video {
    1: required i64 id
    2: required User author
    3: required string play_url
    4: required string cover_url
    5: required i64 favorite_count
    6: required i64 comment_count
    7: required bool is_favorite
    8: required string title
}

struct FeedResponse {
	1: required i32 status_code
	2: optional string status_msg
	3: list<Video> video_list
	4: optional i64 next_time
}

/*
*
* user
*
* */
struct UserRegisterRequest {
    1: required string username (api.query="username", api.vd="len($) > 0")
    2: required string password (api.query="password", api.vd="len($) > 0")
}

struct UserRegisterResponse {
    1: required i32 status_code
    2: optional string status_msg
    3: required i64 user_id
    4: required string token
}

struct UserLoginRequest {
    1: required string username (api.query="username", api.vd="len($) > 0")
    2: required string password (api.query="password", api.vd="len($) > 0")
}

struct UserLoginResponse {
    1: required i32 status_code
    2: optional string status_msg
    3: required i64 user_id
    4: required string token
}

struct UserRequest {
    1: required i64 user_id
    2: required string token
}

struct UserResponse {
    1: required i32 status_code
    2: optional string status_msg
    3: required User user
}

struct User {
    1: required i64 id
    2: required string name
    3: optional i64 follow_count
    4: optional i64 follower_count
    5: required bool is_follow
}

/*
*
* publish
*
* */

struct PublishActionRequest {
    1: required string token
    2: required byte data
    3: required string title
}

struct PublishActionResponse {
    1: required i32 status_code
    2: optional string status_msg
}

struct PublishListRequest {
    1: required i64 user_id
    2: required string token
}

struct PublishListResponse {
    1: required i32 status_code
    2: optional string status_msg
    3: list<Video> video_list
}

/*
*
* favorite
*
* */

struct FavoriteActionRequest {
    1: required string token
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
    3: list<Video> video_list
}

/*
*
* comment
*
* */

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
    2: required User user
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

/*
*
* relation
*
* */

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
    3: list<User> user_list
}

struct RelationFriendListRequest {
    1: required i64 user_id
    2: required string token
}

struct RelationFriendListResponse {
    1: required i32 status_code
    2: optional string status_msg
    3: list<User> user_list
}

/*
*
* message
*
* */

struct MessageChatRequest {
    1: required string token
    2: required i64 to_user_id
}

struct MessageChatResponse {
    1: required i32 status_code
    2: optional string status_msg
    3: list<Message> message_list
}

struct Message {
    1: required i64 id
    2: required string content
    3: optional string create_time
}

struct MessageActionRequest {
    1: required string token
    2: required i64 to_user_id
    3: required i32 action_type
    4: required string content
}

struct MessageActionResponse {
    1: required i32 status_code
    2: optional string status_msg
}

service ApiService {
    FeedResponse Feed(1: FeedRequest req) (api.get="/douyin/feed/")

    UserResponse User(1: UserRequest req) (api.get="/douyin/user/")
    UserLoginResponse UserLogin(1: UserLoginRequest req) (api.post="/douyin/user/login/")
    UserRegisterResponse UserRegister(1: UserRegisterRequest req) (api.post="/douyin/user/register/")

    PublishActionResponse PublishAction(1: PublishActionRequest req) (api.post="/douyin/publish/action/")
    PublishListResponse PublishList(1: PublishListRequest req) (api.get="/douyin/publish/list/")

    FavoriteActionResponse FavoriteAction(1: FavoriteActionRequest req) (api.post="/douyin/favorite/action/")
    FavoriteListResponse FavoriteList(1: FavoriteListRequest req) (api.get="douyin/favorite/list/")

    CommentActionResponse CommentAction(1: CommentActionRequest req) (api.post="douyin/comment/action/")
    CommentListResponse CommentList(1: CommentListRequest req) (api.get="douyin/comment/list/")

    RelationActionResponse RelationAction(1: RelationActionRequest req) (api.post="douyin/relation/action/")
    RelationFollowListResponse RelationFollowList(1: RelationFollowListRequest req) (api.get="douyin/relation/follow/list/")
    RelationFriendListResponse RelationFriendList(1: RelationFriendListRequest req) (api.get="douyin/relation/friend/list/")

    MessageChatResponse MessageChat(1: MessageChatRequest req) (api.get="douyin/message/chat/")
    MessageActionResponse MessageAction(1: MessageActionRequest req) (api.post="douyin/message/action/")
}