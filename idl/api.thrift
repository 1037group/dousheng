namespace go douyin_api

struct FeedRequest {
    1: optional string latest_time
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

struct UserRegisterRequest {
    1: required string username
    2: required string password
}

struct UserRegisterResponse {
    1: required i32 status_code
    2: optional string status_msg
    3: required i64 user_id
    4: required string token
}

struct UserLoginRequest {
    1: required string username
    2: required string password
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

service ApiService {
    FeedResponse Feed(1: FeedRequest req) (api.get="/douyin/feed")
    UserLoginResponse UserLogin(1: UserLoginRequest req) (api.post="/douin/user/login")
    UserRegisterResponse UserRegister(1: UserRegisterRequest req) (api.post="/douin/user/register")
    UserResponse User(1: UserRequest req) (api.get="/douin/user")
}