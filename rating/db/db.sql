
CREATE TABLE `rate_item` (
`id` BIGINT(64) NOT NULL AUTO_INCREMENT,
`rater` varchar(64) NOT NULL,
`rated` varchar(64) NOT NULL,
`ride` varchar(64) NOT NULL,
`rate_value` tinyint NOT NULL,
`message` varchar(64),
`rater_type` tinyint NOT NULL,
`rated_type` tinyint NOT NULL,
`creation_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
`modification_time` DATETIME ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`id`),
UNIQUE (`rater`, `rated`, `ride`),
INDEX (`rater`),
INDEX (`rated`)
);

CREATE TABLE `rate_item_feedback` (
`rate_id` BIGINT(64) NOT NULL,
`feedback` tinyint NOT NULL,
`positive` tinyint NOT NULL,
`creation_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
`modification_time` DATETIME ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`rate_id`, `feedback`, `positive`),
FOREIGN KEY (`rate_id`) REFERENCES `rate_item`(`id`) ON DELETE CASCADE
);

CREATE TABLE `rate` (
`rated` varchar(64) NOT NULL,
`rate_total` bigint NOT NULL DEFAULT 0,
`rate_count` bigint NOT NULL DEFAULT 0,
`creation_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
`modification_time` DATETIME ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`rated`)
);
