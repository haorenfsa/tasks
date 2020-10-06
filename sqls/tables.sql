CREATE TABLE IF NOT EXISTS task (
    id BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT 'primary key',
    parent_task_id BIGINT(20) NOT NULL DEFAULT -1 COMMENT 'parent task id, -1 means no parent',
    name VARCHAR(128) NOT NULL COMMENT 'name, unique',
    status TINYINT NOT NULL DEFAULT 0 COMMENT 'task status',
    position BIGINT(20) NOT NULL DEFAULT 0 COMMENT 'position for order by',
    schedule_level INT(8) NOT NULL DEFAULT 0 COMMENT 'time level of start_time and end_time like year,month,week, 0 means not set',
    start_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'start time',
    end_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'end time',
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