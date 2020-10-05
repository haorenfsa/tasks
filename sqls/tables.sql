CREATE TABLE IF NOT EXISTS task (
    id BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT 'primary key',
    parent_task_id BIGINT(20) NOT NULL DEFAULT -1 'parent task id, -1 means no parent',
    name VARCHAR(128) NOT NULL COMMENT 'name, unique',
    year SMALLINT(5) NOT NULL DEFAULT -1 COMMENT 'planed year, -1 means not set',
    month TINYINT(2) NOT NULL DEFAULT -1 COMMENT 'planed month, -1 means not set',
    week TINYINT(2) NOT NULL DEFAULT -1 COMMENT 'planed week, -1 means not set',
    day TINYINT(1) NOT NULL DEFAULT -1 COMMENT 'planed day, -1 means not set',
    status TINYINT NOT NULL DEFAULT 0 COMMENT 'task status',
    position BIGINT(20) NOT NULL DEFAULT 0 COMMENT 'position for order by',
    created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time auto',
    updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update time auto',
    PRIMARY KEY (id),
    UNIQUE KEY uk_name (name),
    UNIQUE KEY uk_position (position)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='task';


CREATE TABLE IF NOT EXISTS task_description (
    id BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT 'primary key',
    task_id BIGINT(20) NOT NULL COMMENT 'task id',
    description TEXT NOT NULL COMMENT 'description of the task',
    created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time auto',
    updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update time auto',
    PRIMARY KEY (id),
    UNIQUE KEY uk_task_id (task_id)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='task description';