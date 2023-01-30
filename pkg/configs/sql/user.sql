CREATE TABLE `user` (
                        `user_id` bigint NOT NULL AUTO_INCREMENT,
                        `user_name` varchar(255) NOT NULL,
                        `user_follow_count` bigint(20) unsigned zerofill NOT NULL,
                        `user_follower_count` bigint(20) unsigned zerofill NOT NULL,
                        `ctime` datetime NOT NULL,
                        `utime` datetime NOT NULL,
                        `password_hash` varchar(255) NOT NULL,
                        PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;