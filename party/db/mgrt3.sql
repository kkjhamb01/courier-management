CREATE TABLE `admin` (
`id` varchar(64) NOT NULL,
`first_name` varchar(64) DEFAULT '',
`last_name` varchar(64) DEFAULT '',
`username` varchar(64) NOT NULL,
`password` varchar(64) NOT NULL,
`creation_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
`modification_time` DATETIME ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`id`)
);
