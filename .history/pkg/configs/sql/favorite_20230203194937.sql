CREATE TABLE `favorite` (
                            `user_id` bigint NOT NULL,
                           `video_id` bigint NOT NULL,
                           `isfavorite` tinyint NOT NULL,
                           `delstate` tinyint NOT NULL,
                           `utime` datetime NOT NULL,
                           PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;