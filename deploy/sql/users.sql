CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `username` longtext,
  `type` bigint DEFAULT NULL,
  `student_id` longtext,
  `jh_password` longtext,
  `wechat_open_id` longtext,
  `union_id` longtext,
  `lib_password` longtext,
  `zf_password` longtext,
  `device_id` longtext,
  `phone_num` longtext,
  `yxy_uid` longtext,
  `oauth_password` longtext,
  `create_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

