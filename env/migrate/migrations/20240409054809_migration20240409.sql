-- Create "pictures" table
CREATE TABLE `pictures` (`id` char(36) NOT NULL, `content` mediumblob NOT NULL, `extension` varchar(255) NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
