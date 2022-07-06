CREATE DATABASE IF NOT EXISTS `elcorteingles`;
USE `elcorteingles`;
CREATE TABLE IF NOT EXISTS `products` (
	`id` varchar(100) NOT NULL,
	`title` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
	`originalPrice` float DEFAULT NULL,
	`finalPrice` float NOT NULL,
	`discount` float DEFAULT NULL,
	`category` varchar(100) DEFAULT NULL,
	`link` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;