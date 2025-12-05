CREATE TABLE `configs` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `key` longtext,
  `value` longtext,
  `update_time` timestamp NULL DEFAULT NULL COMMENT '''设置时间''',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

