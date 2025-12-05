CREATE TABLE `school_buses` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `line` longtext,
  `departure` longtext,
  `destination` longtext,
  `type` bigint DEFAULT NULL,
  `start_time` longtext,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

