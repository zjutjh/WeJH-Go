CREATE TABLE `announcements` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` longtext,
  `content` longtext,
  `publish_time` timestamp NULL DEFAULT NULL COMMENT '''发布时间''',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

