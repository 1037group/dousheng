CREATE TABLE `message` (
                           `message_id` bigint NOT NULL AUTO_INCREMENT,
                           `store_by_user_id` bigint NOT NULL,
                           `user_id` bigint NOT NULL,
                           `to_user_id` bigint NOT NULL,
                           `message_content` varchar(255) NOT NULL,
                           `is_read` integer NOT NULL,
                           `ctime` datetime NOT NULL,
                           `utime` datetime NOT NULL,
                           PRIMARY KEY (`message_id`),
                           CONSTRAINT `comment_mss_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`),
                           CONSTRAINT `comment_mss_to_user_id` FOREIGN KEY (`to_user_id`) REFERENCES `user` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;