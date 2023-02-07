CREATE TABLE `message` (
                           `message_id` bigint NOT NULL AUTO_INCREMENT,
                           `user_id` bigint NOT NULL,
                           `to_user_id` bigint NOT NULL,
                           `message_content` varchar(255) NOT NULL,
                           `ctime` datetime NOT NULL,
                           `utime` datetime NOT NULL,
                           PRIMARY KEY (`comment_id`),
                           UNIQUE KEY `2ids` (`user_id`,`to_user_id`),
                           CONSTRAINT `comment_fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`),
                           CONSTRAINT `comment_fk_to_user_id` FOREIGN KEY (`to_user_id`) REFERENCES `user` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;