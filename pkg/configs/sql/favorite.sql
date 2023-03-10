CREATE TABLE `favorite` (
                            `favorite_id` bigint NOT NULL AUTO_INCREMENT,
                            `user_id` bigint NOT NULL,
                           `video_id` bigint NOT NULL,
                           `isfavorite` integer NOT NULL,
                           `delstate` integer NOT NULL,
                           `utime` datetime NOT NULL,
                           `ctime` datetime NOT NULL,
                           PRIMARY KEY (`favorite_id`),
                           UNIQUE KEY `uk_user_video_id` (`user_id`,`video_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;