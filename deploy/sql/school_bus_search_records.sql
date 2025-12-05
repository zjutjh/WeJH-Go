CREATE TABLE `school_bus_search_records` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `username` longtext,
  `departure` longtext,
  `destination` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

