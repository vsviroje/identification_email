CREATE TABLE IF NOT EXISTS `user` (
    `id` int NOT NULL AUTO_INCREMENT,
    `userName` varchar(50) NOT NULL,
    `name` varchar(50) DEFAULT '',
    `email` varchar(100) NOT NULL,
    `phoneNum` varchar(10) DEFAULT '',
    `password` varchar(1000) NOT NULL,
    `createdAt` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updatedAt` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);


ALTER TABLE `user` ADD UNIQUE KEY `unique_user`(`userName`, `email`);
