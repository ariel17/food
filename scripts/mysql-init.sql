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
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=sjis COLLATE=sjis_bin;

INSERT INTO `ingredients` (`id`, `name`, `type`) VALUES
(1,	'Pescado',	'Carne'),
(2,	'Pan rallado',	'Harina'),
(3,	'Sal',	'Condimento'),
(4,	'Provenzal',	'Condimento'),
(5,	'Aceite',	'Grasa'),
(6,	'Carne de nalga',	'Carne'),
(7,	'Papa',	'Verdura'),
(8,	'Carne picada especial',	'Carne'),
(9,	'Cebolla',	'Verdura'),
(10,	'Morron',	'Verdura'),
(11,	'Tapa pascualina',	'Harina'),
(12,	'Leche',	'Grasa'),
(13,	'Nuez moscada',	'Condimento'),
(14,	'Manteca',	'Grasa'),
(15,	'Pechuga de pollo',	'Carne'),
(16,	'Costilla de chancho',	'Carne'),
(17,	'Chorizos',	'Carne'),
(18,	'Morcilla',	'Carne'),
(19,	'Pata-muslo de pollo',	'Carne'),
(20,	'Salchicha parrillera',	'Carne'),
(21,	'Prepizza',	'Harina'),
(22,	'Salsa de tomate',	'Verdura'),
(23,	'Condimento para pizza',	'Condimento'),
(24,	'Huevo',	'Carne'),
(25,	'Jamon cocido',	'Carne'),
(26,	'Queso fresco',	'Grasa'),
(27,	'Pimenton',	'Condimento'),
(28,	'Oregano',	'Condimento'),
(29,	'Tomate',	'Verdura'),
(30,	'Arroz',	'Verdura');

DROP TABLE IF EXISTS `plates`;
CREATE TABLE `plates` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET sjis COLLATE sjis_bin NOT NULL,
  `type` varchar(15) CHARACTER SET sjis COLLATE sjis_bin NOT NULL,
  `only_on` varchar(15) CHARACTER SET sjis COLLATE sjis_bin DEFAULT NULL,
  `needs_mixing` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=sjis COLLATE=sjis_bin;

INSERT INTO `plates` (`id`, `name`, `type`, `only_on`, `needs_mixing`) VALUES
(1,	'Milanesa de pescado',	'main',	NULL,	0),
(2,	'Milanesa de carne',	'main',	NULL,	0),
(3,	'Pastel de papa y carne',	'main',	NULL,	0),
(4,	'Milanesa de pollo',	'main',	NULL,	0),
(5,	'Asado',	'main',	NULL,	0),
(6,	'Pizza',	'main',	NULL,	0),
(7,	'Omellete',	'main',	NULL,	0),
(8,	'Guiso de arroz con carne',	'main',	NULL,	0),
(9,	'Tortilla de papa',	'main',	NULL,	0);

DROP TABLE IF EXISTS `plates_ingredients`;
CREATE TABLE `plates_ingredients` (
  `plate_id` int NOT NULL,
  `ingredient_id` int NOT NULL,
  `amount` float NOT NULL,
  `unit` varchar(10) CHARACTER SET sjis COLLATE sjis_bin DEFAULT NULL,
  UNIQUE KEY `plates_ingredients_plate_id_ingredient_id` (`plate_id`,`ingredient_id`),
  KEY `plate_id` (`plate_id`),
  KEY `ingredient_id` (`ingredient_id`),
  CONSTRAINT `plates_ingredients_ingredient_id` FOREIGN KEY (`ingredient_id`) REFERENCES `ingredients` (`id`),
  CONSTRAINT `plates_ingredients_plate_id` FOREIGN KEY (`plate_id`) REFERENCES `plates` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=sjis COLLATE=sjis_bin;

INSERT INTO `plates_ingredients` (`plate_id`, `ingredient_id`, `amount`, `unit`) VALUES
(1,	1,	1000,	'g'),
(1,	2,	500,	'g'),
(1,	3,	10,	'g'),
(1,	4,	10,	'g'),
(1,	5,	300,	'ml'),
(2,	2,	500,	'g'),
(2,	3,	10,	'g'),
(2,	4,	10,	'g'),
(2,	5,	300,	'ml'),
(2,	6,	1000,	'g'),
(3,	7,	500,	'g'),
(3,	8,	1000,	'g'),
(3,	9,	500,	'g'),
(3,	10,	0.25,	NULL),
(3,	11,	1,	NULL),
(3,	12,	100,	'ml'),
(3,	13,	10,	'g'),
(3,	14,	20,	'g'),
(4,	2,	500,	'g'),
(4,	3,	10,	'g'),
(4,	4,	10,	'g'),
(4,	15,	1000,	'g'),
(5,	16,	1000,	'g'),
(5,	17,	6,	NULL),
(5,	18,	2,	NULL),
(5,	19,	2,	NULL),
(5,	20,	1,	NULL),
(6,	5,	100,	'ml'),
(6,	21,	2,	NULL),
(6,	22,	1,	NULL),
(6,	23,	30,	'g'),
(6,	24,	4,	NULL),
(6,	25,	250,	'g'),
(6,	26,	400,	'g'),
(7,	3,	10,	'g'),
(7,	4,	80,	'g'),
(7,	24,	8,	NULL),
(7,	25,	150,	'g'),
(7,	26,	300,	'g'),
(7,	27,	50,	'g'),
(7,	28,	50,	'g'),
(7,	29,	2,	NULL),
(8,	6,	500,	'g'),
(8,	9,	500,	'g'),
(8,	10,	0.5,	NULL),
(8,	30,	500,	'g'),
(9,	3,	100,	'g'),
(9,	5,	200,	'ml'),
(9,	7,	1000,	'g'),
(9,	24,	4,	NULL);

-- 2021-04-02 05:04:28
