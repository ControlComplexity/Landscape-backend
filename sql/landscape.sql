/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50727
Source Host           : localhost:3306
Source Database       : landscape

Target Server Type    : MYSQL
Target Server Version : 50727
File Encoding         : 65001

Date: 2022-12-29 11:08:56
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for essay
-- ----------------------------
DROP TABLE IF EXISTS `essay`;
CREATE TABLE `essay` (
                         `id` int(11) NOT NULL,
                         `uuid` varchar(255) DEFAULT NULL,
                         `title` varchar(255) DEFAULT NULL,
                         `content` varchar(255) DEFAULT NULL,
                         `url` varchar(255) DEFAULT NULL,
                         `time` timestamp(6) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(6),
                         `type` varchar(255) DEFAULT NULL,
                         `city` varchar(255) DEFAULT NULL,
                         `hits` bigint(233) DEFAULT NULL,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `swiper_image`;
CREATE TABLE `swiper_image` (
                         `id` int(11) NOT NULL,
                         `uuid` varchar(255) DEFAULT NULL,
                         `title` varchar(255) DEFAULT NULL,
                         `url` varchar(255) DEFAULT NULL,
                         `time` timestamp(6) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(6),
                         `type` varchar(255) DEFAULT NULL,
                         `city` varchar(255) DEFAULT NULL,
                         `hits` bigint(233) DEFAULT NULL,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
