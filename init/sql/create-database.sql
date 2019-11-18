DROP DATABASE IF EXISTS `master`;
CREATE DATABASE `master`;

USE `master`;
CREATE TABLE `works` (
    `id` int(10) NOT NULL AUTO_INCREMENT COMMENT "id",
    `name` varchar(255) NOT NULL COMMENT "name",
    PRIMARY KEY (`id`)
);