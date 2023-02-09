namespace go douyin_message

struct MessageChatRequest {
    1: required i64 user_id
    2: required i64 to_user_id
}

struct MessageChatResponse {
    1: required i32 status_code
    2: optional string status_msg
    3: list<Message> message_list
}

struct Message {
    1: required i64 id
    2: required i64 to_user_id
    3: required i64 from_user_id
    4: required string content
    5: optional string create_time
}

struct MessageActionRequest {
    1: required i64 user_id
    2: required i64 to_user_id
    3: required i32 action_type
    4: required string content
}

struct MessageActionResponse {
    1: required i32 status_code
    2: optional string status_msg
}

service MessageService {
    MessageChatResponse MessageChat(1: MessageChatRequest req)
    MessageActionResponse MessageAction(1: MessageActionRequest req)
}
