CREATE TABLE IF NOT EXISTS  `task` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT "primary key",
    `name` VARCHAR(128) NOT NULL COMMENT "name, unique",
    `due_time` datetime NOT NULL COMMENT "due time of task",
    `status` TINYINT(8) NOT NULL COMMENT "task status",
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "create time auto",
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT "update time auto",
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='tasks';