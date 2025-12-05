CREATE TABLE `lessons` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `username` longtext,
  `campus` longtext,
  `lesson_name` longtext,
  `lesson_place` longtext,
  `sections` longtext,
  `week` longtext,
  `weekday` longtext,
  `term` longtext,
  `year` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

