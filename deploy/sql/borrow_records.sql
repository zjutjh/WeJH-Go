CREATE TABLE `borrow_records` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` longtext,
  `gender` longtext,
  `student_id` longtext,
  `college` longtext,
  `dormitory` longtext,
  `contact` longtext,
  `campus` tinyint unsigned DEFAULT NULL,
  `supplies_id` bigint DEFAULT NULL,
  `count` bigint unsigned DEFAULT NULL,
  `organization` longtext,
  `apply_time` timestamp NULL DEFAULT NULL,
  `borrow_time` timestamp NULL DEFAULT NULL,
  `return_time` timestamp NULL DEFAULT NULL,
  `status` bigint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

