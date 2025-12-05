CREATE TABLE `theme_permissions` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `student_id` longtext,
  `current_theme_id` bigint DEFAULT NULL,
  `current_theme_dark_id` bigint DEFAULT NULL,
  `theme_permission` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

