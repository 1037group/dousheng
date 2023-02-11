namespace go douyin_user

struct UserRegisterRequest {
    1: required string username
    2: required string password
}

struct UserRegisterResponse {
    1: required i32 status_code
    2: optional string status_msg
    3: required i64 user_id
}

struct UserLoginRequest {
    1: required string username
    2: required string password
}

struct UserLoginResponse {
    1: required i32 status_code
    2: optional string status_msg
    3: required i64 user_id
}

struct UserRequest {
    1: required i64 req_user_id
    2: required i64 user_id
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

struct FriendUser {
    1: required i64 id
    2: required string name
    3: optional i64 follow_count
    4: optional i64 follower_count
    5: required bool is_follow
    6: required string avatar
    7: optional string message
    8: required i64 msgType
}

service UserService {
    UserLoginResponse UserLogin(1: UserLoginRequest req)
    UserRegisterResponse UserRegister(1: UserRegisterRequest req)
    UserResponse User(1: UserRequest req)
}