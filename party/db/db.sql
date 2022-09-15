
CREATE TABLE `courier` (
`id` varchar(64) NOT NULL,
`first_name` varchar(64) DEFAULT '',
`last_name` varchar(64) NOT NULL,
`email` varchar(64),
`phone_number` varchar(32),
`status` int NOT NULL DEFAULT 0,
`birth_date` varchar(32),
`transport_type` tinyint,
`transport_size` tinyint,
`citizen` tinyint,
`creation_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
`modification_time` DATETIME ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`id`)
);

CREATE TABLE `courier_address` (
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
FOREIGN KEY (`user_id`) REFERENCES `courier`(`id`) ON DELETE CASCADE
);

CREATE TABLE `courier_claim` (
`user_id` varchar(64) NOT NULL,
`claim_type` tinyint NOT NULL,
`identifier` varchar(64) NOT NULL,
`creation_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
`modification_time` DATETIME ON UPDATE CURRENT_TIMESTAMP,
UNIQUE (`identifier`),
PRIMARY KEY (`identifier`),
FOREIGN KEY (`user_id`) REFERENCES `courier`(`id`) ON DELETE CASCADE
);

CREATE TABLE `courier_id_card` (
`user_id` varchar(64) NOT NULL,
`first_name` varchar(64) NOT NULL,
`last_name` varchar(64) NOT NULL,
`number` varchar(64) NOT NULL,
`expiration_date` varchar(64) NOT NULL,
`issue_place` varchar(64) NOT NULL,
`type` tinyint NOT NULL,
`creation_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
`modification_time` DATETIME ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`user_id`),
FOREIGN KEY (`user_id`) REFERENCES `courier`(`id`) ON DELETE CASCADE
);

CREATE TABLE `courier_driving_license` (
`user_id` varchar(64) NOT NULL,
`driving_license_number` varchar(64) NOT NULL,
`expiration_date` varchar(64) NOT NULL,
`creation_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
`modification_time` DATETIME ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`user_id`),
FOREIGN KEY (`user_id`) REFERENCES `courier`(`id`) ON DELETE CASCADE
);

CREATE TABLE `courier_driver_background` (
`user_id` varchar(64) NOT NULL,
`national_insurance_number` varchar(64) NOT NULL,
`upload_dbs_later` tinyint,
`creation_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
`modification_time` DATETIME ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`user_id`),
FOREIGN KEY (`user_id`) REFERENCES `courier`(`id`) ON DELETE CASCADE
);

CREATE TABLE `courier_residence_card` (
`user_id` varchar(64) NOT NULL,
`number` varchar(64) NOT NULL,
`expiration_date` varchar(64) NOT NULL,
`issue_date` varchar(64) NOT NULL,
`creation_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
`modification_time` DATETIME ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`user_id`),
FOREIGN KEY (`user_id`) REFERENCES `courier`(`id`) ON DELETE CASCADE
);

CREATE TABLE `courier_bank_account` (
`user_id` varchar(64) NOT NULL,
`bank_name` varchar(64) NOT NULL,
`account_number` varchar(64) NOT NULL UNIQUE,
`account_holder_name` varchar(64) NOT NULL UNIQUE,
`sort_code` varchar(64) NOT NULL,
`creation_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
`modification_time` DATETIME ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`user_id`),
FOREIGN KEY (`user_id`) REFERENCES `courier`(`id`) ON DELETE CASCADE
);

CREATE TABLE `document` (
`user_id` varchar(64) NOT NULL,
`object_id` varchar(64) NOT NULL UNIQUE,
`document_info_type` tinyint NOT NULL,
`document_type` tinyint NOT NULL,
`file_type` varchar(16),
`creation_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
`modification_time` DATETIME ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`user_id`, `object_id`),
FOREIGN KEY (`user_id`) REFERENCES `courier`(`id`) ON DELETE CASCADE
);

CREATE TABLE `document_storage` (
`object_id` varchar(64) NOT NULL,
`data` MEDIUMBLOB NOT NULL,
`creation_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
`modification_time` DATETIME ON UPDATE CURRENT_TIMESTAMP,
FOREIGN KEY (`object_id`) REFERENCES `document`(`object_id`) ON DELETE CASCADE
);

CREATE TABLE `courier_status` (
`user_id` varchar(64) NOT NULL,
`status_type` tinyint NOT NULL,
`status` tinyint NOT NULL,
`message` varchar(128),
`creation_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
`modification_time` DATETIME ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`user_id`,`status_type`),
FOREIGN KEY (`user_id`) REFERENCES `courier`(`id`) ON DELETE CASCADE
);

CREATE TABLE `courier_mot` (
`user_id` varchar(64) NOT NULL,
`registration_number` varchar(64) NOT NULL,
`co2_emissions` int,
`engine_capacity` int,
`euro_status` varchar(64),
`marked_for_export` int,
`fuel_type` varchar(64),
`mot_status` tinyint,
`revenue_weight` int,
`colour` varchar(64),
`make` varchar(64),
`type_approval` varchar(64),
`year_of_manufacture` int,
`tax_due_date` varchar(64),
`tax_status` tinyint,
`date_of_last_v5c_issued` varchar(64),
`real_driving_emissions` varchar(64),
`wheelplan` varchar(64),
`month_of_first_registration` varchar(64),
`creation_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
`modification_time` DATETIME ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`user_id`,`registration_number`),
FOREIGN KEY (`user_id`) REFERENCES `courier`(`id`) ON DELETE CASCADE
);

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

CREATE TABLE `client` (
`id` varchar(64) NOT NULL,
`first_name` varchar(64) DEFAULT '',
`last_name` varchar(64) NOT NULL,
`email` varchar(64),
`phone_number` varchar(32),
`status` int NOT NULL DEFAULT 0,
`payment_method` tinyint,
`birth_date` varchar(32),
`referral` varchar(16) NOT NULL,
`creation_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
`modification_time` DATETIME ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`id`)
);

CREATE TABLE `client_claim` (
`user_id` varchar(64) NOT NULL,
`claim_type` tinyint NOT NULL,
`identifier` varchar(64) NOT NULL,
`creation_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
`modification_time` DATETIME ON UPDATE CURRENT_TIMESTAMP,
UNIQUE (`identifier`),
PRIMARY KEY (`identifier`),
FOREIGN KEY (`user_id`) REFERENCES `client`(`id`) ON DELETE CASCADE
);

CREATE TABLE `client_card` (
`user_id` varchar(64) NOT NULL,
`card_number` varchar(64) NOT NULL,
`issue_date` varchar(64) NOT NULL,
`cvv` varchar(64) NOT NULL,
`zip_code` varchar(64) NOT NULL,
`country` varchar(64) NOT NULL,
`creation_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
`modification_time` DATETIME ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`user_id`, `card_number`),
FOREIGN KEY (`user_id`) REFERENCES `client`(`id`) ON DELETE CASCADE
);

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
