CREATE TABLE IF NOT EXISTS `config` (
    `id` int NOT NULL AUTO_INCREMENT,
    `type` varchar(100) NOT NULL,
    `key` varchar(100) NOT NULL,
    `value` varchar(500),
    `creationAt` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updatedAt` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);	