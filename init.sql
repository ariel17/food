-- Adminer 4.8.0 MySQL 8.0.23 dump

DROP DATABASE IF EXISTS `food`;
CREATE DATABASE `food`;
USE `food`;

SET NAMES utf8;
SET time_zone = '-03:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

DROP TABLE IF EXISTS `ingredients`;
CREATE TABLE `ingredients` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE sjis_bin NOT NULL,
  `type` varchar(10) COLLATE sjis_bin NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=sjis COLLATE=sjis_bin;


DROP TABLE IF EXISTS `plates`;
CREATE TABLE `plates` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE sjis_bin NOT NULL,
  `only_on` varchar(15) COLLATE sjis_bin DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=sjis COLLATE=sjis_bin;


DROP TABLE IF EXISTS `plates_ingredients`;
CREATE TABLE `plates_ingredients` (
  `plate_id` int NOT NULL,
  `ingredient_id` int NOT NULL,
  `amount` int NOT NULL,
  `unit` varchar(10) COLLATE sjis_bin NOT NULL,
  KEY `plate_id` (`plate_id`),
  KEY `ingredient_id` (`ingredient_id`),
  CONSTRAINT `plates_ingredients_plate_id` FOREIGN KEY (`plate_id`) REFERENCES `plates` (`id`),
  CONSTRAINT `plates_ingredients_ingredient_id` FOREIGN KEY (`ingredient_id`) REFERENCES `ingredients` (`id`),
  UNIQUE KEY `plates_ingredients_plate_id_ingredient_id` (`plate_id`, `ingredient_id`)
) ENGINE=InnoDB DEFAULT CHARSET=sjis COLLATE=sjis_bin;


-- 2021-03-29 00:31:39
