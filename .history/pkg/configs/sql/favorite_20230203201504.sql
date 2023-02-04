CREATE TABLE `favorite` (
                            `user_id` bigint NOT NULL,
                           `video_id` bigint NOT NULL,
                           `isfavorite` tinyint NOT NULL,
                           `delstate` tinyint NOT NULL,
                           `utime` datetime NOT NULL,
                           PRIMARY KEY (`user_id`),
                           UNIQUE KEY `uk_user_video_id` (`user_id`,`video_id`),
                           CONSTRAINT `idx_update_time` FOREIGN KEY (`utime`) REFERENCES `user` (`utime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;