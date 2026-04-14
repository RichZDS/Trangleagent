-- 创建通用多态点赞记录表
CREATE TABLE IF NOT EXISTS `user_likes` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `user_id` BIGINT UNSIGNED NOT NULL COMMENT '点赞用户ID',
  `target_type` VARCHAR(32) NOT NULL COMMENT '被点赞对象类型（post/comment）',
  `target_id` BIGINT UNSIGNED NOT NULL COMMENT '被点赞对象ID',
  `created_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_user_target` (`user_id`, `target_type`, `target_id`),
  KEY `idx_target` (`target_type`, `target_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户点赞记录表';
