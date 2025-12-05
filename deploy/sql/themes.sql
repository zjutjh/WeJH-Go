CREATE TABLE `themes` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` longtext,
  `type` longtext,
  `is_dark_mode` tinyint(1) DEFAULT NULL,
  `theme_config` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

