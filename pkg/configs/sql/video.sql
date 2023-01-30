CREATE TABLE `video` (
                         `video_id` bigint NOT NULL AUTO_INCREMENT,
                         `user_id` bigint NOT NULL,
                         `video_play_url` varchar(255) NOT NULL,
                         `video_cover_url` varchar(255) NOT NULL,
                         `video_favorite_count` bigint(20) unsigned zerofill NOT NULL,
                         `video_comment_count` bigint(20) unsigned zerofill NOT NULL,
                         `video_title` varchar(255) NOT NULL,
                         `ctime` datetime NOT NULL,
                         `utime` datetime NOT NULL,
                         PRIMARY KEY (`video_id`),
                         KEY `video_fk_user_id` (`user_id`),
                         CONSTRAINT `video_fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;