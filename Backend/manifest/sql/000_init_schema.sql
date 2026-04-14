-- =============================================================================
-- 三角机构 TRPG 完整建表脚本（可直接执行）
-- 执行前请先创建数据库：CREATE DATABASE trangleagent CHARACTER SET utf8mb4;
-- =============================================================================

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- -----------------------------------------------------------------------------
-- 1. 用户表
-- -----------------------------------------------------------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `account` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '账号',
  `password` VARCHAR(128) NOT NULL DEFAULT '' COMMENT '密码哈希',
  `nickname` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '昵称',
  `gender` TINYINT NOT NULL DEFAULT 0 COMMENT '性别：0未知 1男 2女',
  `birth_date` DATE NULL DEFAULT NULL COMMENT '生日',
  `user_type` VARCHAR(32) NOT NULL DEFAULT 'user' COMMENT '用户类型：user/admin',
    `created_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `extra_info` JSON NULL COMMENT '扩展配置信息',
    `email` VARCHAR(128) NOT NULL DEFAULT '' COMMENT '邮箱',
  `vip_start_at` DATETIME NULL DEFAULT NULL COMMENT 'VIP开始时间',
  `vip_end_at` DATETIME NULL DEFAULT NULL COMMENT 'VIP结束时间',
  `active_role_id` BIGINT UNSIGNED NULL DEFAULT NULL COMMENT '当前选中的角色卡ID',
  `exp` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '经验值（经验条）',
  `level` TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '等级（1+exp/100）',
  `last_checkin_at` DATETIME NULL DEFAULT NULL COMMENT '上次签到时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_account` (`account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- -----------------------------------------------------------------------------
-- 2. 部门表
-- -----------------------------------------------------------------------------
DROP TABLE IF EXISTS `department`;
CREATE TABLE `department` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属用户ID',
  `branch_name` VARCHAR(128) NOT NULL DEFAULT '' COMMENT '分部名称',
  `terminal_count` INT NOT NULL DEFAULT 0 COMMENT '分部散逸端数量',
  `weather` VARCHAR(128) NOT NULL DEFAULT '' COMMENT '分部天气/气候',
  `manager_name` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '分部经理名称',
  `location` VARCHAR(256) NOT NULL DEFAULT '' COMMENT '地址',
  `created_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='部门/分部表';

-- -----------------------------------------------------------------------------
-- 3. 角色卡表（含轨道、QA）
-- -----------------------------------------------------------------------------
DROP TABLE IF EXISTS `role_cards`;
CREATE TABLE `role_cards` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '角色卡ID',
  `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属用户ID',
  `department_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属部门ID',
  `commendation` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '嘉奖次数',
  `reprimand` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '申戒次数',
  `blue_track` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '蓝轨（0-40）',
  `yellow_track` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '黄轨（0-40）',
  `red_track` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '红轨（0-40）',
  `arc_abnormal` VARCHAR(128) NOT NULL DEFAULT '' COMMENT 'ARC：异常',
  `arc_reality` VARCHAR(128) NOT NULL DEFAULT '' COMMENT 'ARC：现实',
  `arc_position` VARCHAR(128) NOT NULL DEFAULT '' COMMENT 'ARC：职位',
  `agent_name` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '特工名字',
  `code_name` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '代号',
  `gender` VARCHAR(16) NOT NULL DEFAULT '' COMMENT '性别',
  `qa_meticulous` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'QA Meticulousness (0-100)',
  `qa_deception` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'QA Deception (0-100)',
  `qa_vigor` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'QA Vigor (0-100)',
  `qa_empathy` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'QA Empathy (0-100)',
  `qa_initiative` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'QA Initiative (0-100)',
  `qa_resilience` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'QA Resilience (0-100)',
  `qa_presence` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'QA Presence (0-100)',
  `qa_professional` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'QA Professionalism (0-100)',
  `qa_discretion` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'QA Discretion (0-100)',
  `extra_info` JSON NULL COMMENT '扩展配置信息',
  `created_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_department_id` (`department_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色卡表（含轨道、QA）';

-- -----------------------------------------------------------------------------
-- 4. TRPG 房间表
-- -----------------------------------------------------------------------------
DROP TABLE IF EXISTS `trpg_room`;
CREATE TABLE `trpg_room` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `room_id` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '房间全局唯一ID',
  `room_code` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '房间号',
  `room_name` VARCHAR(128) NOT NULL DEFAULT '' COMMENT '房间名称',
  `host_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '主持人用户ID',
  `host_nickname` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '主持人昵称',
  `max_players` TINYINT UNSIGNED NOT NULL DEFAULT 4 COMMENT '最大玩家人数',
  `current_players` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '当前玩家人数',
  `has_password` TINYINT NOT NULL DEFAULT 0 COMMENT '是否有密码：0无 1有',
  `room_password` VARCHAR(128) NOT NULL DEFAULT '' COMMENT '房间密码',
  `is_private` TINYINT NOT NULL DEFAULT 0 COMMENT '是否私密：0公开 1私密',
  `status` TINYINT NOT NULL DEFAULT 0 COMMENT '状态：0未开始 1进行中 2已结束 3已关闭',
  `system_name` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '规则系统',
  `scenario_name` VARCHAR(128) NOT NULL DEFAULT '' COMMENT '模组/剧本名称',
  `description` TEXT COMMENT '房间简介',
  `created_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `started_at` DATETIME NULL DEFAULT NULL COMMENT '开团时间',
  `ended_at` DATETIME NULL DEFAULT NULL COMMENT '结束时间',
  `updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_host_id` (`host_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='TRPG房间表';

-- -----------------------------------------------------------------------------
-- 5. 房间成员表
-- -----------------------------------------------------------------------------
DROP TABLE IF EXISTS `trpg_room_member`;
CREATE TABLE `trpg_room_member` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `room_id` BIGINT UNSIGNED NOT NULL COMMENT '房间ID',
  `user_id` BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
  `role_card_id` BIGINT UNSIGNED NOT NULL COMMENT '使用的角色卡ID',
  `joined_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '加入时间',
  `left_at` DATETIME NULL DEFAULT NULL COMMENT '离开时间',
  `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：0已离开 1在线',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_room_user` (`room_id`, `user_id`),
  KEY `idx_room_id` (`room_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_role_card_id` (`role_card_id`),
  CONSTRAINT `fk_room_member_room` FOREIGN KEY (`room_id`) REFERENCES `trpg_room` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_room_member_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_room_member_role` FOREIGN KEY (`role_card_id`) REFERENCES `role_cards` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='TRPG房间成员';

-- -----------------------------------------------------------------------------
-- 6. 论坛频道
-- -----------------------------------------------------------------------------
DROP TABLE IF EXISTS `forum_sections`;
CREATE TABLE `forum_sections` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '频道ID',
  `name` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '频道名称',
  `description` VARCHAR(256) NOT NULL DEFAULT '' COMMENT '频道简介',
  `icon` VARCHAR(256) NOT NULL DEFAULT '' COMMENT '频道图标URL',
  `status` VARCHAR(32) NOT NULL DEFAULT 'normal' COMMENT '状态：normal/hidden/deleted',
  `display_order` INT NOT NULL DEFAULT 0 COMMENT '排序值',
  `created_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='论坛频道';

-- -----------------------------------------------------------------------------
-- 7. 论坛版块
-- -----------------------------------------------------------------------------
DROP TABLE IF EXISTS `forum_boards`;
CREATE TABLE `forum_boards` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '版块ID',
  `section_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属频道ID',
  `name` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '版块名称',
  `description` VARCHAR(256) NOT NULL DEFAULT '' COMMENT '版块简介',
  `cover_image` VARCHAR(256) NOT NULL DEFAULT '' COMMENT '版块封面图URL',
  `status` VARCHAR(32) NOT NULL DEFAULT 'normal' COMMENT '状态：normal/hidden/deleted',
  `display_order` INT NOT NULL DEFAULT 0 COMMENT '排序值',
  `post_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '帖子总数',
  `today_post_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '今日发帖数',
  `last_post_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后一篇帖子ID',
  `last_post_at` DATETIME NULL DEFAULT NULL COMMENT '最后发帖时间',
  `created_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_section_id` (`section_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='论坛版块';

-- -----------------------------------------------------------------------------
-- 8. 论坛帖子
-- -----------------------------------------------------------------------------
DROP TABLE IF EXISTS `forum_posts`;
CREATE TABLE `forum_posts` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '帖子ID',
  `board_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属版块ID',
  `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '发帖用户ID',
  `title` VARCHAR(256) NOT NULL DEFAULT '' COMMENT '帖子标题',
  `content` TEXT COMMENT '帖子正文',
  `cover_image` VARCHAR(256) NOT NULL DEFAULT '' COMMENT '封面图URL',
  `status` VARCHAR(32) NOT NULL DEFAULT 'normal' COMMENT '状态：normal/deleted/audit/reject',
  `is_pinned` TINYINT NOT NULL DEFAULT 0 COMMENT '是否置顶：0否 1是',
  `is_essence` TINYINT NOT NULL DEFAULT 0 COMMENT '是否精华：0否 1是',
  `view_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '浏览量',
  `like_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '点赞数',
  `comment_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '评论数',
  `collect_count` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '收藏数',
  `last_comment_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后评论ID',
  `last_comment_at` DATETIME NULL DEFAULT NULL COMMENT '最后评论时间',
  `created_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发帖时间',
  `updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` DATETIME NULL DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_board_id` (`board_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='论坛帖子';

-- -----------------------------------------------------------------------------
-- 9. 论坛评论
-- -----------------------------------------------------------------------------

-- -----------------------------------------------------------------------------
-- 10. 收容库
-- -----------------------------------------------------------------------------
DROP TABLE IF EXISTS `containment_repo`;
CREATE TABLE `containment_repo` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `anomaly_name` VARCHAR(128) NOT NULL DEFAULT '' COMMENT '异常体名称',
  `agent_name` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '特工名称',
  `repo_name` VARCHAR(128) NOT NULL DEFAULT '' COMMENT '收容库名称/代码',
  `department` INT NOT NULL DEFAULT 0 COMMENT '部门',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='收容库';

SET FOREIGN_KEY_CHECKS = 1;

-- =============================================================================
-- 执行完成。可运行：gf gen dao
-- =============================================================================
