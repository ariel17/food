-- Adminer 4.8.0 MySQL 8.0.23 dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP DATABASE IF EXISTS `food`;
CREATE DATABASE `food` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `food`;

DROP TABLE IF EXISTS `ingredients`;
CREATE TABLE `ingredients` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET sjis COLLATE sjis_bin NOT NULL,
  `type` varchar(10) CHARACTER SET sjis COLLATE sjis_bin NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=sjis COLLATE=sjis_bin;


DROP TABLE IF EXISTS `plates`;
CREATE TABLE `plates` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET sjis COLLATE sjis_bin NOT NULL,
  `type` varchar(15) CHARACTER SET sjis COLLATE sjis_bin NOT NULL,
  `only_on` varchar(15) CHARACTER SET sjis COLLATE sjis_bin DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=sjis COLLATE=sjis_bin;


DROP TABLE IF EXISTS `plates_ingredients`;
CREATE TABLE `plates_ingredients` (
  `plate_id` int NOT NULL,
  `ingredient_id` int NOT NULL,
  `amount` int NOT NULL,
  `unit` varchar(10) CHARACTER SET sjis COLLATE sjis_bin NOT NULL,
  UNIQUE KEY `plates_ingredients_plate_id_ingredient_id` (`plate_id`,`ingredient_id`),
  KEY `plate_id` (`plate_id`),
  KEY `ingredient_id` (`ingredient_id`),
  CONSTRAINT `plates_ingredients_ingredient_id` FOREIGN KEY (`ingredient_id`) REFERENCES `ingredients` (`id`),
  CONSTRAINT `plates_ingredients_plate_id` FOREIGN KEY (`plate_id`) REFERENCES `plates` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=sjis COLLATE=sjis_bin;


-- 2021-03-29 14:20:41
