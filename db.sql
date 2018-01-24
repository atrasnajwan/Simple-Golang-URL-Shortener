CREATE TABLE `shorten` (
    `id` VARCHAR(64) NOT NULL,
    `shorturl` VARCHAR(64) NOT NULL,
    `longurl` VARCHAR(64) NOT NULL,
    PRIMARY KEY (`id`)
);