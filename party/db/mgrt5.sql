ALTER TABLE `client` ADD COLUMN `birth_date` varchar(32);

CREATE TABLE `client_address` (
`user_id` varchar(64) NOT NULL,
`street` varchar(64),
`building` varchar(64),
`city` varchar(64),
`county` varchar(64),
`post_code` varchar(64),
`address_details` varchar(64),
`creation_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
`modification_time` DATETIME ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`user_id`),
FOREIGN KEY (`user_id`) REFERENCES `client`(`id`) ON DELETE CASCADE
);
