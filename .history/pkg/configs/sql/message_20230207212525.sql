CREATE TABLE `message` (
                           `message_id` bigint NOT NULL AUTO_INCREMENT,
                           `user_id` bigint NOT NULL,
                           `to_user_id` bigint NOT NULL,
                           `message_content` varchar(255) NOT NULL,
                           `ctime` datetime NOT NULL,
                           `utime` datetime NOT NULL,
                           PRIMARY KEY (`comment_id`),
                           KEY `comment_fk_video_id` (`video_id`),
                           KEY `comment_fk_user_id` (`user_id`),
                           CONSTRAINT `comment_fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`),
                           CONSTRAINT `comment_fk_video_id` FOREIGN KEY (`video_id`) REFERENCES `video` (`video_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;