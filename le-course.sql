-- DROP DATABASE IF EXISTS `course2`;
CREATE DATABASE IF NOT EXISTS `course2`;
USE `course2`;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` text NOT NULL,
  `email` text NOT NULL,
  `password` text NOT NULL,
  `no_hp` text NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

LOCK TABLES `users` WRITE;
INSERT INTO `users` VALUES
(1,'super admin','admin@gmail.com','ini password','08171122233333','2022-06-15 13:02:09','2022-06-15 13:02:09'),
(2,'ahsan','ahsan@mail.com','$2a$10$mM7/GAbcxBE1.Z2ALg83puE3Vcqn75UlAEexxa/xbIaQCmjb8PKoa','','2022-06-15 23:45:44','2022-06-15 23:45:44'),
(3,'ahsan','ahsan2@mail.com','$2a$10$UOJAFVPbm4QtObtejd0fy.RduB5brNCqGx4Kv10XpAVxmYIhOT9Vi','','2022-06-16 19:21:49','2022-06-16 19:21:49');
UNLOCK TABLES;

DROP TABLE IF EXISTS `exercises`;
CREATE TABLE `exercises` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` text NOT NULL,
  `description` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

LOCK TABLES `exercises` WRITE;
INSERT INTO `exercises` VALUES
(1,'Olimpiade Matematika SMA','Olimpiade Matematika tingkat SMA Jawa Timur 2099');
UNLOCK TABLES;

DROP TABLE IF EXISTS `questions`;
CREATE TABLE `questions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `exercise_id` int NOT NULL,
  `body` text NOT NULL,
  `option_a` text NOT NULL,
  `option_b` text NOT NULL,
  `option_c` text NOT NULL,
  `option_d` text NOT NULL,
  `correct_answer` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `score` int NOT NULL,
  `creator_id` int NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `questions_FK` (`exercise_id`),
  KEY `questions_FK_1` (`creator_id`),
  CONSTRAINT `questions_FK` FOREIGN KEY (`exercise_id`) REFERENCES `exercises` (`id`),
  CONSTRAINT `questions_FK_1` FOREIGN KEY (`creator_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

LOCK TABLES `questions` WRITE;
INSERT INTO `questions` VALUES
(1,1,'Berapa Jumlah hasil dari 1 + 1?','0.2','2','2.2','22','b',10,1,'2022-06-15 14:01:08','2022-06-15 14:01:08'),
(2,1,'Berapa Jumlah hasil dari 2 + 2?','4','4.4','44','0.4','a',10,1,'2022-06-15 14:01:08','2022-06-15 14:01:08'),
(3,1,'Berapa Jumlah hasil dari 1 x 1?','0.1','1','1.1','11','b',10,1,'2022-06-15 14:09:22','2022-06-15 14:09:22'),
(4,1,'Berapa Jumlah hasil dari 3 x 3?','999','9','9.9','99','b',10,1,'2022-06-15 14:09:50','2022-06-15 14:09:50'),
(5,1,'Berapa hasil dari 2 + 3?','0.5','0.55','-5','5','d',10,1,'2022-06-15 14:11:13','2022-06-15 14:11:13'),
(6,1,'Berapa hasil dari 23 x 0.1?','0.1','0.23','23','2.3','d',10,1,'2022-06-15 14:12:07','2022-06-15 14:12:07'),
(7,1,'Jika 3 - 2 = 1, berapakah hasil dari 3 + 1?','4','5','6','7','a',10,1,'2022-06-15 14:15:16','2022-06-15 14:15:16'),
(8,1,'Jika 2 + 2 = 4, berapakah hasil dari 3 + 3?','23','33','6','5','c',10,1,'2022-06-15 14:15:43','2022-06-15 14:15:43'),
(9,1,'Jika 10 + 1 = 11, berapakah hasil dari 30 x 1?','31','13','4','30','d',10,1,'2022-06-15 14:15:47','2022-06-15 14:15:47'),
(10,1,'Berapa hasil dari 9 + 3?','11','12','13','14','b',10,1,'2022-06-15 14:15:50','2022-06-15 14:15:50');
UNLOCK TABLES;

DROP TABLE IF EXISTS `answers`;
CREATE TABLE `answers` (
  `id` int NOT NULL AUTO_INCREMENT,
  `exercise_id` int NOT NULL,
  `question_id` int NOT NULL,
  `user_id` int NOT NULL,
  `answer` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `answers_user_id_IDX` (`user_id`,`question_id`) USING BTREE,
  KEY `answers_FK` (`exercise_id`),
  KEY `answers_FK_1` (`question_id`),
  CONSTRAINT `answers_FK` FOREIGN KEY (`exercise_id`) REFERENCES `exercises` (`id`),
  CONSTRAINT `answers_FK_1` FOREIGN KEY (`question_id`) REFERENCES `questions` (`id`),
  CONSTRAINT `answers_FK_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

LOCK TABLES `answers` WRITE;
INSERT INTO `answers` VALUES
(1,1,1,3,'b','2022-06-17 12:55:44','2022-06-17 12:55:44'),
(2,1,2,3,'c','2022-06-17 12:55:44','2022-06-17 12:55:44'),
(3,1,3,3,'a','2022-06-17 12:55:44','2022-06-17 12:55:44'),
(4,1,4,3,'c','2022-06-17 12:55:44','2022-06-17 12:55:44'),
(5,1,5,3,'d','2022-06-17 12:55:44','2022-06-17 12:55:44'),
(6,1,6,3,'b','2022-06-17 12:55:44','2022-06-17 12:55:44'),
(7,1,7,3,'d','2022-06-17 13:01:35','2022-06-17 13:01:35'),
(8,1,8,3,'c','2022-06-17 13:01:35','2022-06-17 13:01:35'),
(9,1,9,3,'b','2022-06-17 13:01:35','2022-06-17 13:01:35'),
(10,1,10,3,'b','2022-06-17 13:01:35','2022-06-17 13:01:35');
UNLOCK TABLES;
