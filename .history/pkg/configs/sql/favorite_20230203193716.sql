CREATE TABLE `favorite` (
                           `video_id` bigint NOT NULL,
                           `user_id` bigint NOT NULL,
                           `isfavorite` tinyint NOT NULL,
                           `delstate` tinyint NOT NULL,
                           `utime` datetime NOT NULL,
                           PRIMARY KEY (`comment_id`),
                           KEY `comment_fk_video_id` (`video_id`),
                           KEY `comment_fk_user_id` (`user_id`),
                           CONSTRAINT `comment_fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`),
                           CONSTRAINT `comment_fk_video_id` FOREIGN KEY (`video_id`) REFERENCES `video` (`video_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;