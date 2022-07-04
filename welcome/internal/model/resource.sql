-- welcome.resource definition

CREATE TABLE `resource` (
                            `id` int NOT NULL AUTO_INCREMENT,
                            `x` int DEFAULT NULL,
                            `y` int DEFAULT NULL,
                            `width` int DEFAULT NULL,
                            `height` int DEFAULT NULL,
                            `type` varchar(100) DEFAULT NULL,
                            `path` varchar(100) DEFAULT NULL,
                            `thumbnail` varchar(100) DEFAULT NULL,
                            `is_thumbnail` tinyint(1) NOT NULL,
                            `alias` varchar(100) DEFAULT NULL,
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;