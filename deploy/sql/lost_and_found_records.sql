CREATE TABLE `lost_and_found_records` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `type` tinyint(1) DEFAULT NULL,
  `campus` longtext,
  `kind` longtext,
  `publish_time` timestamp NULL DEFAULT NULL,
  `is_processed` tinyint(1) DEFAULT NULL,
  `img1` longtext,
  `img2` longtext,
  `img3` longtext,
  `publisher` longtext,
  `item_name` longtext,
  `lost_or_found_place` longtext,
  `lost_or_found_time` longtext,
  `pickup_place` longtext,
  `contact` longtext,
  `introduction` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

