CREATE TABLE `relation` (
                            `relation_id` bigint NOT NULL AUTO_INCREMENT,
                            `user_id` bigint NOT NULL,
                            `to_user_id` bigint NOT NULL,
                            `status` tinyint(3) unsigned zerofill NOT NULL,
                            `ctime` datetime NOT NULL,
                            `utime` datetime NOT NULL,
                            PRIMARY KEY (`relation_id`),
                            UNIQUE KEY `2ids` (`user_id`,`to_user_id`),
                            CONSTRAINT `relation_fk_to_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`),
                            CONSTRAINT `relation_fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;