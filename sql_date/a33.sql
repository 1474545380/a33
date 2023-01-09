/*
 Navicat Premium Data Transfer

 Source Server         : mysql_1
 Source Server Type    : MySQL
 Source Server Version : 80025
 Source Host           : localhost:3306
 Source Schema         : a33

 Target Server Type    : MySQL
 Target Server Version : 80025
 File Encoding         : 65001

 Date: 10/01/2023 00:05:51
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for schedul_rules
-- ----------------------------
DROP TABLE IF EXISTS `schedul_rules`;
CREATE TABLE `schedul_rules` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `rule_type` varchar(255) NOT NULL COMMENT '规则类型；可选值:开店规则、关店规则、客流规则。可扩展',
  `store_identity` longtext COMMENT '为空时，为系统通用规则。\n不为空时，为门店规则。\n当门店有门店规则时，使用门店规则进行排班，没\n有门店规则时，使用系统通用规则进行排班',
  `rule_date` text NOT NULL COMMENT '需要开发者自行设计。\n示例:\n- 客流规则:"3.8" 表示按照业务预测数据，每 3.8 个客流必须安排至少一个员工当值\n- 开店规则:"1.5,23.5" 表示开店 1 个半小时前需要 有员工当值，当值员工数为门店面积除以 23.5\n- 关店规则:"2.5,3,13" 表示关店 2 个半小时内需要 有员工当值，当值员工数不小于 3 并且不小于门店 面积除以 13\n为了提高规则的灵活性，建议使用 json 格式保存规 则值，如关店规则: {"after":"2.5","count":"3","fomula":"size/13"}',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for staff
-- ----------------------------
DROP TABLE IF EXISTS `staff`;
CREATE TABLE `staff` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `identity` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '员工表的唯一标识',
  `password` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '员工登录密码',
  `name` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '员工真实姓名',
  `email` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '员工邮箱',
  `position` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '员工职位',
  `store_identity` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '商店唯一标识',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_staff_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for staff_preference
-- ----------------------------
DROP TABLE IF EXISTS `staff_preference`;
CREATE TABLE `staff_preference` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `staff_preference_identity` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '员工偏好唯一标识',
  `preference_type` varchar(255) NOT NULL COMMENT '偏好类型',
  `staff_identity` longtext NOT NULL COMMENT '员工唯一标识',
  `preference_value` text NOT NULL COMMENT '偏好值',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for store
-- ----------------------------
DROP TABLE IF EXISTS `store`;
CREATE TABLE `store` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `identity` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '门店表的唯一标识',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '门店名称',
  `address` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '门店地址',
  `size` float unsigned NOT NULL COMMENT '门店大小',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;
