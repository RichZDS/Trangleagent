-- =============================================================================
-- 三角机构 TRPG 数据库结构优化
-- 规则：每个用户可创建多个角色，每个角色拥有不同的素质保障（QA）
-- =============================================================================

-- 1. 用户表：添加当前选中角色 ID，便于进入房间时默认使用
ALTER TABLE `users` ADD COLUMN `active_role_id` BIGINT UNSIGNED NULL DEFAULT NULL COMMENT '当前选中的角色卡ID';
ALTER TABLE `users` ADD INDEX `idx_users_active_role_id` (`active_role_id`);

-- 2. 用户表：轨道字段标记弃用（数据以 role_cards 为准）
-- reality_role/abnormal_role/job_title 保留为用户级身份描述
ALTER TABLE `users` MODIFY COLUMN `red_trace` INT DEFAULT 0 COMMENT '【弃用】红轨，请使用 role_cards.red_track';
ALTER TABLE `users` MODIFY COLUMN `yellow_trace` INT DEFAULT 0 COMMENT '【弃用】黄轨，请使用 role_cards.yellow_track';
ALTER TABLE `users` MODIFY COLUMN `blue_trace` INT DEFAULT 0 COMMENT '【弃用】蓝轨，请使用 role_cards.blue_track';

-- 3. 角色卡表：添加索引，优化按用户/部门查询
ALTER TABLE `role_cards` ADD INDEX `idx_user_id` (`user_id`);
ALTER TABLE `role_cards` ADD INDEX `idx_department_id` (`department_id`);

-- 4. 轨道与 QA 范围约束（可选，需 MySQL 8.0.16+，若已有越界数据请先修正）
-- ALTER TABLE role_cards ADD CONSTRAINT chk_blue_track CHECK (blue_track <= 40);
-- ALTER TABLE role_cards ADD CONSTRAINT chk_yellow_track CHECK (yellow_track <= 40);
-- ALTER TABLE role_cards ADD CONSTRAINT chk_red_track CHECK (red_track <= 40);
-- ALTER TABLE role_cards ADD CONSTRAINT chk_qa_meticulous CHECK (qa_meticulous <= 100);
-- ... 其余 QA 同理

-- 5. 新建：房间成员表（用户加入房间时选用角色）
-- 记录：谁(user) 在哪个房间(room) 使用哪个角色(role_card)
CREATE TABLE IF NOT EXISTS `trpg_room_member` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `room_id` BIGINT UNSIGNED NOT NULL COMMENT '房间ID',
  `user_id` BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
  `role_card_id` BIGINT UNSIGNED NOT NULL COMMENT '使用的角色卡ID',
  `joined_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '加入时间',
  `left_at` DATETIME NULL DEFAULT NULL COMMENT '离开时间',
  `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：0已离开 1在线',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_room_user` (`room_id`, `user_id`) COMMENT '同一房间同一用户仅一条有效记录',
  KEY `idx_room_id` (`room_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_role_card_id` (`role_card_id`),
  CONSTRAINT `fk_room_member_room` FOREIGN KEY (`room_id`) REFERENCES `trpg_room` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_room_member_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_room_member_role` FOREIGN KEY (`role_card_id`) REFERENCES `role_cards` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='TRPG房间成员：用户-房间-角色关联';

-- 执行后请运行: gf gen dao，并在 hack/config.yaml 的 table 中加入 trpg_room_member
