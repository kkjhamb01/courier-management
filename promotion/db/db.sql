
CREATE TABLE `promotion` (
`id` BIGINT(64) NOT NULL AUTO_INCREMENT,
`name` varchar(64) NOT NULL,
`start_date` DATE DEFAULT NULL,
`exp_date` DATE DEFAULT NULL,
`discount_percentage` float DEFAULT 0,
`discount_value` float DEFAULT 0,
`type` tinyint NOT NULL,
`creation_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
`modification_time` DATETIME ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`id`)
);

CREATE TABLE `promotion_user` (
`promotion_id` BIGINT(64) NOT NULL,
`user_id` varchar(64) NOT NULL,
`metadata` varchar(256),
`status` tinyint DEFAULT 0,
`creation_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
`modification_time` DATETIME ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`promotion_id`, `user_id`),
FOREIGN KEY (`promotion_id`) REFERENCES `promotion`(`id`) ON DELETE CASCADE
);

CREATE TABLE `promotion_history` (
`promotion_id` BIGINT(64) NOT NULL,
`user_id` varchar(64) NOT NULL,
`transaction_id` varchar(256) NOT NULL,
`date` DATETIME NOT NULL,
`creation_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
`modification_time` DATETIME ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`promotion_id`, `user_id`),
FOREIGN KEY (`promotion_id`) REFERENCES `promotion`(`id`) ON DELETE CASCADE
);
