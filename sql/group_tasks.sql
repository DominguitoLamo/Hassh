CREATE TABLE
    `group_tasks` (
        `id` int NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
        `group_id` int NOT NULL,
        `task_id` int NOT NULL,
        PRIMARY KEY (`id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = 'record group_id to task_id. 1 to x'