-- 签到功能：每日签到增加经验
ALTER TABLE `users` ADD COLUMN `last_checkin_at` DATETIME NULL DEFAULT NULL COMMENT '上次签到时间' AFTER `level`;
