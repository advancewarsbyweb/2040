-- MySQL dump 10.13  Distrib 5.5.47, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: awbw
-- ------------------------------------------------------
-- Server version	5.5.47-0+deb8u1-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `awbw_actions`
--


DROP TABLE IF EXISTS `awbw_actions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_actions` (
  `actions_id` int(11) NOT NULL AUTO_INCREMENT,
  `actions_games_id` int(11) DEFAULT NULL,
  `actions_object` varchar(8000) DEFAULT NULL,
  `actions_time` datetime DEFAULT NULL,
  `actions_players_id` int(11) DEFAULT NULL,
  `actions_day` int(11) DEFAULT NULL,
  PRIMARY KEY (`actions_id`),
  KEY `idx_actions_games_id` (`actions_games_id`),
  KEY `idx_actions_players_id` (`actions_players_id`)
) ENGINE=InnoDB AUTO_INCREMENT=190597948 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_banco`
--

DROP TABLE IF EXISTS `awbw_banco`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_banco` (
  `banco_id` int(11) NOT NULL AUTO_INCREMENT,
  `banco_games_id` int(11) DEFAULT NULL,
  `banco_co_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`banco_id`),
  KEY `banco_games_id_co_id` (`banco_games_id`,`banco_co_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6551864 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_banco_tags`
--

DROP TABLE IF EXISTS `awbw_banco_tags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_banco_tags` (
  `tags_banco_id` int(11) NOT NULL AUTO_INCREMENT,
  `tags_banco_games_id` int(11) DEFAULT NULL,
  `tags_banco_co_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`tags_banco_id`),
  KEY `banco_tags_games_id` (`tags_banco_games_id`)
) ENGINE=InnoDB AUTO_INCREMENT=59807 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_banunit`
--

DROP TABLE IF EXISTS `awbw_banunit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_banunit` (
  `banunit_id` int(11) NOT NULL AUTO_INCREMENT,
  `banunit_games_id` int(11) NOT NULL DEFAULT '0',
  `banunit_units_name` varchar(60) NOT NULL DEFAULT '0',
  PRIMARY KEY (`banunit_id`),
  KEY `banunit_games_id_units_name` (`banunit_games_id`,`banunit_units_name`)
) ENGINE=InnoDB AUTO_INCREMENT=940206 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_boots`
--

DROP TABLE IF EXISTS `awbw_boots`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_boots` (
  `boots_id` int(11) NOT NULL AUTO_INCREMENT,
  `boots_users_id` int(11) NOT NULL,
  `boots_games_id` int(11) NOT NULL,
  PRIMARY KEY (`boots_id`)
) ENGINE=InnoDB AUTO_INCREMENT=24923 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_buildings`
--

DROP TABLE IF EXISTS `awbw_buildings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_buildings` (
  `buildings_id` int(11) NOT NULL AUTO_INCREMENT,
  `buildings_games_id` int(11) DEFAULT NULL,
  `buildings_terrain_id` int(11) DEFAULT NULL,
  `buildings_x` int(11) DEFAULT NULL,
  `buildings_y` int(11) DEFAULT NULL,
  `buildings_capture` int(11) DEFAULT NULL,
  `buildings_last_capture` int(11) DEFAULT NULL,
  `buildings_last_updated` datetime DEFAULT NULL,
  PRIMARY KEY (`buildings_id`),
  KEY `buildings_games_id` (`buildings_games_id`,`buildings_x`,`buildings_y`),
  KEY `buildings_terrain_id` (`buildings_terrain_id`)
) ENGINE=MyISAM AUTO_INCREMENT=35841356 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_categories`
--

DROP TABLE IF EXISTS `awbw_categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_categories` (
  `categories_id` int(11) NOT NULL AUTO_INCREMENT,
  `categories_name` varchar(200) DEFAULT NULL,
  `categories_type` int(2) DEFAULT NULL,
  PRIMARY KEY (`categories_id`)
) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_changelog`
--

DROP TABLE IF EXISTS `awbw_changelog`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_changelog` (
  `changelog_id` int(11) NOT NULL AUTO_INCREMENT,
  `changelog_text` varchar(1000) DEFAULT NULL,
  `changelog_date` datetime DEFAULT NULL,
  PRIMARY KEY (`changelog_id`)
) ENGINE=InnoDB AUTO_INCREMENT=421 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_clanmembers`
--

DROP TABLE IF EXISTS `awbw_clanmembers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_clanmembers` (
  `clanmembers_clans_id` int(11) NOT NULL DEFAULT '0',
  `clanmembers_users_id` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`clanmembers_clans_id`,`clanmembers_users_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_clans`
--

DROP TABLE IF EXISTS `awbw_clans`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_clans` (
  `clans_id` int(11) NOT NULL AUTO_INCREMENT,
  `clans_name` varchar(200) DEFAULT NULL,
  `clans_users_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`clans_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_co`
--

DROP TABLE IF EXISTS `awbw_co`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_co` (
  `co_id` int(11) NOT NULL AUTO_INCREMENT,
  `co_name` varchar(40) DEFAULT NULL,
  `co_image` varchar(50) DEFAULT NULL,
  `co_countries_id` int(11) DEFAULT NULL,
  `co_max_power` int(11) NOT NULL DEFAULT '0',
  `co_max_spower` int(11) DEFAULT NULL,
  `co_power_name` varchar(50) NOT NULL DEFAULT '',
  `co_power_text` varchar(255) NOT NULL DEFAULT '',
  `co_spower_name` varchar(50) DEFAULT NULL,
  `co_spower_text` varchar(255) DEFAULT NULL,
  `co_text` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`co_id`),
  KEY `co_name` (`co_name`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_co_percent`
--

DROP TABLE IF EXISTS `awbw_co_percent`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_co_percent` (
  `co_percent_id` int(11) NOT NULL AUTO_INCREMENT,
  `co_percent_percent` int(11) NOT NULL DEFAULT '0',
  `co_percent_co_id` int(11) NOT NULL DEFAULT '0',
  `co_percent_units_id` int(11) NOT NULL DEFAULT '0',
  `co_percent_attack` char(1) NOT NULL DEFAULT '',
  `co_percent_power` char(1) NOT NULL DEFAULT '',
  PRIMARY KEY (`co_percent_id`),
  KEY `co_percent_co_id` (`co_percent_co_id`,`co_percent_units_id`)
) ENGINE=InnoDB AUTO_INCREMENT=138581 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_comments`
--

DROP TABLE IF EXISTS `awbw_comments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_comments` (
  `comments_id` int(11) NOT NULL AUTO_INCREMENT,
  `comments_users_id` int(11) DEFAULT NULL,
  `comments_maps_id` int(11) DEFAULT NULL,
  `comments_text` text,
  `comments_date` datetime DEFAULT NULL,
  `comments_last_edit_date` datetime DEFAULT NULL,
  PRIMARY KEY (`comments_id`),
  KEY `comments_maps_id` (`comments_maps_id`)
) ENGINE=MyISAM AUTO_INCREMENT=152857 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_cotiers`
--

DROP TABLE IF EXISTS `awbw_cotiers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_cotiers` (
  `cotiers_co_id` int(2) NOT NULL,
  `cotiers_tier` int(2) NOT NULL,
  `cotiers_game_type` varchar(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_countries`
--

DROP TABLE IF EXISTS `awbw_countries`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_countries` (
  `countries_id` int(11) NOT NULL AUTO_INCREMENT,
  `countries_name` varchar(40) DEFAULT NULL,
  `countries_code` char(2) DEFAULT NULL,
  `countries_dark` varchar(15) DEFAULT NULL,
  `countries_light` varchar(15) DEFAULT NULL,
  PRIMARY KEY (`countries_id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_damage`
--

DROP TABLE IF EXISTS `awbw_damage`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_damage` (
  `damage_id` int(11) NOT NULL AUTO_INCREMENT,
  `damage_attacker` int(11) DEFAULT NULL,
  `damage_defender` int(11) DEFAULT NULL,
  `damage_percentage` int(11) DEFAULT NULL,
  `damage_second_weapon` char(1) DEFAULT NULL,
  PRIMARY KEY (`damage_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1736 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_digest`
--

DROP TABLE IF EXISTS `awbw_digest`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_digest` (
  `digest_id` int(11) NOT NULL AUTO_INCREMENT,
  `digest_users_id` int(11) NOT NULL DEFAULT '0',
  `digest_games_id` int(11) NOT NULL DEFAULT '0',
  `digest_code` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`digest_id`),
  KEY `digest_users_id` (`digest_users_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_favorites`
--

DROP TABLE IF EXISTS `awbw_favorites`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_favorites` (
  `favorites_users_id` int(11) NOT NULL DEFAULT '0',
  `favorites_maps_id` int(11) NOT NULL DEFAULT '0',
  KEY `favorites_index` (`favorites_users_id`,`favorites_maps_id`),
  KEY `favorites_maps_id` (`favorites_maps_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_featmaps`
--

DROP TABLE IF EXISTS `awbw_featmaps`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_featmaps` (
  `featmaps_id` int(11) NOT NULL,
  `featmaps_date` date DEFAULT NULL,
  PRIMARY KEY (`featmaps_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_follows`
--

DROP TABLE IF EXISTS `awbw_follows`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_follows` (
  `follows_id` int(11) NOT NULL AUTO_INCREMENT,
  `follows_users_id` int(11) DEFAULT NULL,
  `follows_games_id` int(11) DEFAULT NULL,
  `follows_last_view` datetime DEFAULT NULL,
  PRIMARY KEY (`follows_id`),
  KEY `games_ids` (`follows_games_id`)
) ENGINE=InnoDB AUTO_INCREMENT=77007 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_friend_requests`
--

DROP TABLE IF EXISTS `awbw_friend_requests`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_friend_requests` (
  `requests_id` int(11) NOT NULL AUTO_INCREMENT,
  `requests_from_users_id` int(11) DEFAULT NULL,
  `requests_to_users_id` int(11) DEFAULT NULL,
  `requests_accepted` char(1) DEFAULT NULL,
  `requests_time` datetime DEFAULT NULL,
  PRIMARY KEY (`requests_id`),
  KEY `request_to_user` (`requests_to_users_id`),
  KEY `request_from_user` (`requests_from_users_id`)
) ENGINE=InnoDB AUTO_INCREMENT=38665 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_games`
--

DROP TABLE IF EXISTS `awbw_games`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_games` (
  `games_id` int(11) NOT NULL AUTO_INCREMENT,
  `games_name` varchar(100) DEFAULT NULL,
  `games_password` varchar(50) DEFAULT NULL,
  `games_creator` int(11) DEFAULT NULL,
  `games_start_date` datetime DEFAULT NULL,
  `games_end_date` date DEFAULT NULL,
  `games_activity_date` datetime DEFAULT NULL,
  `games_maps_id` int(11) DEFAULT NULL,
  `games_weather_type` varchar(10) DEFAULT NULL,
  `games_weather_start` int(11) DEFAULT NULL,
  `games_weather_code` char(1) DEFAULT NULL,
  `games_private` char(1) DEFAULT NULL,
  `games_turn` int(11) DEFAULT NULL,
  `games_day` int(11) DEFAULT NULL,
  `games_active` char(1) DEFAULT NULL,
  `games_funds` int(11) DEFAULT NULL,
  `games_capture_win` int(11) DEFAULT NULL,
  `games_days` int(11) DEFAULT '0',
  `games_fog` char(1) DEFAULT NULL,
  `games_comment` text,
  `games_type` char(1) DEFAULT NULL,
  `games_boot_interval` int(11) NOT NULL DEFAULT '0',
  `games_starting_funds` int(11) NOT NULL DEFAULT '0',
  `games_official` char(1) DEFAULT NULL,
  `games_min_rating` int(11) DEFAULT NULL,
  `games_unit_limit` int(11) DEFAULT NULL,
  `games_league` char(1) DEFAULT NULL,
  `games_team` char(1) NOT NULL DEFAULT 'N',
  `games_aet_interval` int(11) DEFAULT NULL,
  `games_aet_date` datetime DEFAULT NULL,
  `games_use_powers` char(1) DEFAULT NULL,
  PRIMARY KEY (`games_id`),
  KEY `games_maps` (`games_id`,`games_maps_id`),
  KEY `games_maps_id` (`games_maps_id`),
  KEY `games_turn` (`games_turn`),
  KEY `games_creator` (`games_creator`),
  KEY `games_end_date` (`games_end_date`),
  KEY `games_type` (`games_type`)
) ENGINE=MyISAM AUTO_INCREMENT=731167 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_games_test`
--

DROP TABLE IF EXISTS `awbw_games_test`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_games_test` (
  `games_id` int(11) NOT NULL AUTO_INCREMENT,
  `games_name` varchar(100) DEFAULT NULL,
  `games_password` varchar(50) DEFAULT NULL,
  `games_creator` int(11) DEFAULT NULL,
  `games_start_date` date DEFAULT NULL,
  `games_end_date` date DEFAULT NULL,
  `games_activity_date` datetime DEFAULT NULL,
  `games_maps_id` int(11) DEFAULT NULL,
  `games_weather_type` varchar(10) DEFAULT NULL,
  `games_weather_start` int(11) DEFAULT NULL,
  `games_weather_code` char(1) DEFAULT NULL,
  `games_win_condition` char(2) DEFAULT NULL,
  `games_turn` int(11) DEFAULT NULL,
  `games_day` int(11) DEFAULT NULL,
  `games_active` char(1) DEFAULT NULL,
  `games_funds` int(11) DEFAULT NULL,
  `games_capture_win` int(11) DEFAULT NULL,
  `games_days` int(11) DEFAULT '0',
  `games_fog` char(1) DEFAULT NULL,
  `games_comment` text,
  `games_type` char(1) DEFAULT NULL,
  `games_boot_interval` int(11) NOT NULL DEFAULT '0',
  `games_starting_funds` int(11) NOT NULL DEFAULT '0',
  `games_official` char(1) DEFAULT NULL,
  `games_min_rating` int(11) DEFAULT NULL,
  `games_max_rating` int(11) DEFAULT NULL,
  `games_league` char(1) DEFAULT NULL,
  `games_team` char(1) NOT NULL DEFAULT 'N',
  `games_aet_interval` int(11) DEFAULT NULL,
  `games_aet_date` datetime DEFAULT NULL,
  `games_use_powers` char(1) DEFAULT NULL,
  PRIMARY KEY (`games_id`),
  KEY `games_maps` (`games_id`,`games_maps_id`),
  KEY `games_maps_id` (`games_maps_id`),
  KEY `games_turn` (`games_turn`)
) ENGINE=InnoDB AUTO_INCREMENT=149210 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_games_timers`
--

DROP TABLE IF EXISTS `awbw_games_timers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_games_timers` (
  `timers_id` int(11) NOT NULL AUTO_INCREMENT,
  `timers_games_id` int(11) NOT NULL,
  `timers_max_turn` int(11) DEFAULT NULL,
  `timers_initial` int(11) DEFAULT NULL,
  `timers_increment` int(11) DEFAULT NULL,
  `timers_pause_time` datetime DEFAULT NULL,
  PRIMARY KEY (`timers_id`),
  KEY `games_ids` (`timers_games_id`)
) ENGINE=InnoDB AUTO_INCREMENT=441174 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_gauntlets`
--

DROP TABLE IF EXISTS `awbw_gauntlets`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_gauntlets` (
  `gauntlets_id` int(11) NOT NULL AUTO_INCREMENT,
  `gauntlets_name` varchar(200) DEFAULT NULL,
  `gauntlets_start_date` datetime DEFAULT NULL,
  PRIMARY KEY (`gauntlets_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_gauntlets_levels`
--

DROP TABLE IF EXISTS `awbw_gauntlets_levels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_gauntlets_levels` (
  `levels_gauntlets_id` int(11) NOT NULL,
  `levels_level` int(11) NOT NULL,
  `levels_maps_id` int(11) NOT NULL,
  `levels_users_id` int(11) NOT NULL,
  `levels_co_id` int(11) NOT NULL,
  KEY `levels_gauntlets_level` (`levels_gauntlets_id`,`levels_level`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_gauntlets_players`
--

DROP TABLE IF EXISTS `awbw_gauntlets_players`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_gauntlets_players` (
  `players_gauntlets_id` int(11) NOT NULL,
  `players_users_id` int(11) NOT NULL,
  `players_level` int(11) DEFAULT NULL,
  KEY `players_gauntlets_users` (`players_gauntlets_id`,`players_users_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_history`
--

DROP TABLE IF EXISTS `awbw_history`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_history` (
  `history_users_id` int(11) NOT NULL DEFAULT '0',
  `history_units_id` int(11) NOT NULL DEFAULT '0',
  `history_type` char(1) NOT NULL DEFAULT '',
  `history_number` int(11) DEFAULT NULL,
  PRIMARY KEY (`history_users_id`,`history_units_id`,`history_type`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_history_games`
--

DROP TABLE IF EXISTS `awbw_history_games`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_history_games` (
  `history_users_id` int(11) NOT NULL,
  `history_games_id` int(11) NOT NULL,
  `history_units_id` int(11) NOT NULL,
  `history_type` char(1) NOT NULL,
  `history_number` int(11) DEFAULT NULL,
  PRIMARY KEY (`history_users_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_ip_bans`
--

DROP TABLE IF EXISTS `awbw_ip_bans`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_ip_bans` (
  `bans_id` int(11) NOT NULL AUTO_INCREMENT,
  `bans_ip` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`bans_id`)
) ENGINE=InnoDB AUTO_INCREMENT=145 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_ips`
--

DROP TABLE IF EXISTS `awbw_ips`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_ips` (
  `ips_id` int(11) NOT NULL AUTO_INCREMENT,
  `ip` varchar(16) DEFAULT NULL,
  `users_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`ips_id`),
  KEY `users_id` (`users_id`)
) ENGINE=InnoDB AUTO_INCREMENT=848169 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_labunit`
--

DROP TABLE IF EXISTS `awbw_labunit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_labunit` (
  `labunit_id` int(11) NOT NULL AUTO_INCREMENT,
  `labunit_games_id` int(11) DEFAULT NULL,
  `labunit_units_name` varchar(60) DEFAULT NULL,
  PRIMARY KEY (`labunit_id`),
  KEY `labunit_games_id` (`labunit_games_id`)
) ENGINE=InnoDB AUTO_INCREMENT=404488 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_league2_games`
--

DROP TABLE IF EXISTS `awbw_league2_games`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_league2_games` (
  `games_id` int(11) NOT NULL,
  PRIMARY KEY (`games_id`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_league2_prefs`
--

DROP TABLE IF EXISTS `awbw_league2_prefs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_league2_prefs` (
  `prefs_id` int(11) NOT NULL,
  PRIMARY KEY (`prefs_id`),
  CONSTRAINT `awbw_league2_prefs_ibfk_1` FOREIGN KEY (`prefs_id`) REFERENCES `awbw_prefs_games` (`prefs_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_league2_stats`
--

DROP TABLE IF EXISTS `awbw_league2_stats`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_league2_stats` (
  `stats_id` int(11) NOT NULL AUTO_INCREMENT,
  `stats_users_id` int(11) DEFAULT NULL,
  `stats_games_type` varchar(11) DEFAULT NULL,
  `stats_wins` int(11) DEFAULT NULL,
  `stats_losses` int(11) DEFAULT NULL,
  `stats_draws` int(11) DEFAULT NULL,
  PRIMARY KEY (`stats_id`),
  KEY `stats_users_id` (`stats_users_id`),
  KEY `stats_games_type` (`stats_games_type`)
) ENGINE=InnoDB AUTO_INCREMENT=44470 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_league2_users`
--

DROP TABLE IF EXISTS `awbw_league2_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_league2_users` (
  `users_id` int(11) NOT NULL,
  `league2_users_concurrent_games` int(11) NOT NULL DEFAULT '0',
  `users_game_type` int(2) DEFAULT NULL,
  `users_last_played` int(11) DEFAULT NULL,
  PRIMARY KEY (`users_id`),
  KEY `league2_users_games` (`league2_users_concurrent_games`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_leagues`
--

DROP TABLE IF EXISTS `awbw_leagues`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_leagues` (
  `leagues_id` int(11) NOT NULL AUTO_INCREMENT,
  `leagues_name` varchar(100) DEFAULT NULL,
  `leagues_start_date` datetime DEFAULT NULL,
  `leagues_boot_threshold` int(11) DEFAULT NULL,
  `leagues_divisions` int(11) DEFAULT NULL,
  `leagues_rain_chance` int(11) DEFAULT NULL,
  `leagues_snow_chance` int(11) DEFAULT NULL,
  `leagues_random_chance` int(11) DEFAULT NULL,
  `leagues_capture_win` int(11) DEFAULT '1000',
  `leagues_days` int(11) DEFAULT '0',
  `leagues_fog_chance` int(11) DEFAULT NULL,
  `leagues_boot_interval` int(11) DEFAULT NULL,
  `leagues_aet_interval` int(11) DEFAULT NULL,
  `leagues_use_powers` char(1) DEFAULT 'Y',
  `leagues_funds` int(11) DEFAULT NULL,
  `leagues_starting_funds` int(11) DEFAULT NULL,
  `leagues_active` int(11) DEFAULT '0',
  `leagues_allow_join` int(11) DEFAULT '1',
  `leagues_admin` int(11) DEFAULT NULL,
  `leagues_tag` int(11) DEFAULT '0',
  PRIMARY KEY (`leagues_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_leagues_banco`
--

DROP TABLE IF EXISTS `awbw_leagues_banco`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_leagues_banco` (
  `co_leagues_id` int(11) NOT NULL DEFAULT '0',
  `co_leagues_co_id` int(11) NOT NULL DEFAULT '0',
  KEY `co_key` (`co_leagues_id`,`co_leagues_co_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_leagues_bancountry`
--

DROP TABLE IF EXISTS `awbw_leagues_bancountry`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_leagues_bancountry` (
  `country_leagues_id` int(11) DEFAULT NULL,
  `country_leagues_countries_id` int(11) DEFAULT NULL,
  KEY `country_league_id` (`country_leagues_id`,`country_leagues_countries_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_leagues_banunit`
--

DROP TABLE IF EXISTS `awbw_leagues_banunit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_leagues_banunit` (
  `unit_leagues_id` int(11) NOT NULL DEFAULT '0',
  `unit_leagues_unit_id` int(11) NOT NULL DEFAULT '0',
  KEY `unit_key` (`unit_leagues_id`,`unit_leagues_unit_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_leagues_games`
--

DROP TABLE IF EXISTS `awbw_leagues_games`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_leagues_games` (
  `games_games_id` int(11) NOT NULL DEFAULT '0',
  `games_leagues_id` int(11) NOT NULL DEFAULT '0',
  KEY `games_key` (`games_games_id`,`games_leagues_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_leagues_labunit`
--

DROP TABLE IF EXISTS `awbw_leagues_labunit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_leagues_labunit` (
  `unit_leagues_id` int(11) NOT NULL DEFAULT '0',
  `unit_leagues_unit_id` int(11) NOT NULL DEFAULT '0',
  KEY `unit_league` (`unit_leagues_id`,`unit_leagues_unit_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_leagues_maps`
--

DROP TABLE IF EXISTS `awbw_leagues_maps`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_leagues_maps` (
  `maps_maps_id` int(11) NOT NULL DEFAULT '0',
  `maps_leagues_id` int(11) NOT NULL DEFAULT '0',
  `maps_approved` int(11) DEFAULT '1',
  KEY `maps_key` (`maps_maps_id`,`maps_leagues_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_leagues_mapvotes`
--

DROP TABLE IF EXISTS `awbw_leagues_mapvotes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_leagues_mapvotes` (
  `mapvotes_leagues_id` int(11) NOT NULL DEFAULT '0',
  `mapvotes_maps_id` int(11) NOT NULL DEFAULT '0',
  `mapvotes_users_id` int(11) NOT NULL DEFAULT '0',
  `mapvotes_vote` int(11) DEFAULT NULL,
  KEY `prim` (`mapvotes_leagues_id`,`mapvotes_maps_id`,`mapvotes_users_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_leagues_players`
--

DROP TABLE IF EXISTS `awbw_leagues_players`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_leagues_players` (
  `players_users_id` int(11) NOT NULL DEFAULT '0',
  `players_leagues_id` int(11) NOT NULL DEFAULT '0',
  `players_division` int(11) NOT NULL DEFAULT '0',
  `players_games_won` int(11) NOT NULL DEFAULT '0',
  `players_games_lost` int(11) NOT NULL DEFAULT '0',
  `players_games_drawn` int(11) NOT NULL DEFAULT '0',
  `players_booted` int(11) NOT NULL DEFAULT '0',
  `players_rating` double(8,2) DEFAULT '800.00',
  `players_allow_games` int(11) DEFAULT '1',
  `players_games_missed` int(11) DEFAULT '0',
  `players_co_image` varchar(100) DEFAULT NULL,
  KEY `players_leagues_key` (`players_users_id`,`players_leagues_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_logs`
--

DROP TABLE IF EXISTS `awbw_logs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_logs` (
  `logs_id` int(11) NOT NULL AUTO_INCREMENT,
  `logs_games_id` int(11) DEFAULT NULL,
  `logs_date` datetime DEFAULT NULL,
  `logs_text` text,
  `logs_day` int(11) NOT NULL DEFAULT '0',
  `logs_countries_id` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`logs_id`),
  KEY `logs_games_id` (`logs_games_id`)
) ENGINE=MyISAM AUTO_INCREMENT=441254815 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_mapcats`
--

DROP TABLE IF EXISTS `awbw_mapcats`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_mapcats` (
  `mapcats_maps_id` int(11) NOT NULL DEFAULT '0',
  `mapcats_categories_id` int(11) NOT NULL DEFAULT '0',
  KEY `prim` (`mapcats_maps_id`,`mapcats_categories_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_mapcats_log`
--

DROP TABLE IF EXISTS `awbw_mapcats_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_mapcats_log` (
  `log_id` int(11) NOT NULL AUTO_INCREMENT,
  `log_maps_id` int(11) NOT NULL,
  `log_users_id` int(11) DEFAULT NULL,
  `log_time` datetime DEFAULT NULL,
  PRIMARY KEY (`log_id`)
) ENGINE=InnoDB AUTO_INCREMENT=12183 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_mapcats_log_list`
--

DROP TABLE IF EXISTS `awbw_mapcats_log_list`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_mapcats_log_list` (
  `mapcats_log_id` int(11) NOT NULL,
  `mapcats_category_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_maps`
--

DROP TABLE IF EXISTS `awbw_maps`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_maps` (
  `maps_id` int(11) NOT NULL AUTO_INCREMENT,
  `maps_name` varchar(100) DEFAULT NULL,
  `maps_players` int(11) DEFAULT NULL,
  `maps_type` char(1) DEFAULT NULL,
  `maps_users_id` int(11) DEFAULT NULL,
  `maps_published` char(1) DEFAULT NULL,
  `maps_published_date` datetime DEFAULT NULL,
  `maps_first_published_date` datetime DEFAULT NULL,
  `maps_awn_number` int(11) DEFAULT NULL,
  `maps_last_viewed` date DEFAULT NULL,
  PRIMARY KEY (`maps_id`),
  KEY `maps_users_id` (`maps_users_id`)
) ENGINE=MyISAM AUTO_INCREMENT=132237 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_maps_size`
--

DROP TABLE IF EXISTS `awbw_maps_size`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_maps_size` (
  `size_maps_id` int(11) NOT NULL,
  `size_max_x` int(11) DEFAULT NULL,
  `size_max_y` int(11) DEFAULT NULL,
  PRIMARY KEY (`size_maps_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_maps_staging`
--

DROP TABLE IF EXISTS `awbw_maps_staging`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_maps_staging` (
  `maps_id` int(11) NOT NULL AUTO_INCREMENT,
  `maps_name` varchar(60) DEFAULT NULL,
  `maps_players` int(11) DEFAULT NULL,
  `maps_type` char(1) DEFAULT NULL,
  `maps_users_id` int(11) DEFAULT NULL,
  `maps_published` char(1) DEFAULT NULL,
  `maps_published_date` datetime DEFAULT NULL,
  `maps_first_published_date` datetime DEFAULT NULL,
  `maps_awn_number` int(11) DEFAULT NULL,
  `maps_awn_last_updated` date DEFAULT NULL,
  PRIMARY KEY (`maps_id`)
) ENGINE=MyISAM AUTO_INCREMENT=3530 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_maps_test`
--

DROP TABLE IF EXISTS `awbw_maps_test`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_maps_test` (
  `maps_id` int(11) NOT NULL AUTO_INCREMENT,
  `maps_name` varchar(100) DEFAULT NULL,
  `maps_players` int(11) DEFAULT NULL,
  `maps_type` char(1) DEFAULT NULL,
  `maps_users_id` int(11) DEFAULT NULL,
  `maps_published` char(1) DEFAULT NULL,
  `maps_published_date` datetime DEFAULT NULL,
  `maps_first_published_date` datetime DEFAULT NULL,
  `maps_awn_number` int(11) DEFAULT NULL,
  `maps_awn_last_updated` date DEFAULT NULL,
  PRIMARY KEY (`maps_id`),
  KEY `maps_users_id` (`maps_users_id`)
) ENGINE=InnoDB AUTO_INCREMENT=55004 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_mapunits`
--

DROP TABLE IF EXISTS `awbw_mapunits`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_mapunits` (
  `mapunits_maps_id` int(11) NOT NULL DEFAULT '0',
  `mapunits_units_id` int(11) NOT NULL DEFAULT '0',
  `mapunits_count` int(11) DEFAULT NULL,
  KEY `maps_units` (`mapunits_maps_id`,`mapunits_units_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_marquees`
--

DROP TABLE IF EXISTS `awbw_marquees`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_marquees` (
  `marquees_id` int(11) NOT NULL AUTO_INCREMENT,
  `marquees_event_id` int(11) NOT NULL DEFAULT '0',
  `marquees_type` char(1) DEFAULT NULL,
  `marquees_active` int(11) DEFAULT NULL,
  `marquees_text` text,
  PRIMARY KEY (`marquees_id`),
  KEY `event_marquee` (`marquees_event_id`,`marquees_type`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_messages`
--

DROP TABLE IF EXISTS `awbw_messages`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_messages` (
  `messages_id` int(11) NOT NULL AUTO_INCREMENT,
  `messages_text` text,
  `messages_to_users_id` int(11) DEFAULT NULL,
  `messages_from_users_id` int(11) DEFAULT NULL,
  `messages_date` datetime DEFAULT NULL,
  `messages_subject` varchar(255) DEFAULT NULL,
  `messages_read` char(1) DEFAULT NULL,
  PRIMARY KEY (`messages_id`),
  KEY `messages_to` (`messages_to_users_id`),
  KEY `messages_from` (`messages_from_users_id`)
) ENGINE=MyISAM AUTO_INCREMENT=373120 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_movement_costs`
--

DROP TABLE IF EXISTS `awbw_movement_costs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_movement_costs` (
  `movement_costs_id` int(11) NOT NULL AUTO_INCREMENT,
  `movement_costs_cost` int(11) DEFAULT NULL,
  `movement_costs_terrain_id` int(11) DEFAULT NULL,
  `movement_costs_type_code` char(1) DEFAULT NULL,
  `movement_costs_weather_code` char(1) DEFAULT NULL,
  PRIMARY KEY (`movement_costs_id`),
  KEY `movement_costs_terrain_id_weather_code` (`movement_costs_terrain_id`,`movement_costs_weather_code`)
) ENGINE=InnoDB AUTO_INCREMENT=5052 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_newleague_banlist`
--

DROP TABLE IF EXISTS `awbw_newleague_banlist`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_newleague_banlist` (
  `users_id` int(11) NOT NULL,
  PRIMARY KEY (`users_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_notifications`
--

DROP TABLE IF EXISTS `awbw_notifications`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_notifications` (
  `notifications_users_id` int(11) NOT NULL,
  `notifications_email` char(1) DEFAULT NULL,
  `notifications_discord` char(1) DEFAULT NULL,
  `notifications_create` char(1) DEFAULT NULL,
  `notifications_join` char(1) DEFAULT NULL,
  `notifications_start` char(1) DEFAULT NULL,
  `notifications_end` char(1) DEFAULT NULL,
  `notifications_boot` int(3) DEFAULT NULL,
  `notifications_message` char(1) DEFAULT NULL,
  `notifications_newturn` char(1) DEFAULT NULL,
  PRIMARY KEY (`notifications_users_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_password_resets`
--

DROP TABLE IF EXISTS `awbw_password_resets`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_password_resets` (
  `reset_users_id` int(11) DEFAULT NULL,
  `reset_token` varchar(25) DEFAULT NULL,
  `reset_token_used` tinyint(4) DEFAULT NULL,
  `reset_created` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_players`
--

DROP TABLE IF EXISTS `awbw_players`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_players` (
  `players_id` int(11) NOT NULL AUTO_INCREMENT,
  `players_users_id` int(11) DEFAULT NULL,
  `players_games_id` int(11) DEFAULT NULL,
  `players_countries_id` int(11) DEFAULT NULL,
  `players_co_id` int(11) DEFAULT NULL,
  `players_funds` int(11) DEFAULT NULL,
  `players_turn` int(11) DEFAULT NULL,
  `players_old_interface` char(1) DEFAULT NULL,
  `players_uniq_id` varchar(30) DEFAULT NULL,
  `players_eliminated` char(1) DEFAULT NULL,
  `players_last_read` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `players_turn_start` datetime DEFAULT NULL,
  `players_turn_clock` int(11) DEFAULT NULL,
  `players_pause_timer` int(1) DEFAULT NULL,
  `players_co_power` int(11) NOT NULL DEFAULT '0',
  `players_co_power_on` char(1) DEFAULT NULL,
  `players_order` int(11) DEFAULT NULL,
  `players_accept_draw` char(1) DEFAULT NULL,
  `players_co_max_power` int(11) DEFAULT NULL,
  `players_co_max_spower` int(11) DEFAULT NULL,
  `players_co_image` varchar(30) DEFAULT NULL,
  `players_team` varchar(10) DEFAULT NULL,
  `players_aet_count` int(11) DEFAULT '0',
  `players_turn_flag` int(1) DEFAULT '0',
  PRIMARY KEY (`players_id`),
  KEY `players_games_id_users_id` (`players_games_id`,`players_users_id`),
  KEY `players_games_id_co_id` (`players_games_id`,`players_co_id`),
  KEY `players_co_users` (`players_users_id`,`players_co_id`),
  KEY `players_id_games` (`players_id`,`players_games_id`)
) ENGINE=MyISAM AUTO_INCREMENT=1805549 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_players_test`
--

DROP TABLE IF EXISTS `awbw_players_test`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_players_test` (
  `players_id` int(11) NOT NULL AUTO_INCREMENT,
  `players_users_id` int(11) DEFAULT NULL,
  `players_games_id` int(11) DEFAULT NULL,
  `players_countries_id` int(11) DEFAULT NULL,
  `players_co_id` int(11) DEFAULT NULL,
  `players_funds` int(11) DEFAULT NULL,
  `players_turn` int(11) DEFAULT NULL,
  `players_email` char(1) DEFAULT NULL,
  `players_uniq_id` varchar(30) DEFAULT NULL,
  `players_eliminated` char(1) DEFAULT NULL,
  `players_last_read` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `players_last_read_broadcasts` datetime DEFAULT NULL,
  `players_emailpress` char(1) NOT NULL DEFAULT '',
  `players_signature` varchar(255) NOT NULL DEFAULT '',
  `players_co_power` int(11) NOT NULL DEFAULT '0',
  `players_co_power_on` char(1) DEFAULT NULL,
  `players_order` int(11) DEFAULT NULL,
  `players_accept_draw` char(1) DEFAULT NULL,
  `players_co_max_power` int(11) DEFAULT NULL,
  `players_co_max_spower` int(11) DEFAULT NULL,
  `players_co_image` varchar(30) DEFAULT NULL,
  `players_team` varchar(10) DEFAULT NULL,
  `players_aet_count` int(11) DEFAULT '0',
  PRIMARY KEY (`players_id`),
  KEY `players_games_id_users_id` (`players_games_id`,`players_users_id`),
  KEY `players_games_id_co_id` (`players_games_id`,`players_co_id`),
  KEY `players_co_users` (`players_users_id`,`players_co_id`),
  KEY `players_id_games` (`players_id`,`players_games_id`)
) ENGINE=InnoDB AUTO_INCREMENT=420760 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_predeployed`
--

DROP TABLE IF EXISTS `awbw_predeployed`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_predeployed` (
  `predeployed_id` int(11) NOT NULL AUTO_INCREMENT,
  `predeployed_maps_id` int(11) DEFAULT NULL,
  `predeployed_units_id` int(11) DEFAULT NULL,
  `predeployed_x` int(11) DEFAULT NULL,
  `predeployed_y` int(11) DEFAULT NULL,
  `predeployed_countries_code` char(2) DEFAULT NULL,
  PRIMARY KEY (`predeployed_id`),
  KEY `predeployed_x_y` (`predeployed_maps_id`,`predeployed_x`,`predeployed_y`)
) ENGINE=InnoDB AUTO_INCREMENT=4855307 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_prefs_banco`
--

DROP TABLE IF EXISTS `awbw_prefs_banco`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_prefs_banco` (
  `banco_id` int(11) NOT NULL AUTO_INCREMENT,
  `banco_prefs_id` int(11) DEFAULT NULL,
  `banco_co_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`banco_id`)
) ENGINE=InnoDB AUTO_INCREMENT=171657 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_prefs_banco_tags`
--

DROP TABLE IF EXISTS `awbw_prefs_banco_tags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_prefs_banco_tags` (
  `tags_banco_id` int(11) NOT NULL AUTO_INCREMENT,
  `tags_banco_prefs_id` int(11) DEFAULT NULL,
  `tags_banco_co_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`tags_banco_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5460 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_prefs_banunit`
--

DROP TABLE IF EXISTS `awbw_prefs_banunit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_prefs_banunit` (
  `banunit_id` int(11) NOT NULL AUTO_INCREMENT,
  `banunit_prefs_id` int(11) NOT NULL DEFAULT '0',
  `banunit_units_name` varchar(60) NOT NULL DEFAULT '0',
  PRIMARY KEY (`banunit_id`),
  KEY `banunit_prefs_id_units_name` (`banunit_prefs_id`,`banunit_units_name`)
) ENGINE=InnoDB AUTO_INCREMENT=40573 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_prefs_games`
--

DROP TABLE IF EXISTS `awbw_prefs_games`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_prefs_games` (
  `prefs_id` int(11) NOT NULL AUTO_INCREMENT,
  `prefs_users_id` int(11) NOT NULL DEFAULT '0',
  `prefs_name` varchar(40) DEFAULT NULL,
  `prefs_password` varchar(50) DEFAULT NULL,
  `prefs_max_turn` int(11) DEFAULT NULL,
  `prefs_initial` int(11) DEFAULT NULL,
  `prefs_increment` int(11) DEFAULT NULL,
  `prefs_activity_date` datetime DEFAULT NULL,
  `prefs_maps_id` int(11) DEFAULT NULL,
  `prefs_weather_type` varchar(10) DEFAULT NULL,
  `prefs_weather_start` int(11) DEFAULT NULL,
  `prefs_weather_code` char(1) DEFAULT NULL,
  `prefs_win_condition` char(2) DEFAULT NULL,
  `prefs_turn` int(11) DEFAULT NULL,
  `prefs_day` int(11) DEFAULT NULL,
  `prefs_active` char(1) DEFAULT NULL,
  `prefs_funds` int(11) DEFAULT NULL,
  `prefs_capture_win` int(11) DEFAULT NULL,
  `prefs_days` int(11) DEFAULT '0',
  `prefs_fog` char(1) DEFAULT NULL,
  `prefs_comment` text,
  `prefs_type` char(1) DEFAULT NULL,
  `prefs_boot_interval` int(11) NOT NULL DEFAULT '0',
  `prefs_starting_funds` int(11) NOT NULL DEFAULT '0',
  `prefs_official` char(1) DEFAULT NULL,
  `prefs_min_rating` int(11) DEFAULT NULL,
  `prefs_unit_limit` int(11) DEFAULT NULL,
  `prefs_tags` char(1) DEFAULT NULL,
  `prefs_team` char(1) NOT NULL DEFAULT 'N',
  `prefs_aet_interval` int(11) DEFAULT NULL,
  PRIMARY KEY (`prefs_id`)
) ENGINE=InnoDB AUTO_INCREMENT=12437 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_prefs_labunit`
--

DROP TABLE IF EXISTS `awbw_prefs_labunit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_prefs_labunit` (
  `labunit_id` int(11) NOT NULL AUTO_INCREMENT,
  `labunit_prefs_id` int(11) DEFAULT NULL,
  `labunit_units_name` varchar(60) DEFAULT NULL,
  PRIMARY KEY (`labunit_id`)
) ENGINE=InnoDB AUTO_INCREMENT=24668 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_press`
--

DROP TABLE IF EXISTS `awbw_press`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_press` (
  `press_id` int(11) NOT NULL AUTO_INCREMENT,
  `press_text` text,
  `press_players_id` int(11) DEFAULT NULL,
  `press_date` datetime DEFAULT NULL,
  `press_subject` varchar(255) DEFAULT NULL,
  `press_type` char(1) DEFAULT NULL,
  PRIMARY KEY (`press_id`)
) ENGINE=InnoDB AUTO_INCREMENT=545848 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_press_to`
--

DROP TABLE IF EXISTS `awbw_press_to`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_press_to` (
  `press_to_id` int(11) NOT NULL AUTO_INCREMENT,
  `press_to_press_id` int(11) DEFAULT NULL,
  `press_to_players_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`press_to_id`),
  KEY `press_to_press_id` (`press_to_press_id`),
  KEY `press_to_players_id` (`press_to_players_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1271235 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_ratings`
--

DROP TABLE IF EXISTS `awbw_ratings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_ratings` (
  `ratings_id` int(11) NOT NULL AUTO_INCREMENT,
  `ratings_users_id` int(11) DEFAULT NULL,
  `ratings_maps_id` int(11) DEFAULT NULL,
  `ratings_score` int(11) DEFAULT NULL,
  PRIMARY KEY (`ratings_id`),
  KEY `ratings_maps_id` (`ratings_maps_id`)
) ENGINE=MyISAM AUTO_INCREMENT=137936 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_readies`
--

DROP TABLE IF EXISTS `awbw_readies`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_readies` (
  `readies_id` int(11) NOT NULL AUTO_INCREMENT,
  `readies_games_id` int(11) DEFAULT NULL,
  `readies_players_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`readies_id`),
  KEY `players_ids` (`readies_players_id`)
) ENGINE=InnoDB AUTO_INCREMENT=594050 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_replays`
--

DROP TABLE IF EXISTS `awbw_replays`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_replays` (
  `replays_id` int(11) NOT NULL AUTO_INCREMENT,
  `replays_games_id` int(11) DEFAULT NULL,
  `replays_users_id` int(11) DEFAULT NULL,
  `replays_save_time` datetime DEFAULT NULL,
  PRIMARY KEY (`replays_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5200 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_saved_replays`
--

DROP TABLE IF EXISTS `awbw_saved_replays`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_saved_replays` (
  `games_id` int(11) NOT NULL,
  PRIMARY KEY (`games_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_states`
--

DROP TABLE IF EXISTS `awbw_states`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_states` (
  `states_id` int(11) NOT NULL AUTO_INCREMENT,
  `states_users_id` int(11) DEFAULT NULL,
  `states_games_id` int(11) DEFAULT NULL,
  `states_turn` int(11) DEFAULT NULL,
  `states_time` datetime DEFAULT NULL,
  PRIMARY KEY (`states_id`),
  KEY `idx_states_users_id` (`states_users_id`),
  KEY `idx_states_games_id` (`states_games_id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_stats`
--

DROP TABLE IF EXISTS `awbw_stats`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_stats` (
  `stats_id` int(11) NOT NULL AUTO_INCREMENT,
  `stats_users_id` int(11) DEFAULT NULL,
  `stats_games_played` int(11) DEFAULT NULL,
  `stats_games_won` int(11) DEFAULT NULL,
  `stats_games_lost` int(11) DEFAULT NULL,
  `stats_games_drawn` int(11) DEFAULT NULL,
  `stats_games_booted` int(11) DEFAULT NULL,
  `stats_rating` double(8,2) DEFAULT NULL,
  PRIMARY KEY (`stats_id`),
  KEY `stats_users_id` (`stats_users_id`),
  KEY `stats_rating` (`stats_rating`)
) ENGINE=MyISAM AUTO_INCREMENT=33948 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_stats_ratings`
--

DROP TABLE IF EXISTS `awbw_stats_ratings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_stats_ratings` (
  `ratings_id` int(11) NOT NULL AUTO_INCREMENT,
  `users_id` int(11) NOT NULL,
  `games_id` int(11) DEFAULT NULL,
  `users_rating` double(8,2) DEFAULT NULL,
  `ratings_games_type` varchar(11) DEFAULT NULL,
  PRIMARY KEY (`ratings_id`),
  KEY `games_ids` (`games_id`),
  KEY `users_id` (`users_id`),
  KEY `games_type` (`ratings_games_type`)
) ENGINE=InnoDB AUTO_INCREMENT=811180 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_symmetry`
--

DROP TABLE IF EXISTS `awbw_symmetry`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_symmetry` (
  `terrain_id` int(11) NOT NULL,
  `symmetry_type` varchar(255) NOT NULL,
  `symmetry_terrain_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_tags`
--

DROP TABLE IF EXISTS `awbw_tags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_tags` (
  `tags_players_id` int(11) NOT NULL AUTO_INCREMENT,
  `tags_games_id` int(11) DEFAULT NULL,
  `tags_co_id` int(11) DEFAULT NULL,
  `tags_co_power` int(11) DEFAULT NULL,
  `tags_co_power_on` char(1) DEFAULT NULL,
  `tags_co_max_power` int(11) DEFAULT NULL,
  `tags_co_max_spower` int(11) DEFAULT NULL,
  `tags_co_image` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`tags_players_id`),
  KEY `tags_games_id` (`tags_games_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1805514 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_terrain`
--

DROP TABLE IF EXISTS `awbw_terrain`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_terrain` (
  `terrain_id` int(11) NOT NULL AUTO_INCREMENT,
  `terrain_name` varchar(30) DEFAULT NULL,
  `terrain_defense` int(11) DEFAULT NULL,
  `terrain_symbol` char(1) DEFAULT NULL,
  `terrain_country_code` char(2) DEFAULT NULL,
  `terrain_build_type` char(1) DEFAULT NULL,
  `terrain_active` char(1) NOT NULL DEFAULT 'N',
  `terrain_offset` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`terrain_id`)
) ENGINE=InnoDB AUTO_INCREMENT=195 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_tiles`
--

DROP TABLE IF EXISTS `awbw_tiles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_tiles` (
  `tiles_id` int(11) NOT NULL AUTO_INCREMENT,
  `tiles_maps_id` int(11) DEFAULT NULL,
  `tiles_terrain_id` int(11) DEFAULT NULL,
  `tiles_x` int(11) DEFAULT NULL,
  `tiles_y` int(11) DEFAULT NULL,
  `tiles_type` char(1) DEFAULT NULL,
  PRIMARY KEY (`tiles_id`),
  KEY `tiles_x_y` (`tiles_maps_id`,`tiles_x`,`tiles_y`),
  KEY `tiles_maps_terrain` (`tiles_maps_id`,`tiles_terrain_id`)
) ENGINE=InnoDB AUTO_INCREMENT=96794133 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_tiles_myisam`
--

DROP TABLE IF EXISTS `awbw_tiles_myisam`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_tiles_myisam` (
  `tiles_id` int(11) NOT NULL AUTO_INCREMENT,
  `tiles_maps_id` int(11) DEFAULT NULL,
  `tiles_terrain_id` int(11) DEFAULT NULL,
  `tiles_x` int(11) DEFAULT NULL,
  `tiles_y` int(11) DEFAULT NULL,
  `tiles_type` char(1) DEFAULT NULL,
  PRIMARY KEY (`tiles_id`),
  KEY `tiles_x_y` (`tiles_maps_id`,`tiles_x`,`tiles_y`),
  KEY `tiles_maps_terrain` (`tiles_maps_id`,`tiles_terrain_id`)
) ENGINE=InnoDB AUTO_INCREMENT=76203 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_tiles_staging`
--

DROP TABLE IF EXISTS `awbw_tiles_staging`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_tiles_staging` (
  `tiles_id` int(11) NOT NULL AUTO_INCREMENT,
  `tiles_maps_id` int(11) DEFAULT NULL,
  `tiles_terrain_id` int(11) DEFAULT NULL,
  `tiles_x` int(11) DEFAULT NULL,
  `tiles_y` int(11) DEFAULT NULL,
  `tiles_type` char(1) DEFAULT NULL,
  PRIMARY KEY (`tiles_id`),
  KEY `tiles_x_y` (`tiles_maps_id`,`tiles_x`,`tiles_y`)
) ENGINE=MyISAM AUTO_INCREMENT=2407348 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_tiles_test`
--

DROP TABLE IF EXISTS `awbw_tiles_test`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_tiles_test` (
  `tiles_id` int(11) NOT NULL AUTO_INCREMENT,
  `tiles_maps_id` int(11) DEFAULT NULL,
  `tiles_terrain_id` int(11) DEFAULT NULL,
  `tiles_x` int(11) DEFAULT NULL,
  `tiles_y` int(11) DEFAULT NULL,
  `tiles_type` char(1) DEFAULT NULL,
  PRIMARY KEY (`tiles_id`),
  KEY `tiles_x_y` (`tiles_maps_id`,`tiles_x`,`tiles_y`),
  KEY `tiles_maps_terrain` (`tiles_maps_id`,`tiles_terrain_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3871 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_tournaments`
--

DROP TABLE IF EXISTS `awbw_tournaments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_tournaments` (
  `tournaments_id` int(11) NOT NULL AUTO_INCREMENT,
  `tournaments_name` varchar(100) DEFAULT NULL,
  `tournaments_shortname` varchar(15) DEFAULT NULL,
  `tournaments_players` int(11) DEFAULT NULL,
  `tournaments_elim` int(1) DEFAULT NULL,
  `tournaments_creator` int(11) DEFAULT NULL,
  `tournaments_prefs_id` int(11) DEFAULT NULL,
  `tournaments_active` char(1) DEFAULT NULL,
  `tournaments_winner_users_id` int(11) DEFAULT NULL,
  `tournaments_create` datetime DEFAULT NULL,
  `tournaments_start` datetime DEFAULT NULL,
  `tournaments_comment` varchar(1000) DEFAULT NULL,
  `tournaments_status` int(11) DEFAULT NULL,
  `tournaments_min_rating` double(8,2) DEFAULT NULL,
  `tournaments_min_games` int(11) DEFAULT NULL,
  `tournaments_type` char(1) DEFAULT NULL,
  `tournaments_end` datetime DEFAULT NULL,
  `tournaments_bracket_start` int(11) DEFAULT NULL,
  `tournaments_bracket_finals` int(11) DEFAULT NULL,
  PRIMARY KEY (`tournaments_id`)
) ENGINE=InnoDB AUTO_INCREMENT=227 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_tournaments_banco`
--

DROP TABLE IF EXISTS `awbw_tournaments_banco`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_tournaments_banco` (
  `banco_id` int(11) NOT NULL AUTO_INCREMENT,
  `banco_tournaments_id` int(11) DEFAULT NULL,
  `banco_tournaments_round` int(11) DEFAULT NULL,
  `banco_co_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`banco_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7324 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_tournaments_banco_tags`
--

DROP TABLE IF EXISTS `awbw_tournaments_banco_tags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_tournaments_banco_tags` (
  `banco_id` int(11) NOT NULL AUTO_INCREMENT,
  `banco_tournaments_id` int(11) DEFAULT NULL,
  `banco_tournaments_round` int(11) DEFAULT NULL,
  `banco_co_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`banco_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Table structure for table `awbw_tournaments_banlist`
--

DROP TABLE IF EXISTS `awbw_tournaments_banlist`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_tournaments_banlist` (
  `banlist_users_id` int(11) NOT NULL,
  `ban_count` int(11) DEFAULT NULL,
  PRIMARY KEY (`banlist_users_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_tournaments_banunit`
--

DROP TABLE IF EXISTS `awbw_tournaments_banunit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_tournaments_banunit` (
  `banunit_id` int(11) NOT NULL AUTO_INCREMENT,
  `banunit_tournaments_id` int(11) DEFAULT NULL,
  `banunit_tournaments_round` int(11) DEFAULT NULL,
  `banunit_units_name` varchar(60) DEFAULT NULL,
  PRIMARY KEY (`banunit_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1026 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_tournaments_games`
--

DROP TABLE IF EXISTS `awbw_tournaments_games`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_tournaments_games` (
  `tournaments_games_id` int(11) NOT NULL AUTO_INCREMENT,
  `tournaments_games_tournaments_id` int(11) DEFAULT NULL,
  `tournaments_games_games_id` int(11) DEFAULT NULL,
  `tournaments_games_round` int(11) DEFAULT NULL,
  `tournaments_games_number` int(11) DEFAULT NULL,
  `tournaments_games_ready` char(1) DEFAULT NULL,
  PRIMARY KEY (`tournaments_games_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2875 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_tournaments_players`
--

DROP TABLE IF EXISTS `awbw_tournaments_players`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_tournaments_players` (
  `tournaments_players_id` int(11) NOT NULL AUTO_INCREMENT,
  `tournaments_players_tournaments_id` int(11) DEFAULT NULL,
  `tournaments_players_users_id` int(11) DEFAULT NULL,
  `tournaments_players_seed` int(11) DEFAULT NULL,
  `tournaments_players_elim` char(1) DEFAULT NULL,
  PRIMARY KEY (`tournaments_players_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4466 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_tournaments_settings`
--

DROP TABLE IF EXISTS `awbw_tournaments_settings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_tournaments_settings` (
  `settings_id` int(11) NOT NULL AUTO_INCREMENT,
  `settings_tournaments_id` int(11) DEFAULT NULL,
  `settings_tournaments_round` int(2) DEFAULT NULL,
  `settings_maps_id` int(11) DEFAULT NULL,
  `settings_days` int(3) DEFAULT NULL,
  `settings_capture` int(3) DEFAULT NULL,
  PRIMARY KEY (`settings_id`)
) ENGINE=InnoDB AUTO_INCREMENT=511 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_tourneys`
--

DROP TABLE IF EXISTS `awbw_tourneys`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_tourneys` (
  `tourneys_id` int(11) NOT NULL AUTO_INCREMENT,
  `tourneys_name` varchar(100) DEFAULT NULL,
  `tourneys_shortname` varchar(15) DEFAULT NULL,
  `tourneys_players` int(11) DEFAULT NULL,
  `tourneys_prelim_type` int(1) DEFAULT NULL,
  `tourneys_finals_type` int(1) DEFAULT NULL,
  `tourneys_admin` int(11) DEFAULT NULL,
  `tourneys_visible` char(1) DEFAULT NULL,
  `tourneys_status` int(1) DEFAULT NULL,
  `tourneys_winner` int(11) DEFAULT NULL,
  `tourneys_create` datetime DEFAULT NULL,
  `tourneys_start` datetime DEFAULT NULL,
  `tourneys_end` datetime DEFAULT NULL,
  `tourneys_comment` varchar(1000) DEFAULT NULL,
  `tourneys_min_rating` double(8,2) DEFAULT NULL,
  `tourneys_min_games` int(11) DEFAULT NULL,
  `tourneys_type` char(1) DEFAULT NULL,
  `tourneys_pools_count` int(11) DEFAULT NULL,
  `tourneys_pools_adv_next` int(1) DEFAULT NULL,
  PRIMARY KEY (`tourneys_id`)
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_tourneys_games`
--

DROP TABLE IF EXISTS `awbw_tourneys_games`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_tourneys_games` (
  `games_id` int(11) NOT NULL AUTO_INCREMENT,
  `games_tourneys_id` int(11) DEFAULT NULL,
  `games_pools_id` int(11) DEFAULT NULL,
  `games_num` int(11) DEFAULT NULL,
  `games_games_id` int(11) DEFAULT NULL,
  `games_round` int(11) DEFAULT NULL,
  `games_ready` char(1) DEFAULT NULL,
  PRIMARY KEY (`games_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4350 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_tourneys_mods`
--

DROP TABLE IF EXISTS `awbw_tourneys_mods`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_tourneys_mods` (
  `mods_tourneys_id` int(11) NOT NULL,
  `mods_users_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_tourneys_players`
--

DROP TABLE IF EXISTS `awbw_tourneys_players`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_tourneys_players` (
  `players_id` int(11) NOT NULL AUTO_INCREMENT,
  `players_tourneys_id` int(11) DEFAULT NULL,
  `players_users_id` int(11) DEFAULT NULL,
  `players_seed` int(11) DEFAULT NULL,
  `players_elim` char(1) DEFAULT NULL,
  `players_pools_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`players_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1755 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_tourneys_pools`
--

DROP TABLE IF EXISTS `awbw_tourneys_pools`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_tourneys_pools` (
  `pools_id` int(11) NOT NULL AUTO_INCREMENT,
  `pools_tourneys_id` int(11) DEFAULT NULL,
  `pools_num` int(11) DEFAULT NULL,
  `pools_elim_type` int(1) DEFAULT NULL,
  `pools_finals` char(1) DEFAULT NULL,
  `pools_name` varchar(40) DEFAULT NULL,
  `pools_old_id` int(11) DEFAULT NULL,
  `pools_players` int(11) DEFAULT NULL,
  PRIMARY KEY (`pools_id`)
) ENGINE=InnoDB AUTO_INCREMENT=232 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_tourneys_seats`
--

DROP TABLE IF EXISTS `awbw_tourneys_seats`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_tourneys_seats` (
  `seats_id` int(11) NOT NULL AUTO_INCREMENT,
  `seats_tourneys_id` int(11) DEFAULT NULL,
  `seats_pools_id` int(11) DEFAULT NULL,
  `seats_round` int(11) DEFAULT NULL,
  `seats_num` int(11) DEFAULT NULL,
  `seats_games_num` int(11) DEFAULT NULL,
  `seats_from_game_num` int(11) DEFAULT NULL,
  `seats_from_type` char(1) DEFAULT NULL,
  `seats_users_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`seats_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5067 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_tourneys_settings`
--

DROP TABLE IF EXISTS `awbw_tourneys_settings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_tourneys_settings` (
  `settings_id` int(11) NOT NULL AUTO_INCREMENT,
  `settings_tourneys_id` int(11) DEFAULT NULL,
  `settings_tourneys_round` int(11) DEFAULT NULL,
  `settings_prefs_id` int(11) DEFAULT NULL,
  `settings_pools_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`settings_id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_tournaments_banco`
--

DROP TABLE IF EXISTS `awbw_tutorial_videos`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_tutorial_videos` (
  `tutorials_id` int(11) NOT NULL AUTO_INCREMENT,
  `tutorials_creator` varchar(100) DEFAULT NULL,
  `tutorials_title` varchar(1000) DEFAULT NULL,
  `tutorials_url` varchar(100) DEFAULT NULL,
  `tutorials_order` int(11) DEFAULT NULL,
  PRIMARY KEY (`tutorials_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

--
-- Table structure for table `awbw_units`
--

DROP TABLE IF EXISTS `awbw_units`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_units` (
  `units_id` int(11) NOT NULL AUTO_INCREMENT,
  `units_games_id` int(11) DEFAULT NULL,
  `units_players_id` int(11) DEFAULT NULL,
  `units_name` varchar(30) DEFAULT NULL,
  `units_movement_points` int(11) DEFAULT NULL,
  `units_vision` int(11) DEFAULT NULL,
  `units_fuel` int(11) DEFAULT NULL,
  `units_fuel_per_turn` int(11) DEFAULT NULL,
  `units_sub_dive` char(1) DEFAULT NULL,
  `units_ammo` int(11) DEFAULT NULL,
  `units_short_range` int(11) DEFAULT NULL,
  `units_long_range` int(11) DEFAULT NULL,
  `units_second_weapon` char(1) DEFAULT NULL,
  `units_symbol` char(1) DEFAULT NULL,
  `units_cost` int(11) DEFAULT NULL,
  `units_movement_type` char(1) DEFAULT NULL,
  `units_x` int(11) DEFAULT NULL,
  `units_y` int(11) DEFAULT NULL,
  `units_moved` int(11) DEFAULT NULL,
  `units_capture` int(11) DEFAULT NULL,
  `units_fired` int(11) DEFAULT NULL,
  `units_hit_points` double(3,1) DEFAULT NULL,
  `units_cargo1_units_id` int(11) NOT NULL DEFAULT '0',
  `units_cargo2_units_id` int(11) NOT NULL DEFAULT '0',
  `units_carried` char(1) NOT NULL DEFAULT '',
  PRIMARY KEY (`units_id`),
  KEY `units_players_id_games_id_x_y` (`units_players_id`,`units_games_id`,`units_x`,`units_y`),
  KEY `units_x_y` (`units_x`,`units_y`),
  KEY `units_games_id_players_id` (`units_games_id`,`units_players_id`),
  KEY `units_games_id` (`units_games_id`,`units_x`,`units_y`),
  KEY `units_cargo2` (`units_cargo2_units_id`),
  KEY `units_cargo1` (`units_cargo1_units_id`)
) ENGINE=MyISAM AUTO_INCREMENT=115521606 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_units_discovered`
--

DROP TABLE IF EXISTS `awbw_units_discovered`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_units_discovered` (
  `discovered_games_id` int(11) NOT NULL,
  `discovered_units_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_units_staging`
--

DROP TABLE IF EXISTS `awbw_units_staging`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_units_staging` (
  `units_id` int(11) NOT NULL AUTO_INCREMENT,
  `units_games_id` int(11) DEFAULT NULL,
  `units_players_id` int(11) DEFAULT NULL,
  `units_name` varchar(30) DEFAULT NULL,
  `units_movement_points` int(11) DEFAULT NULL,
  `units_vision` int(11) DEFAULT NULL,
  `units_fuel` int(11) DEFAULT NULL,
  `units_fuel_per_turn` int(11) DEFAULT NULL,
  `units_sub_dive` char(1) DEFAULT NULL,
  `units_ammo` int(11) DEFAULT NULL,
  `units_short_range` int(11) DEFAULT NULL,
  `units_long_range` int(11) DEFAULT NULL,
  `units_second_weapon` char(1) DEFAULT NULL,
  `units_symbol` char(1) DEFAULT NULL,
  `units_cost` int(11) DEFAULT NULL,
  `units_movement_type` char(1) DEFAULT NULL,
  `units_x` int(11) DEFAULT NULL,
  `units_y` int(11) DEFAULT NULL,
  `units_moved` int(11) DEFAULT NULL,
  `units_capture` int(11) DEFAULT NULL,
  `units_fired` int(11) DEFAULT NULL,
  `units_hit_points` int(11) DEFAULT NULL,
  `units_cargo1_units_id` int(11) NOT NULL DEFAULT '0',
  `units_cargo2_units_id` int(11) NOT NULL DEFAULT '0',
  `units_carried` char(1) NOT NULL DEFAULT '',
  PRIMARY KEY (`units_id`),
  KEY `units_players_id_games_id_x_y` (`units_players_id`,`units_games_id`,`units_x`,`units_y`),
  KEY `units_x_y` (`units_x`,`units_y`),
  KEY `units_games_id_players_id` (`units_games_id`,`units_players_id`)
) ENGINE=MyISAM AUTO_INCREMENT=1141439 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_units_stats`
--

DROP TABLE IF EXISTS `awbw_units_stats`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_units_stats` (
  `units_stats_id` int(11) NOT NULL AUTO_INCREMENT,
  `units_stats_players_id` int(11) NOT NULL,
  `units_stats_name` varchar(30) DEFAULT NULL,
  `units_stats_kills` int(11) DEFAULT NULL,
  `units_stats_destroyed` int(11) DEFAULT NULL,
  `units_stats_funds` int(11) DEFAULT NULL,
  PRIMARY KEY (`units_stats_id`),
  KEY `stats_players_id` (`units_stats_players_id`),
  KEY `stats_units_name` (`units_stats_name`)
) ENGINE=InnoDB AUTO_INCREMENT=6160603 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_units_test`
--

DROP TABLE IF EXISTS `awbw_units_test`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_units_test` (
  `units_id` int(11) NOT NULL AUTO_INCREMENT,
  `units_games_id` int(11) DEFAULT NULL,
  `units_players_id` int(11) DEFAULT NULL,
  `units_name` varchar(30) DEFAULT NULL,
  `units_movement_points` int(11) DEFAULT NULL,
  `units_vision` int(11) DEFAULT NULL,
  `units_fuel` int(11) DEFAULT NULL,
  `units_fuel_per_turn` int(11) DEFAULT NULL,
  `units_sub_dive` char(1) DEFAULT NULL,
  `units_ammo` int(11) DEFAULT NULL,
  `units_short_range` int(11) DEFAULT NULL,
  `units_long_range` int(11) DEFAULT NULL,
  `units_second_weapon` char(1) DEFAULT NULL,
  `units_symbol` char(1) DEFAULT NULL,
  `units_cost` int(11) DEFAULT NULL,
  `units_movement_type` char(1) DEFAULT NULL,
  `units_x` int(11) DEFAULT NULL,
  `units_y` int(11) DEFAULT NULL,
  `units_moved` int(11) DEFAULT NULL,
  `units_capture` int(11) DEFAULT NULL,
  `units_fired` int(11) DEFAULT NULL,
  `units_hit_points` double(3,1) DEFAULT NULL,
  `units_cargo1_units_id` int(11) NOT NULL DEFAULT '0',
  `units_cargo2_units_id` int(11) NOT NULL DEFAULT '0',
  `units_carried` char(1) NOT NULL DEFAULT '',
  PRIMARY KEY (`units_id`),
  KEY `units_players_id_games_id_x_y` (`units_players_id`,`units_games_id`,`units_x`,`units_y`),
  KEY `units_x_y` (`units_x`,`units_y`),
  KEY `units_games_id_players_id` (`units_games_id`,`units_players_id`),
  KEY `units_name` (`units_name`)
) ENGINE=InnoDB AUTO_INCREMENT=637129 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_vision`
--

DROP TABLE IF EXISTS `awbw_vision`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_vision` (
  `vision_players_id` int(11) DEFAULT NULL,
  `vision_games_id` int(11) DEFAULT NULL,
  `vision_requester_id` int(11) DEFAULT NULL,
  `vision_request_status` int(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awbw_zgames`
--

DROP TABLE IF EXISTS `awbw_zgames`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_zgames` (
  `zgames_id` int(11) NOT NULL AUTO_INCREMENT,
  `zgames_games_id` int(11) DEFAULT NULL,
  `zgames_ready` char(1) DEFAULT NULL,
  PRIMARY KEY (`zgames_id`)
) ENGINE=InnoDB AUTO_INCREMENT=70219 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `bankick`
--

DROP TABLE IF EXISTS `bankick`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `bankick` (
  `id` bigint(255) NOT NULL AUTO_INCREMENT,
  `ip` varchar(16) NOT NULL DEFAULT '10.0.0.1',
  `type` varchar(100) NOT NULL DEFAULT '',
  `timestamp` int(14) NOT NULL DEFAULT '0',
  `port` int(5) NOT NULL DEFAULT '0',
  `reason` varchar(200) NOT NULL DEFAULT '',
  `nickname` varchar(14) DEFAULT NULL,
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `ip` (`ip`),
  KEY `port` (`port`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `boxes`
--

DROP TABLE IF EXISTS `boxes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `boxes` (
  `id` bigint(60) NOT NULL AUTO_INCREMENT,
  `window_title` varchar(255) NOT NULL DEFAULT '',
  `last_activity` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `uid` bigint(255) NOT NULL DEFAULT '0',
  `port` bigint(5) NOT NULL DEFAULT '0',
  `submit_ip` varchar(255) DEFAULT NULL,
  `customername` varchar(255) NOT NULL DEFAULT '-',
  `customeraddress` varchar(255) DEFAULT NULL,
  `customerpostcode` varchar(50) DEFAULT NULL,
  `customerplaats` varchar(50) DEFAULT NULL,
  `customerphone` varchar(30) DEFAULT NULL,
  `customeremail` varchar(255) DEFAULT NULL,
  `customerurl` varchar(255) DEFAULT NULL,
  `background` varchar(50) NOT NULL DEFAULT '',
  `profile_url` varchar(255) DEFAULT NULL,
  `adminfullname` varchar(50) NOT NULL DEFAULT '',
  `adminname` varchar(50) NOT NULL DEFAULT '',
  `adminpass` varchar(50) NOT NULL DEFAULT '',
  `name` varchar(50) NOT NULL DEFAULT '',
  `maxusers` bigint(20) DEFAULT '15',
  `numberofrooms` int(3) NOT NULL DEFAULT '0',
  `boxtype` set('free','lite','medium','heavy','pro') NOT NULL DEFAULT 'free',
  `userlist` set('yes','no') NOT NULL DEFAULT 'yes',
  `userdisplay` varchar(50) DEFAULT NULL,
  `entername` varchar(50) DEFAULT NULL,
  `chattext` text,
  `afterlogintext` text,
  `standardsendtext` text,
  `socketconnectfailuretext` text,
  `bannedtext` text,
  `nickisinusetext` text,
  `chatsessionended` text,
  `serverfulltext` text,
  `socketclosetext` text,
  `unknownfailuretext` text,
  `flushemployeetext` text,
  `waitsecs` text,
  `logouttext` text,
  `chatSessionHolded` text,
  `welcomeText` text,
  `deparray` text,
  `backpicture` text,
  `logopicture` text,
  `maxoperators` bigint(20) NOT NULL DEFAULT '10',
  `standardthemecolor` text,
  `enterfoto` text,
  `startdate` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `soundtriggerchannel` set('on','off') DEFAULT NULL,
  `charsize` set('10','12','14','15','16') DEFAULT '12',
  `helpurl` varchar(255) DEFAULT NULL,
  `active` set('yes','no') NOT NULL DEFAULT 'yes',
  `reactivate_code` varchar(10) NOT NULL DEFAULT '',
  `themetrigger` set('on','off') DEFAULT 'on',
  `entercv` varchar(255) DEFAULT NULL,
  `warning_youwerebanned` varchar(255) DEFAULT NULL,
  `warning_youwerekicked` varchar(255) DEFAULT NULL,
  `language` varchar(255) DEFAULT NULL,
  `freebox` set('true','false') NOT NULL DEFAULT 'true',
  `newbox` set('true','false') NOT NULL DEFAULT 'true',
  `webcambox` set('true','false') NOT NULL DEFAULT 'false',
  `featured` set('true','false') NOT NULL DEFAULT 'false',
  `profile_location` varchar(255) NOT NULL DEFAULT '',
  `profile_picture_location` varchar(255) NOT NULL DEFAULT '',
  `operator_unban` set('true','false') NOT NULL DEFAULT 'true',
  `allow_guest` set('yes','no') NOT NULL DEFAULT 'yes',
  PRIMARY KEY (`id`,`customername`,`port`),
  UNIQUE KEY `port` (`port`),
  KEY `port_2` (`port`)
) ENGINE=MyISAM AUTO_INCREMENT=29 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `log`
--

DROP TABLE IF EXISTS `log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user` varchar(50) NOT NULL DEFAULT '',
  `code` varchar(50) NOT NULL DEFAULT '',
  `text` text,
  `port` varchar(50) NOT NULL DEFAULT '',
  `time` bigint(255) DEFAULT NULL,
  `referrer` text,
  `timestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `ip` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `ofua_users`
--

DROP TABLE IF EXISTS `ofua_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ofua_users` (
  `users_id` int(11) NOT NULL AUTO_INCREMENT,
  `users_username` varchar(30) DEFAULT NULL,
  `users_firstname` varchar(20) DEFAULT NULL,
  `users_lastname` varchar(20) DEFAULT NULL,
  `users_password` varchar(100) DEFAULT NULL,
  `users_email` varchar(60) DEFAULT NULL,
  `users_email_display` enum('Y','N') DEFAULT NULL,
  `users_inactive` int(11) DEFAULT NULL,
  `users_games` int(11) DEFAULT NULL,
  `users_uniq_id` varchar(30) DEFAULT NULL,
  `users_discord_id` char(20) DEFAULT NULL,
  `users_session_id` varchar(60) DEFAULT NULL,
  `users_activity_date` datetime DEFAULT NULL,
  `users_boot_record` int(11) DEFAULT NULL,
  `pn_name` varchar(60) NOT NULL DEFAULT '',
  `pn_uname` varchar(25) NOT NULL DEFAULT '',
  `pn_email` varchar(60) NOT NULL DEFAULT '',
  `pn_femail` varchar(60) NOT NULL DEFAULT '',
  `pn_url` varchar(254) NOT NULL DEFAULT '',
  `pn_user_avatar` varchar(30) DEFAULT NULL,
  `pn_user_regdate` varchar(20) NOT NULL DEFAULT '',
  `pn_user_icq` varchar(15) DEFAULT NULL,
  `pn_user_occ` varchar(100) DEFAULT NULL,
  `pn_user_from` varchar(100) DEFAULT NULL,
  `pn_user_intrest` varchar(150) DEFAULT NULL,
  `pn_user_sig` varchar(255) DEFAULT NULL,
  `pn_user_viewemail` tinyint(2) DEFAULT NULL,
  `pn_user_theme` tinyint(3) DEFAULT NULL,
  `pn_user_aim` varchar(18) DEFAULT NULL,
  `pn_user_yim` varchar(25) DEFAULT NULL,
  `pn_user_msnm` varchar(25) DEFAULT NULL,
  `pn_pass` varchar(40) NOT NULL DEFAULT '',
  `pn_storynum` tinyint(4) DEFAULT '10',
  `pn_umode` varchar(10) NOT NULL DEFAULT '',
  `pn_uorder` tinyint(1) DEFAULT '0',
  `pn_thold` tinyint(1) DEFAULT '0',
  `pn_noscore` tinyint(1) DEFAULT '0',
  `pn_bio` tinytext NOT NULL,
  `pn_ublockon` tinyint(1) DEFAULT '0',
  `pn_ublock` text NOT NULL,
  `pn_theme` varchar(255) NOT NULL DEFAULT '',
  `pn_commentmax` int(11) DEFAULT '4096',
  `pn_counter` int(11) DEFAULT '0',
  `pn_timezone_offset` float(3,1) DEFAULT '0.0',
  `pn_uid` int(11) NOT NULL DEFAULT '0',
  `users_shoals_display` varchar(30) DEFAULT NULL,
  `users_countries_id` int(11) DEFAULT NULL,
  `users_awbw_yim` varchar(30) DEFAULT NULL,
  `users_awbw_msn` varchar(30) DEFAULT NULL,
  `users_awbw_image_dir` text,
  `users_friend_date` datetime DEFAULT NULL,
  `users_auto_scroll` int(11) DEFAULT NULL,
  `users_changelog_date` datetime DEFAULT NULL,
  `ip` varchar(16) NOT NULL DEFAULT '',
  `users_awbw_last_page` varchar(60) NOT NULL DEFAULT '',
  `users_awbw_moderator` char(1) NOT NULL DEFAULT '',
  `users_awbw_days` int(11) NOT NULL DEFAULT '0',
  `users_yourgames_date` datetime DEFAULT NULL,
  `users_awbw_turn_email` varchar(10) NOT NULL DEFAULT '',
  `users_timezone` varchar(5) NOT NULL DEFAULT '-4',
  `users_email_messages` char(1) NOT NULL DEFAULT '',
  `users_show_warnings` int(11) DEFAULT NULL,
  `users_co_portraits` varchar(10) DEFAULT NULL,
  `users_tournament_date` datetime DEFAULT NULL,
  `users_gridlines` char(3) DEFAULT NULL,
  `users_vacation` int(11) DEFAULT NULL,
  `users_last_vacation` datetime DEFAULT NULL,
  `users_map_committee` int(11) DEFAULT '0',
  `users_dor_fc` varchar(30) DEFAULT NULL,
  `users_game_animations` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`users_id`),
  KEY `users_password` (`users_password`,`users_id`),
  KEY `users_username` (`users_username`),
  KEY `users_inactive` (`users_inactive`)
) ENGINE=MyISAM AUTO_INCREMENT=310426 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `online`
--

DROP TABLE IF EXISTS `online`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `online` (
  `port` int(11) NOT NULL DEFAULT '0',
  `user` varchar(50) NOT NULL DEFAULT '',
  `room` varchar(50) NOT NULL DEFAULT '',
  `isop` int(1) NOT NULL DEFAULT '0'
) ENGINE=MyISAM DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `site_games`
--

DROP TABLE IF EXISTS `site_games`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `site_games` (
  `games_id` int(11) DEFAULT NULL,
  `games_name` varchar(30) DEFAULT NULL,
  `games_code` varchar(4) DEFAULT NULL,
  `games_is_running` int(11) DEFAULT NULL,
  `games_version` varchar(10) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-10-14 15:30:48

-- MySQL dump 10.13  Distrib 5.5.47, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: awbw
-- ------------------------------------------------------
-- Server version	5.5.47-0+deb8u1-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `awbw_co`
--

DROP TABLE IF EXISTS `awbw_co`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_co` (
  `co_id` int(11) NOT NULL AUTO_INCREMENT,
  `co_name` varchar(40) DEFAULT NULL,
  `co_image` varchar(50) DEFAULT NULL,
  `co_countries_id` int(11) DEFAULT NULL,
  `co_max_power` int(11) NOT NULL DEFAULT '0',
  `co_max_spower` int(11) DEFAULT NULL,
  `co_power_name` varchar(50) NOT NULL DEFAULT '',
  `co_power_text` varchar(255) NOT NULL DEFAULT '',
  `co_spower_name` varchar(50) DEFAULT NULL,
  `co_spower_text` varchar(255) DEFAULT NULL,
  `co_text` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`co_id`),
  KEY `co_name` (`co_name`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `awbw_co`
--

LOCK TABLES `awbw_co` WRITE;
/*!40000 ALTER TABLE `awbw_co` DISABLE KEYS */;
INSERT INTO `awbw_co` VALUES (1,'Andy','andy.png',1,270000,540000,'Hyper Repair','All units gain +2 HP.','Hyper Upgrade','All units gain +5HP, +10% attack, and +1 movement.','No day-to-day abilities.'),(2,'Grit','grit.png',2,270000,540000,'Snipe Attack','Indirect units gain +1 range and their attack is increased to +40%.','Super Snipe','Indirect units gain +2 range and their attack is increased to +40%.','Indirect units have +1 range and gain +20% attack. Direct units lose -20% attack (footsoldiers are normal).'),(3,'Kanbei','kanbei.png',4,360000,630000,'Morale Boost','All units\' attack is increased to +40%.','Samurai Spirit','All units\' attack is increased to +40% and defense to +50%. Counterattacks do 1.5x more damage.','Units cost +20% more to build, but gain +30% attack and defense.'),(5,'Drake','drake.png',3,360000,630000,'Tsunami','All enemy units lose 1 HP (to a minimum of 0.1HP) and half their fuel.','Typhoon','All enemy units lose 2 HP (to a minimum of 0.1HP) and half their fuel. Weather changes to Rain for 1 day.','Naval units gain +1 movement and +25 defense. Air units lose -20% attack. Unaffected by rain (except vision), and has a higher chance of Rain in random weather.'),(7,'Max','max.png',1,270000,540000,'Max Force','Direct units gain +1 movement and their attack increases to +30%.','Max Blast','Direct units gain +2 movement and their attack increases to +50%.','Direct units gain +20% attack. Indirect units lose -10% attack and have -1 range.'),(8,'Sami','sami.png',1,270000,720000,'Double Time','Footsoldiers gain +1 movement and their attack is increased to +50%.','Victory March','Footsoldiers gain +2 movement and their attack is increased to +70%. Footsoldiers can capture buildings instantly (even with 1 HP).','Footsoldiers gain +30% attack and a 50% capture point bonus (rounded down). Other direct units lose -10% attack. Transports gain +1 movement.'),(9,'Olaf','olaf.png',2,270000,630000,'Blizzard','Changes the weather to Snow for 1 day.','Winter Fury','Enemy units lose 2HP (to a minimum of 0.1HP), and the weather changes to Snow for 1 day.','Unaffected by snow, but rain affects him the same as snow would for others.'),(10,'Eagle','eagle.png',3,270000,810000,'Lightning Drive','Air units attack and defense are increased to +20%.','Lightning Strike','Air units attack and defense are increased to +20%. All non-footsoldier units may move and fire again, even if built this turn (use this power after moving!).','Air units gain +15% attack and +10% defense, and consume -2 fuel per day. Naval units lose -30% attack.'),(11,'Adder','adder.png',5,180000,450000,'Sideslip','All units gain +1 movement.','Sidewinder','All units gain +2 movement.','No day-to-day abilities.'),(12,'Hawke','hawke.png',5,450000,810000,'Black Wave','All units gain +1HP, and all enemy units lose -1HP (to a minimum of 0.1HP).','Black Storm','All units gain +2HP, and all enemy units lose -2HP (to a minimum of 0.1HP).','Units gain +10% attack.'),(13,'Sensei','sensei.png',4,180000,540000,'Copter Command','Copters\' attack is increased to +65%. 9HP unwaited infantry are placed on every owned, empty city.','Airborne Assault','Copters\' attack is increased to +65%. 9HP unwaited mechs are placed on every owned, empty city.','Copters gain +50% attack, footsoldiers gain +40% attack, but all other non-air units lose -10% attack. Transports gain +1 movement.'),(14,'Jess','jess.png',3,270000,540000,'Turbo Charge','Vehicles gain +1 movement and their attack is increased to +20%. All units resupply fuel and ammo.','Overdrive','Vehicles gain +2 movement and their attack is increased to +40%. All units resupply fuel and ammo.','Vehicles gain +10% attack, but all other units (including footsoldiers) lose -10% attack.'),(15,'Colin','colin.png',2,180000,540000,'Gold Rush','Funds are multiplied by 1.5x.','Power of Money','Unit attack percentage increases by (3 * Funds / 1000)%.','Units cost -20% less to build and lose -10% attack.'),(16,'Lash','lash.png',5,360000,630000,'Terrain Tactics','Movement cost for all terrain is reduced to 1, except in Snow.','Prime Tactics','Terrain stars are doubled (attack and defense). Movement cost over all terrain is reduced to 1, except in Snow.','Units gain +10% attack for every terrain star (note: air units are unaffected by terrain).'),(17,'Hachi','hachi.png',1,270000,450000,'Barter','Units cost -50% to build.','Merchant Union','Units cost -50% and ground units can deploy from cities.','Units cost -10% less to build.'),(18,'Sonja','sonja.png',4,270000,450000,'Enhanced Vision','All units gain +1 vision, and can see into forests and reefs.','Counter Break','All units gain +1 vision, and can see into forests and reefs. A unit being attacked will attack first (even if it would be destroyed by the attack).','Units gain +1 vision in Fog of War, and counterattacks do 1.5x more damage. All units have hidden HP, but have bad luck (-9% to +9% luck).'),(19,'Sasha','sasha.png',2,180000,540000,'Market Crash','Reduces enemy power bar(s) by 10% per 5000 funds.','War Bonds','Receives funds equal to 50% of the damage dealt when attacking enemy units.','Receives +100 income per property. Note that labs and towers do not count towards this.'),(20,'Grimm','grimm.png',4,270000,540000,'Knucklebuster','All units\' attack is increased to +50%.','Haymaker','All units\' attack is increased to +80%.','Units gain +30% attack, but lose -20% defense.'),(21,'Koal','koal.png',5,270000,450000,'Forced March','All units gain +1 movement, and the road bonus is increased to +20%.','Trail of Woe','All units gain +2 movement, and the road bonus is increased to +30%.','Units (even air units) gain +10% attack power on roads.'),(22,'Jake','jake.png',1,270000,540000,'Beat Down','Land indirects gain +1 range, and plains bonus is increased to +20%.','Block Rock','Land indirects gain +1 range, plains bonus is increased to +40%, and vehicles gain +2 movement.','Units (even air units) gain +10% attack power on plains.'),(23,'Kindle','kindle.png',5,270000,540000,'Urban Blight','All enemy units lose -3HP on urban terrain. Urban bonus is increased to +80%.','High Society','Urban bonus is increased to +130%, and attack for all units is increased by +3% for every city owned.','All units (even air units) gain +40% attack while on urban terrain (HQs, bases, (air)ports, cities, labs, and comtowers).'),(24,'Nell','nell.png',1,270000,540000,'Lucky Star',' Luck is improved to +0% to +59%.','Lady Luck','Luck is improved to +0% to +99%.','Luck on attacks is +0 to +19%.'),(25,'Flak','flak.png',5,270000,540000,'Brute Force','Luck range is changed to -19% to +49%.','Barbaric Blow','Luck range is changed to -39% to +89%.','Luck on attacks is -9% to +24%.'),(26,'Jugger','jugger.png',5,270000,630000,'Overclock','Luck range is changed to -24% to +54%.','System Crash','Luck range is changed to -44% to +94%.','Luck on attacks is -14% to +29%.'),(27,'Javier','javier.png',3,270000,540000,'Tower Shield','Indirect defense is increased to +40%. Comtower bonuses are doubled.','Tower of Power','Indirect defense is increased to +60%. Comtower bonuses are tripled.','Units gain +20% defense against indirect units. Comtowers grant all units additional +10% defense.'),(28,'Rachel','rachel.png',1,270000,540000,'Lucky Lass','Luck is improved to +0% to +39%.','Covering Fire','Three 2-range missiles deal 3HP damage each. The missiles target the opponents\' greatest accumulation of footsoldier HP, unit value, and unit HP (in that order).','Units repair +1 additional HP (note: liable for costs).'),(29,'Sturm','sturm.png',5,540000,900000,'Meteor Strike','A 2-range missile deals 4HP damage. The missile targets an enemy unit located at the greatest accumulation of unit value.','Meteor Strike II','A 2-range missile deals 8HP damage. The missile targets an enemy unit located at the greatest accumulation of unit value.','Movement cost over all terrain is reduced to 1, except in Snow. Units lose -20% attack and gain +20% defense.'),(30,'Von Bolt','vonbolt.png',5,900000,900000,'None','','Ex Machina','A 2-range missile deals 3HP damage and prevents all affected enemy units from acting next turn. The missile targets the opponents\' greatest accumulation of unit value.','Units gain +10% attack and +10% defense.'),(31,'Blank',NULL,NULL,0,0,'','','','','');
/*!40000 ALTER TABLE `awbw_co` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `awbw_terrain`
--

DROP TABLE IF EXISTS `awbw_terrain`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_terrain` (
  `terrain_id` int(11) NOT NULL AUTO_INCREMENT,
  `terrain_name` varchar(30) DEFAULT NULL,
  `terrain_defense` int(11) DEFAULT NULL,
  `terrain_symbol` char(1) DEFAULT NULL,
  `terrain_country_code` char(2) DEFAULT NULL,
  `terrain_build_type` char(1) DEFAULT NULL,
  `terrain_active` char(1) NOT NULL DEFAULT 'N',
  `terrain_offset` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`terrain_id`)
) ENGINE=InnoDB AUTO_INCREMENT=195 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `awbw_terrain`
--

LOCK TABLES `awbw_terrain` WRITE;
/*!40000 ALTER TABLE `awbw_terrain` DISABLE KEYS */;
INSERT INTO `awbw_terrain` VALUES (1,'Plain',1,'.','',NULL,'Y',0),(2,'Mountain',4,'^','','','Y',5),(3,'Wood',2,'@','',NULL,'Y',0),(4,'HRiver',0,'{','',NULL,'Y',0),(5,'VRiver',0,'}','',NULL,'Y',0),(6,'CRiver',0,'~','',NULL,'Y',0),(7,'ESRiver',0,'I','',NULL,'Y',0),(8,'SWRiver',0,'J','',NULL,'Y',0),(9,'WNRiver',0,'K','','','Y',0),(10,'NERiver',0,'L','',NULL,'Y',0),(11,'ESWRiver',0,'M','',NULL,'Y',0),(12,'SWNRiver',0,'N','',NULL,'Y',0),(13,'WNERiver',0,'O','',NULL,'Y',0),(14,'NESRiver',0,'P','',NULL,'Y',0),(15,'HRoad',0,'-','',NULL,'Y',0),(16,'VRoad',0,'=','',NULL,'Y',0),(17,'CRoad',0,'+','',NULL,'Y',0),(18,'ESRoad',0,'A','',NULL,'Y',0),(19,'SWRoad',0,'B','',NULL,'Y',0),(20,'WNRoad',0,'C','',NULL,'Y',0),(21,'NERoad',0,'D','',NULL,'Y',0),(22,'ESWRoad',0,'E','',NULL,'Y',0),(23,'SWNRoad',0,'F','',NULL,'Y',0),(24,'WNERoad',0,'G','',NULL,'Y',0),(25,'NESRoad',0,'H','',NULL,'Y',0),(26,'HBridge',0,'[','',NULL,'Y',0),(27,'VBridge',0,']','',NULL,'Y',0),(28,'Sea',0,',','',NULL,'Y',0),(29,'HShoal',0,'<','',NULL,'Y',0),(30,'HShoalN',0,'(','',NULL,'Y',0),(31,'VShoal',0,'>','',NULL,'Y',0),(32,'VShoalE',0,')','',NULL,'Y',0),(33,'Reef',1,'%','',NULL,'Y',0),(34,'Neutral City',3,'a','',NULL,'Y',6),(35,'Neutral Base',3,'b','','','Y',9),(36,'Neutral Airport',3,'c','',NULL,'Y',2),(37,'Neutral Port',3,'d','',NULL,'Y',6),(38,'Orange Star City',3,'e','os',NULL,'Y',6),(39,'Orange Star Base',3,'f','os','L','Y',9),(40,'Orange Star Airport',3,'g','os','A','Y',2),(41,'Orange Star Port',3,'h','os','S','Y',6),(42,'Orange Star HQ',4,'i','os',NULL,'Y',15),(43,'Blue Moon City',3,'j','bm',NULL,'Y',6),(44,'Blue Moon Base',3,'l','bm','L','Y',9),(45,'Blue Moon Airport',3,'m','bm','A','Y',2),(46,'Blue Moon Port',3,'n','bm','S','Y',6),(47,'Blue Moon HQ',4,'o','bm',NULL,'Y',15),(48,'Green Earth City',3,'p','ge',NULL,'Y',6),(49,'Green Earth Base',3,'q','ge','L','Y',9),(50,'Green Earth Airport',3,'r','ge','A','Y',2),(51,'Green Earth Port',3,'s','ge','S','Y',6),(52,'Green Earth HQ',4,'t','ge',NULL,'Y',15),(53,'Yellow Comet City',3,'u','yc',NULL,'Y',6),(54,'Yellow Comet Base',3,'v','yc','L','Y',9),(55,'Yellow Comet Airport',3,'w','yc','A','Y',2),(56,'Yellow Comet Port',3,'x','yc','S','Y',6),(57,'Yellow Comet HQ',4,'y','yc',NULL,'Y',15),(81,'Red Fire City',3,'U','rf',NULL,'Y',6),(82,'Red Fire Base',3,'T','rf','L','Y',9),(83,'Red Fire Airport',3,'S','rf','A','Y',2),(84,'Red Fire Port',3,'R','rf','S','Y',6),(85,'Red Fire HQ',4,'Q','rf',NULL,'Y',15),(86,'Grey Sky City',3,'Z','gs',NULL,'Y',6),(87,'Grey Sky Base',3,'Y','gs','L','Y',9),(88,'Grey Sky Airport',3,'X','gs','A','Y',2),(89,'Grey Sky Port',3,'W','gs','S','Y',6),(90,'Grey Sky HQ',4,'V','gs',NULL,'Y',15),(91,'Black Hole City',3,'1','bh',NULL,'Y',6),(92,'Black Hole Base',3,'2','bh','L','Y',9),(93,'Black Hole Airport',3,'3','bh','A','Y',2),(94,'Black Hole Port',3,'4','bh','S','Y',6),(95,'Black Hole HQ',4,'5','bh','','Y',16),(96,'Brown Desert City',3,'','bd',NULL,'Y',6),(97,'Brown Desert Base',3,'','bd','L','Y',9),(98,'Brown Desert Airport',3,'','bd','A','Y',2),(99,'Brown Desert Port',3,'','bd','S','Y',6),(100,'Brown Desert HQ',4,'','bd',NULL,'Y',15),(101,'VPipe',0,'k','','','Y',0),(102,'HPipe',0,'z','','','Y',0),(103,'NEPipe',0,'!','','','Y',0),(104,'ESPipe',0,'#','','','Y',0),(105,'SWPipe',0,'$','','','Y',0),(106,'WNPipe',0,'&','','','Y',0),(107,'NPipe End',0,'*','','','Y',0),(108,'EPipe End',0,'|','','','Y',0),(109,'SPipe End',0,'`','','','Y',0),(110,'WPipe End',0,'\'','','','Y',0),(111,'Missile Silo',3,'\"','','','Y',8),(112,'Missile Silo Empty',3,';','','','Y',0),(113,'HPipe Seam',0,':','','','Y',0),(114,'VPipe Seam',0,'?','','','Y',0),(115,'HPipe Rubble',1,'/','','','Y',0),(116,'VPipe Rubble',1,'0','','','Y',0),(117,'Amber Blaze Airport',3,'','ab','A','Y',2),(118,'Amber Blaze Base',3,'','ab','L','Y',9),(119,'Amber Blaze City',3,'','ab','','Y',6),(120,'Amber Blaze HQ',4,'','ab','','Y',15),(121,'Amber Blaze Port',3,'','ab','S','Y',6),(122,'Jade Sun Airport',3,'','js','A','Y',2),(123,'Jade Sun Base',3,'','js','L','Y',9),(124,'Jade Sun City',3,'','js','','Y',6),(125,'Jade Sun HQ',4,'','js','','Y',15),(126,'Jade Sun Port',3,'','js','S','Y',6),(127,'Amber Blaze Com Tower',3,'','ab','','Y',9),(128,'Black Hole Com Tower',3,'','bh','','Y',9),(129,'Blue Moon Com Tower',3,'','bm','','Y',9),(130,'Brown Desert Com Tower',3,'','bd','','Y',9),(131,'Green Earth Com Tower',3,'','ge','','Y',9),(132,'Jade Sun Com Tower',3,'','js','','Y',9),(133,'Neutral Com Tower',3,'_','','','Y',9),(134,'Orange Star Com Tower',3,'','os','','Y',9),(135,'Red Fire Com Tower',3,'','rf','','Y',9),(136,'Yellow Comet Com Tower',3,'','yc','','Y',9),(137,'Grey Sky Com Tower',3,'','gs','','Y',9),(138,'Amber Blaze Lab',3,'','ab','','Y',8),(139,'Black Hole Lab',3,'','bh','','Y',8),(140,'Blue Moon Lab',3,'','bm','','Y',8),(141,'Brown Desert Lab',3,'','bd','','Y',8),(142,'Green Earth Lab',3,'','ge','','Y',8),(143,'Grey Sky Lab',3,'','gs','','Y',8),(144,'Jade Sun Lab',3,'','js','','Y',8),(145,'Neutral Lab',3,'6','','','Y',8),(146,'Orange Star Lab',3,'','os','','Y',8),(147,'Red Fire Lab',3,'','rf','','Y',8),(148,'Yellow Comet Lab',3,'','yc','','Y',8),(149,'Cobalt Ice Airport',3,'','ci','A','Y',2),(150,'Cobalt Ice Base',3,'','ci','L','Y',9),(151,'Cobalt Ice City',3,'','ci','','Y',6),(152,'Cobalt Ice Com Tower',3,'','ci','','Y',9),(153,'Cobalt Ice HQ',4,'','ci','','Y',13),(154,'Cobalt Ice Lab',3,'','ci','','Y',8),(155,'Cobalt Ice Port',3,'','ci','S','Y',6),(156,'Pink Cosmos Airport',3,'','pc','A','Y',2),(157,'Pink Cosmos Base',3,'','pc','L','Y',9),(158,'Pink Cosmos City',3,'','pc','','Y',6),(159,'Pink Cosmos Com Tower',3,'','pc','','Y',9),(160,'Pink Cosmos HQ',4,'','pc','','Y',15),(161,'Pink Cosmos Lab',3,'','pc','','Y',8),(162,'Pink Cosmos Port',3,'','pc','S','Y',6),(163,'Teal Galaxy Airport',3,'','tg','A','Y',2),(164,'Teal Galaxy Base',3,'','tg','L','Y',9),(165,'Teal Galaxy City',3,'','tg','','Y',6),(166,'Teal Galaxy Com Tower',3,'','tg','','Y',9),(167,'Teal Galaxy HQ',4,'','tg','','Y',16),(168,'Teal Galaxy Lab',3,'','tg','','Y',8),(169,'Teal Galaxy Port',3,'','tg','S','Y',6),(170,'Purple Lightning Airport',3,'','pl','A','Y',2),(171,'Purple Lightning Base',3,'','pl','L','Y',9),(172,'Purple Lightning City',3,'','pl','','Y',6),(173,'Purple Lightning Com Tower',3,'','pl','','Y',9),(174,'Purple Lightning HQ',4,'','pl','','Y',16),(175,'Purple Lightning Lab',3,'','pl','','Y',8),(176,'Purple Lightning Port',3,'','pl','S','Y',6),(181,'Acid Rain Airport',3,'','ar','A','Y',2),(182,'Acid Rain Base',3,'','ar','L','Y',9),(183,'Acid Rain City',3,'','ar','','Y',6),(184,'Acid Rain Com Tower',3,'','ar','','Y',9),(185,'Acid Rain HQ',4,'','ar','','Y',15),(186,'Acid Rain Lab',3,'','ar','','Y',8),(187,'Acid Rain Port',3,'','ar','S','Y',6),(188,'White Nova Airport',3,'','wn','A','Y',2),(189,'White Nova Base',3,'','wn','L','Y',9),(190,'White Nova City',3,'','wn','','Y',6),(191,'White Nova Com Tower',3,'','wn','','Y',9),(192,'White Nova HQ',4,'','wn','','Y',15),(193,'White Nova Lab',3,'','wn','','Y',8),(194,'White Nova Port',3,'','wn','S','Y',6);
/*!40000 ALTER TABLE `awbw_terrain` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-10-17 11:43:12

-- MySQL dump 10.13  Distrib 5.5.47, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: awbw
-- ------------------------------------------------------
-- Server version	5.5.47-0+deb8u1-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `awbw_units`
--

DROP TABLE IF EXISTS `awbw_units`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awbw_units` (
  `units_id` int(11) NOT NULL AUTO_INCREMENT,
  `units_games_id` int(11) DEFAULT NULL,
  `units_players_id` int(11) DEFAULT NULL,
  `units_name` varchar(30) DEFAULT NULL,
  `units_movement_points` int(11) DEFAULT NULL,
  `units_vision` int(11) DEFAULT NULL,
  `units_fuel` int(11) DEFAULT NULL,
  `units_fuel_per_turn` int(11) DEFAULT NULL,
  `units_sub_dive` char(1) DEFAULT NULL,
  `units_ammo` int(11) DEFAULT NULL,
  `units_short_range` int(11) DEFAULT NULL,
  `units_long_range` int(11) DEFAULT NULL,
  `units_second_weapon` char(1) DEFAULT NULL,
  `units_symbol` char(1) DEFAULT NULL,
  `units_cost` int(11) DEFAULT NULL,
  `units_movement_type` char(1) DEFAULT NULL,
  `units_x` int(11) DEFAULT NULL,
  `units_y` int(11) DEFAULT NULL,
  `units_moved` int(11) DEFAULT NULL,
  `units_capture` int(11) DEFAULT NULL,
  `units_fired` int(11) DEFAULT NULL,
  `units_hit_points` double(3,1) DEFAULT NULL,
  `units_cargo1_units_id` int(11) NOT NULL DEFAULT '0',
  `units_cargo2_units_id` int(11) NOT NULL DEFAULT '0',
  `units_carried` char(1) NOT NULL DEFAULT '',
  PRIMARY KEY (`units_id`),
  KEY `units_players_id_games_id_x_y` (`units_players_id`,`units_games_id`,`units_x`,`units_y`),
  KEY `units_x_y` (`units_x`,`units_y`),
  KEY `units_games_id_players_id` (`units_games_id`,`units_players_id`),
  KEY `units_games_id` (`units_games_id`,`units_x`,`units_y`),
  KEY `units_cargo2` (`units_cargo2_units_id`),
  KEY `units_cargo1` (`units_cargo1_units_id`)
) ENGINE=MyISAM AUTO_INCREMENT=115755395 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `awbw_units`
--
-- WHERE:  units_games_id=-1

LOCK TABLES `awbw_units` WRITE;
/*!40000 ALTER TABLE `awbw_units` DISABLE KEYS */;
INSERT INTO `awbw_units` VALUES (1,-1,-1,'Infantry',3,2,99,0,'N',0,0,0,'N','A',1000,'F',NULL,NULL,NULL,0,0,NULL,0,0,'N'),(2,-1,-1,'Mech',2,2,70,0,'N',3,0,0,'Y','B',3000,'B',NULL,NULL,NULL,0,0,NULL,0,0,'N'),(3,-1,-1,'Md.Tank',5,1,50,0,'N',8,0,0,'Y','C',16000,'T',NULL,NULL,NULL,0,0,NULL,0,0,'N'),(4,-1,-1,'Tank',6,3,70,0,'N',9,0,0,'Y','D',7000,'T',NULL,NULL,NULL,0,0,NULL,0,0,'N'),(5,-1,-1,'Recon',8,5,80,0,'N',0,0,0,'N','E',4000,'W',NULL,NULL,NULL,0,0,NULL,0,0,'N'),(6,-1,-1,'APC',6,1,70,0,'N',0,0,0,'N','F',5000,'T',NULL,NULL,NULL,0,0,NULL,0,0,'N'),(7,-1,-1,'Artillery',5,1,50,0,'N',9,2,3,'N','G',6000,'T',NULL,NULL,NULL,0,0,NULL,0,0,'N'),(8,-1,-1,'Rocket',5,1,50,0,'N',6,3,5,'N','H',15000,'W',NULL,NULL,NULL,0,0,NULL,0,0,'N'),(9,-1,-1,'Anti-Air',6,2,60,0,'N',9,0,0,'N','I',8000,'T',NULL,NULL,NULL,0,0,NULL,0,0,'N'),(10,-1,-1,'Missile',4,5,50,0,'N',6,3,5,'N','J',12000,'W',NULL,NULL,NULL,0,0,NULL,0,0,'N'),(11,-1,-1,'Fighter',9,2,99,5,'N',9,0,0,'N','L',20000,'A',NULL,NULL,NULL,0,0,NULL,0,0,'N'),(12,-1,-1,'Bomber',7,2,99,5,'N',9,0,0,'N','M',22000,'A',NULL,NULL,NULL,0,0,NULL,0,0,'N'),(13,-1,-1,'B-Copter',6,3,99,2,'N',6,0,0,'Y','N',9000,'A',NULL,NULL,NULL,0,0,NULL,0,0,'N'),(14,-1,-1,'T-Copter',6,2,99,2,'N',0,0,0,'N','O',5000,'A',NULL,NULL,NULL,0,0,NULL,0,0,'N'),(15,-1,-1,'Battleship',5,2,99,1,'N',9,2,6,'N','P',28000,'S',NULL,NULL,NULL,0,0,NULL,0,0,'N'),(16,-1,-1,'Cruiser',6,3,99,1,'N',9,0,0,'N','Q',18000,'S',NULL,NULL,NULL,0,0,NULL,0,0,'N'),(17,-1,-1,'Lander',6,1,99,1,'N',0,0,0,'N','R',12000,'L',NULL,NULL,NULL,0,0,NULL,0,0,'N'),(18,-1,-1,'Sub',5,5,60,1,'N',6,0,0,'N','S',20000,'S',NULL,NULL,NULL,0,0,NULL,0,0,'N'),(46,-1,-1,'Neotank',6,1,99,1,'N',9,0,0,'Y','T',22000,'T',NULL,NULL,NULL,0,0,10.0,0,0,'N'),(960900,-1,NULL,'Piperunner',9,4,99,0,NULL,9,2,5,'Y','U',20000,'P',NULL,NULL,NULL,NULL,NULL,NULL,0,0,''),(968731,-1,NULL,'Black Bomb',9,1,45,5,NULL,0,0,0,'N','V',25000,'A',NULL,NULL,NULL,NULL,NULL,NULL,0,0,''),(1141438,-1,-1,'Mega Tank',4,1,50,0,'N',3,0,0,'Y','Z',28000,'T',NULL,NULL,NULL,NULL,NULL,10.0,0,0,'N'),(28,-1,NULL,'Black Boat',7,1,60,1,'N',0,0,0,'N','A',7500,'L',0,2,1,1,1,10.0,0,0,'N'),(30,-1,0,'Stealth',6,4,60,5,'N',6,0,0,'N','D',24000,'A',4,6,0,0,0,10.0,0,0,'N'),(29,-1,0,'Carrier',5,4,99,1,'N',9,3,8,'N','B',30000,'S',6,8,0,0,0,10.0,0,0,'N');
/*!40000 ALTER TABLE `awbw_units` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-10-17 11:49:20
