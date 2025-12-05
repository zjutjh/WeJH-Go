CREATE TABLE `app_lists` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` longtext,
  `route` longtext,
  `background_color` longtext,
  `icon` longtext,
  `require` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
