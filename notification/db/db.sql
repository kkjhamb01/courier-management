
CREATE TABLE `device` (
`device_id` varchar(128) NOT NULL,
`phone_number` varchar(64) NOT NULL,
`manufacturer` varchar(64) NOT NULL,
`device_model` varchar(64) NOT NULL,
`device_os` tinyint NOT NULL,
`device_version` varchar(32) NOT NULL,
`device_token` varchar(255) NOT NULL,
`creation_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
`modification_time` DATETIME ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`device_id`),
INDEX (`phone_number`)
);