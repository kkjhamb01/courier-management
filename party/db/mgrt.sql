DROP INDEX email ON courier;
DROP INDEX email ON client;

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
`mot_status` varchar(64),
`revenue_weight` int,
`colour` varchar(64),
`make` varchar(64),
`type_approval` varchar(64),
`year_of_manufacture` int,
`tax_due_date` varchar(64),
`tax_status` varchar(64),
`date_of_last_v5c_issued` varchar(64),
`real_driving_emissions` varchar(64),
`wheelplan` varchar(64),
`month_of_first_registration` varchar(64),
`creation_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
`modification_time` DATETIME ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`user_id`,`registration_number`),
FOREIGN KEY (`user_id`) REFERENCES `courier`(`id`) ON DELETE CASCADE
);


