-- 1. 合并 fans 与 subscribe 表为 user_follows 表
DROP TABLE IF EXISTS `fans`;
DROP TABLE IF EXISTS `subscribe`;

CREATE TABLE `user_follows` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `follower_id` BIGINT UNSIGNED NOT NULL COMMENT '关注者（粉丝）ID',
  `followed_id` BIGINT UNSIGNED NOT NULL COMMENT '被关注者ID',
  `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1=关注中 0=取消关注',
  `remark` VARCHAR(128) NOT NULL DEFAULT '' COMMENT '备注名',
  `created_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_follow` (`follower_id`, `followed_id`),
  KEY `idx_followed_id` (`followed_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户关注关系表';

-- 2. 删除 forum_comments 表（复用现有的多态 comments 表）
DROP TABLE IF EXISTS `forum_comments`;

-- 3. 剥离 users 表中冗余的 ARC 角色字段，并新增 extra_info
ALTER TABLE `users` 
  DROP COLUMN `reality_role`,
  DROP COLUMN `abnormal_role`,
  DROP COLUMN `job_title`,
  ADD COLUMN `extra_info` JSON NULL COMMENT '扩展配置信息';

-- 4. role_cards 表新增 extra_info 字段（可以用于存放 qa_metrics 等）
ALTER TABLE `role_cards`
  ADD COLUMN `extra_info` JSON NULL COMMENT '扩展配置信息';
