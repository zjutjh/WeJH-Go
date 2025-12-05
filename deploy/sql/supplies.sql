CREATE TABLE `supplies` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` longtext,
  `kind` longtext,
  `spec` longtext,
  `campus` tinyint unsigned DEFAULT NULL,
  `organization` longtext,
  `stock` bigint unsigned DEFAULT NULL,
  `borrowed` bigint unsigned DEFAULT NULL,
  `img` longtext,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_supplies_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

