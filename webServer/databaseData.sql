# ************************************************************
# Sequel Pro SQL dump
# Version 5446
#
# https://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.26)
# Database: NoTrackWebsiteStats
# Generation Time: 2020-06-01 13:03:28 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table failed_jobs
# ------------------------------------------------------------

DROP TABLE IF EXISTS `failed_jobs`;

CREATE TABLE `failed_jobs` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `connection` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `queue` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `payload` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `exception` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `failed_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;



# Dump of table migrations
# ------------------------------------------------------------

DROP TABLE IF EXISTS `migrations`;

CREATE TABLE `migrations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `migration` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `batch` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `migrations` WRITE;
/*!40000 ALTER TABLE `migrations` DISABLE KEYS */;

INSERT INTO `migrations` (`id`, `migration`, `batch`)
VALUES
	(1,'2014_10_12_000000_create_users_table',1),
	(2,'2014_10_12_100000_create_password_resets_table',1),
	(3,'2019_08_19_000000_create_failed_jobs_table',1),
	(4,'2020_05_18_140530_create_initial_tables',1),
	(5,'2020_05_19_084405_create_users_api_token',2),
	(6,'2020_05_23_151840_create_websockets_statistics_entries_table',3),
	(7,'2020_05_29_120550_add_ipaddress_collumn',4);

/*!40000 ALTER TABLE `migrations` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table password_resets
# ------------------------------------------------------------

DROP TABLE IF EXISTS `password_resets`;

CREATE TABLE `password_resets` (
  `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `token` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  KEY `password_resets_email_index` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;



# Dump of table users
# ------------------------------------------------------------

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email_verified_at` timestamp NULL DEFAULT NULL,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `api_token` varchar(80) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `remember_token` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `users_email_unique` (`email`),
  UNIQUE KEY `users_api_token_unique` (`api_token`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;

INSERT INTO `users` (`id`, `name`, `email`, `email_verified_at`, `password`, `api_token`, `remember_token`, `created_at`, `updated_at`)
VALUES
	(1,'Scott Andrews','scottandrews490@gmail.com',NULL,'$2y$10$fuwq2mqm9tHLMMfiXhRqq.Ag/tLPP.S9g71tlEGnqlBINDWedm13C','LqEYJ1rpIntd8A9ThQfwHqrypdhWCUDKc3jUQjr4YGrG21AxUhMMJRVhb8dh',NULL,'2020-05-18 15:09:28','2020-05-18 15:09:28'),
	(2,'Test User','scottandrews4900@gmail.com',NULL,'$2y$10$px3fzQ.LtzJP7UVuTXUXwuVUyKO9p7jU/eb8Olw04xh6WjDBpoIU6',NULL,NULL,'2020-05-18 17:35:41','2020-05-18 17:35:41'),
	(5,'test api 2','scottandrews49000@gmail.com',NULL,'$2y$10$7wPc7zdEGFgwTsX.d7miy.jkPq6zw46oowJ8BI5OPG4F2GtVKrdGK','LqEYJ1rpIntd8A9ThQfwHqrypdhWCUDKc3jUQjr4YGrG21AxUhMMJRVhb8du',NULL,'2020-05-19 08:59:21','2020-05-19 08:59:21');

/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table views
# ------------------------------------------------------------

DROP TABLE IF EXISTS `views`;

CREATE TABLE `views` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `referrer` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `currentPage` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `screenWidth` int(11) DEFAULT NULL,
  `language` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `browser` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `country` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `ipEncrypted` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT '',
  `website_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `views_website_id_foreign` (`website_id`),
  CONSTRAINT `views_website_id_foreign` FOREIGN KEY (`website_id`) REFERENCES `websites` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `views` WRITE;
/*!40000 ALTER TABLE `views` DISABLE KEYS */;

INSERT INTO `views` (`id`, `referrer`, `currentPage`, `screenWidth`, `language`, `browser`, `country`, `created_at`, `updated_at`, `ipEncrypted`, `website_id`)
VALUES
	(1,'http://127.0.0.1:8000/register','http://127.0.0.1:8000/home',3008,'en-GB','Safari','N/A',NULL,NULL,'',1),
	(2,'http://127.0.0.1:8000/register','http://127.0.0.1:8000/home',3008,'en-GB','Safari','N/A','2020-05-18 15:58:13','2020-05-18 15:58:13','',1),
	(3,'http://127.0.0.1:8000/register','http://127.0.0.1:8000/home',3008,'en-GB','Safari','N/A','2020-05-18 16:10:48','2020-05-18 16:10:48','',1),
	(4,'http://127.0.0.1:8000/register','http://127.0.0.1:8000/home',3008,'en-GB','Safari','N/A','2020-05-18 16:12:56','2020-05-18 16:12:56','',1),
	(5,'http://127.0.0.1:8000/register','http://127.0.0.1:8000/home',3008,'en-GB','Safari','N/A','2020-05-18 16:13:26','2020-05-18 16:13:26','',1),
	(6,'http://127.0.0.1:8000/register','http://127.0.0.1:8000/home',3008,'en-GB','Safari','N/A','2020-05-18 16:16:37','2020-05-18 16:16:37','',1),
	(7,'http://127.0.0.1:8000/register','http://127.0.0.1:8000/home',3008,'en-GB','Safari','N/A','2020-05-18 16:17:07','2020-05-18 16:17:07','',1),
	(8,'http://127.0.0.1:8000/register','http://127.0.0.1:8000/home',3008,'en-GB','Safari','N/A','2020-05-18 16:18:17','2020-05-18 16:18:17','fdfgfdgdfg',1),
	(9,'http://127.0.0.1:8000/register','http://127.0.0.1:8000/home',3008,'en-GB','Safari','N/A','2020-05-18 16:28:02','2020-05-18 16:28:02','fdfgfdgdfg',1),
	(10,'http://127.0.0.1:8000/register','http://127.0.0.1:8000/home',3008,'en-GB','Safari','N/A','2020-05-18 16:30:02','2020-05-18 16:30:02','dgfdgfdg',1),
	(11,'http://127.0.0.1:8000/register','http://127.0.0.1:8000/home',3008,'en-GB','Safari','N/A','2020-05-18 16:30:37','2020-05-18 16:30:37','dgfdgfdg',1),
	(12,'http://127.0.0.1:8000/register','http://127.0.0.1:8000/home',3008,'en-GB','Safari','N/A','2020-05-18 16:31:08','2020-05-18 16:31:08','',1),
	(13,'http://127.0.0.1:8000/register','http://127.0.0.1:8000/home',3008,'en-GB','Safari','N/A','2020-05-18 16:31:44','2020-05-18 16:31:44','',1),
	(14,'http://127.0.0.1:8000/register','http://127.0.0.1:8000/home',3008,'en-GB','Safari','N/A','2020-05-18 16:32:14','2020-05-18 16:32:14','',1),
	(15,'http://127.0.0.1:8000/register','http://127.0.0.1:8000/home',3008,'en-GB','Safari','N/A','2020-05-18 16:33:26','2020-05-18 16:33:26','',1),
	(16,'http://127.0.0.1:8000/register','http://127.0.0.1:8000/home',3008,'en-GB','Safari','N/A','2020-05-18 16:33:51','2020-05-18 16:33:51','',1),
	(17,'http://127.0.0.1:8000/register','http://127.0.0.1:8000/home',3008,'en-GB','Safari','N/A','2020-05-18 16:34:06','2020-05-18 16:34:06','',1),
	(18,'http://127.0.0.1:8000/register','http://127.0.0.1:8000/home',3008,'en-GB','Safari','N/A','2020-05-18 16:35:40','2020-05-18 16:35:40','',1),
	(19,'http://127.0.0.1:8000/register','http://127.0.0.1:8000/home',3008,'en-GB','Safari','N/A','2020-05-18 16:35:53','2020-05-18 16:35:53','',1),
	(20,'http://127.0.0.1:8000/register','http://127.0.0.1:8000/home',3008,'en-GB','Safari','N/A','2020-05-18 16:36:23','2020-05-18 16:36:23','',1),
	(21,'http://127.0.0.1:8000/register','http://127.0.0.1:8000/home',3008,'en-GB','Safari','N/A','2020-05-18 16:37:17','2020-05-18 16:37:17','',1),
	(22,'http://127.0.0.1:8000/register','http://127.0.0.1:8000/home',3008,'en-GB','Safari','N/A','2020-05-18 17:34:12','2020-05-18 17:34:12','',1),
	(23,'http://127.0.0.1:8000/home','http://127.0.0.1:8000/viewSite/1',3008,'en-GB','Safari','N/A','2020-05-18 17:35:00','2020-05-18 17:35:00','',1),
	(24,'http://127.0.0.1:8000/','http://127.0.0.1:8000/register',3008,'en-GB','Safari','N/A','2020-05-18 17:35:21','2020-05-18 17:35:21','',1),
	(25,'http://127.0.0.1:8000/register','http://127.0.0.1:8000/home',3008,'en-GB','Safari','N/A','2020-05-18 17:35:42','2020-05-18 17:35:42','',1),
	(26,'http://127.0.0.1:8000/register','http://127.0.0.1:8000/home',3008,'en-GB','Safari','N/A','2020-05-18 17:36:03','2020-05-18 17:36:03','',1),
	(27,'http://127.0.0.1:8000/home','http://127.0.0.1:8000/viewSite/4',3008,'en-GB','Safari','N/A','2020-05-18 17:36:06','2020-05-18 17:36:06','',1),
	(28,'','http://127.0.0.1:8000/home',3008,'en-GB','Safari','N/A','2020-05-18 17:36:11','2020-05-18 17:36:11','',1),
	(29,'http://127.0.0.1:8000/home','http://127.0.0.1:8000/viewSite/4',3008,'en-GB','Safari','N/A','2020-05-18 17:39:03','2020-05-18 17:39:03','',1),
	(30,'http://127.0.0.1:8000/','http://127.0.0.1:8000/login',3008,'en-GB','Safari','N/A','2020-05-18 21:03:32','2020-05-18 21:03:32','',1),
	(31,'http://127.0.0.1:8000/login','http://127.0.0.1:8000/home',3008,'en-GB','Safari','N/A','2020-05-18 21:03:35','2020-05-18 21:03:35','',1),
	(32,'http://127.0.0.1:8000/home','http://127.0.0.1:8000/viewSite/1',3008,'en-GB','Safari','N/A','2020-05-18 21:03:40','2020-05-18 21:03:40','',1),
	(33,'http://127.0.0.1:8000/home','http://127.0.0.1:8000/viewSite/1',3008,'en-GB','Safari','N/A','2020-05-18 21:19:14','2020-05-18 21:19:14','',1),
	(34,'http://127.0.0.1:8000/home','http://127.0.0.1:8000/viewSite/1',3008,'en-GB','Safari','N/A','2020-05-18 21:22:53','2020-05-18 21:22:53','',1),
	(35,'http://127.0.0.1:8000/home','http://127.0.0.1:8000/viewSite/1',3008,'en-GB','Safari','N/A','2020-05-18 21:23:43','2020-05-18 21:23:43','',1),
	(36,'http://127.0.0.1:8000/home','http://127.0.0.1:8000/viewSite/1',3008,'en-GB','Safari','N/A','2020-05-18 21:23:46','2020-05-18 21:23:46','',1),
	(37,'http://127.0.0.1:8000/home','http://127.0.0.1:8000/viewSite/1',3008,'en-GB','Safari','N/A','2020-05-18 21:24:37','2020-05-18 21:24:37','',1),
	(38,'','http://127.0.0.1:8000/viewSite/1',3008,'en-gb','Safari','N/A','2020-05-23 22:03:15','2020-05-23 22:03:15','',1),
	(39,'','http://127.0.0.1:8000/viewSite/1',3008,'en-gb','Safari','N/A','2020-05-23 22:10:35','2020-05-23 22:10:35','',1),
	(40,'','http://127.0.0.1:8000/viewSite/1',3008,'en-US','Chrome','N/A','2020-05-23 22:11:40','2020-05-23 22:11:40','',1),
	(41,'http://127.0.0.1:8000/','http://127.0.0.1:8000/login',3008,'en-gb','Safari','N/A','2020-05-24 09:42:53','2020-05-24 09:42:53','',1),
	(42,'http://127.0.0.1:8000/login','http://127.0.0.1:8000/home',3008,'en-gb','Safari','N/A','2020-05-24 09:42:55','2020-05-24 09:42:55','',1),
	(43,'http://127.0.0.1:8000/home','http://127.0.0.1:8000/viewSite/1',3008,'en-gb','Safari','N/A','2020-05-24 09:43:00','2020-05-24 09:43:00','',1),
	(44,'','http://127.0.0.1:8000/viewSite/1',3008,'en-gb','Safari','N/A','2020-05-24 09:43:28','2020-05-24 09:43:28','',1),
	(45,'http://127.0.0.1:8000/home','http://127.0.0.1:8000/viewSite/1',3008,'en-gb','Safari','N/A','2020-05-24 10:21:57','2020-05-24 10:21:57','',1),
	(46,'http://127.0.0.1:8000/home','http://127.0.0.1:8000/viewSite/1',3008,'en-gb','Safari','N/A','2020-05-24 10:22:52','2020-05-24 10:22:52','',1),
	(47,'http://127.0.0.1:8000/home','http://127.0.0.1:8000/viewSite/1',3008,'en-gb','Safari','N/A','2020-05-24 10:23:44','2020-05-24 10:23:44','',1),
	(48,'http://127.0.0.1:8000/home','http://127.0.0.1:8000/viewSite/1',3008,'en-gb','Safari','N/A','2020-05-24 10:37:14','2020-05-24 10:37:14','',1),
	(49,'http://127.0.0.1:8000/home','http://127.0.0.1:8000/viewSite/1',3008,'en-gb','Safari','N/A','2020-05-24 10:38:09','2020-05-24 10:38:09','',1),
	(50,'http://127.0.0.1:8000/home','http://127.0.0.1:8000/viewSite/1',3008,'en-gb','Safari','N/A','2020-05-24 10:41:58','2020-05-24 10:41:58','',1),
	(51,'http://127.0.0.1:8000/home','http://127.0.0.1:8000/viewSite/1',3008,'en-gb','Safari','N/A','2020-05-24 10:43:46','2020-05-24 10:43:46','',1),
	(52,'http://127.0.0.1:8000/home','http://127.0.0.1:8000/viewSite/1',3008,'en-gb','Safari','N/A','2020-05-24 10:51:28','2020-05-24 10:51:28','',1),
	(53,'http://127.0.0.1:8000/home','http://127.0.0.1:8000/viewSite/1',3008,'en-gb','Safari','N/A','2020-05-24 10:52:23','2020-05-24 10:52:23','',1),
	(54,'http://127.0.0.1:8000/home','http://127.0.0.1:8000/viewSite/1',3008,'en-gb','Safari','N/A','2020-05-24 10:53:54','2020-05-24 10:53:54','',1),
	(55,'http://127.0.0.1:8000/home','http://127.0.0.1:8000/viewSite/1',3008,'en-gb','Safari','N/A','2020-05-24 10:54:19','2020-05-24 10:54:19','',1),
	(56,'http://127.0.0.1:8000/home','http://127.0.0.1:8000/viewSite/1',3008,'en-gb','Safari','N/A','2020-05-24 10:54:37','2020-05-24 10:54:37','',1),
	(57,'http://127.0.0.1:8000/home','http://127.0.0.1:8000/viewSite/1',3008,'en-gb','Safari','N/A','2020-05-24 10:55:06','2020-05-24 10:55:06','',1),
	(58,'','http://127.0.0.1:8000/viewSite/1',1920,'en-gb','Safari','N/A','2020-05-24 10:55:16','2020-05-24 10:55:16','',1),
	(59,'','http://127.0.0.1:8000/viewSite/1',1920,'en-gb','Safari','N/A','2020-05-24 10:55:35','2020-05-24 10:55:35','',1),
	(60,'http://127.0.0.1:8000/','http://127.0.0.1:8000/home',3008,'en-gb','Safari','N/A','2020-05-24 11:23:23','2020-05-24 11:23:23','',1),
	(61,'http://127.0.0.1:8000/home','http://127.0.0.1:8000/viewSite/1',3008,'en-gb','Safari','N/A','2020-05-24 11:23:28','2020-05-24 11:23:28','',1),
	(62,'','http://127.0.0.1:8000/viewSite/1',3008,'en-gb','Safari','N/A','2020-05-24 11:24:43','2020-05-24 11:24:43','',1),
	(63,'','http://127.0.0.1:8000/viewSite/1',3008,'en-gb','Safari','N/A','2020-05-24 11:31:24','2020-05-24 11:31:24','',1),
	(64,'http://127.0.0.1:8000/home','http://127.0.0.1:8000/viewSite/1',3008,'en-gb','Safari','N/A','2020-05-24 11:45:46','2020-05-24 11:45:46','',1),
	(65,'','http://127.0.0.1:8000/viewSite/1',3008,'en-gb','Safari','N/A','2020-05-24 11:45:54','2020-05-24 11:45:54','',1),
	(66,'http://127.0.0.1:8000/home','http://127.0.0.1:8000/viewSite/1',3008,'en-gb','Safari','N/A','2020-05-24 11:47:33','2020-05-24 11:47:33','',1),
	(67,'','http://127.0.0.1:8000/viewSite/1',3008,'en-gb','Safari','N/A','2020-05-24 11:47:38','2020-05-24 11:47:38','',1),
	(68,'','http://127.0.0.1:8000/viewSite/1',3008,'en-gb','Safari','N/A','2020-05-24 12:42:33','2020-05-24 12:42:33','',1),
	(69,'','http://127.0.0.1:8000/viewSite/1',3008,'en-gb','Safari','N/A','2020-05-24 12:43:02','2020-05-24 12:43:02','',1),
	(70,'','http://127.0.0.1:8000/viewSite/1',3008,'en-gb','Safari','N/A','2020-05-24 12:43:46','2020-05-24 12:43:46','',1),
	(71,'http://127.0.0.1:8000/','http://127.0.0.1:8000/login',3008,'en-gb','Safari','N/A','2020-05-24 12:51:01','2020-05-24 12:51:01','',1),
	(72,'http://127.0.0.1:8000/home','http://127.0.0.1:8000/login',3008,'en-gb','Safari','N/A','2020-05-24 13:02:54','2020-05-24 13:02:54','',1),
	(73,'http://127.0.0.1:8000/login','http://127.0.0.1:8000/viewSite/1',3008,'en-gb','Safari','N/A','2020-05-24 13:02:58','2020-05-24 13:02:58','',1),
	(74,'http://127.0.0.1:8000/login','http://127.0.0.1:8000/viewSite/1',3008,'en-gb','Safari','N/A','2020-05-24 13:05:27','2020-05-24 13:05:27','',1),
	(75,'','http://127.0.0.1:8000/login',1920,'en-US','Chrome','N/A','2020-05-24 13:08:20','2020-05-24 13:08:20','',1),
	(76,'','http://127.0.0.1:8000/login',1920,'en-US','Chrome','N/A','2020-05-24 13:08:41','2020-05-24 13:08:41','',1);

/*!40000 ALTER TABLE `views` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table websites
# ------------------------------------------------------------

DROP TABLE IF EXISTS `websites`;

CREATE TABLE `websites` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `websiteURL` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `user_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `websites_user_id_foreign` (`user_id`),
  CONSTRAINT `websites_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `websites` WRITE;
/*!40000 ALTER TABLE `websites` DISABLE KEYS */;

INSERT INTO `websites` (`id`, `websiteURL`, `created_at`, `updated_at`, `user_id`)
VALUES
	(1,'http://127.0.0.1:8000',NULL,NULL,1),
	(2,'https://scottandrews.dev',NULL,NULL,1),
	(4,'https://bbc.com',NULL,NULL,2),
	(5,'https://test.com','2020-05-19 18:24:42','2020-05-19 18:24:42',1),
	(6,'https://test2.com','2020-05-19 18:25:06','2020-05-19 18:25:06',1),
	(7,'https://test3.com','2020-05-19 18:26:04','2020-05-19 18:26:04',1),
	(8,'https://www.newtownfootball.co.uk','2020-05-30 21:58:05','2020-05-30 21:58:05',1);

/*!40000 ALTER TABLE `websites` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table websockets_statistics_entries
# ------------------------------------------------------------

DROP TABLE IF EXISTS `websockets_statistics_entries`;

CREATE TABLE `websockets_statistics_entries` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `app_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `peak_connection_count` int(11) NOT NULL,
  `websocket_message_count` int(11) NOT NULL,
  `api_message_count` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
