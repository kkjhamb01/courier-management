ALTER TABLE `courier` ADD COLUMN `transport_size` tinyint;
DELETE FROM `courier_mot`;
ALTER TABLE `courier_mot` MODIFY COLUMN `mot_status` tinyint;
ALTER TABLE `courier_mot` MODIFY COLUMN `tax_status` tinyint;
