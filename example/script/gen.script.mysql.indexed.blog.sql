
USE `ezsearch`;
CREATE TABLE `indexed_blog` (
	`id` INT(11) NOT NULL DEFAULT '0',
	`user_id` INT(11) NOT NULL DEFAULT '0',
	`hash` VARCHAR(100) NOT NULL DEFAULT '',
	`title` VARCHAR(100) NOT NULL DEFAULT '',
	`content` VARCHAR(100) NOT NULL DEFAULT '',
	`readed` INT(11) NOT NULL DEFAULT '0',
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT 'indexed_blog';

