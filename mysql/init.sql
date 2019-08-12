CREATE TABLE `USERS` (
    `id` bigint unsigned PRIMARY KEY AUTO_INCREMENT,
    `name` varchar(255) NOT NULL COMMENT 'user name',
    `age` integer NOT NULL COMMENT 'age'
)
ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC
