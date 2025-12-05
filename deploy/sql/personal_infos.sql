CREATE TABLE `personal_infos` (
  `name` longtext,
  `gender` longtext,
  `student_id` varchar(191) NOT NULL,
  `college` longtext,
  `dormitory` longtext,
  `contact` longtext,
  PRIMARY KEY (`student_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

