CREATE TABLE `notices` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` longtext,
  `publish_time` timestamp NULL DEFAULT NULL,
  `img1` longtext,
  `img2` longtext,
  `img3` longtext,
  `publisher` longtext,
  `content` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
