CREATE TABLE
    `group_info` (
        `id` int NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
        `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
        PRIMARY KEY (`id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = 'group_info for grouping ssh task'