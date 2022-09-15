
CREATE TABLE `announcement` (
`id` BIGINT(64) NOT NULL AUTO_INCREMENT,
`title` varchar(255) NOT NULL,
`text` text NOT NULL,
`type` tinyint NOT NULL,
`message_type` tinyint NOT NULL,
`creation_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
`modification_time` DATETIME ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`id`)
);

CREATE TABLE `announcement_user` (
`announcement_id` BIGINT(64) NOT NULL,
`user_id` varchar(64) NOT NULL,
`creation_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
`modification_time` DATETIME ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`announcement_id`, `user_id`),
FOREIGN KEY (`announcement_id`) REFERENCES `announcement`(`id`) ON DELETE CASCADE
);
