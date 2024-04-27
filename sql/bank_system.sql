CREATE TABLE IF NOT EXISTS `accounts` (
  `id` VARCHAR(255) PRIMARY KEY,
  `owner` VARCHAR(255) NOT NULL,
  `balance` FLOAT NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `entries` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
    `account_id` VARCHAR(255),
    `amount` FLOAT NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS `transfers` (
    `id` VARCHAR(255) PRIMARY KEY,
    `from_account_id` VARCHAR(255),
    `to_account_id` VARCHAR(255),
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`to_account_id`) REFERENCES `accounts` (`id`),
    FOREIGN KEY (`from_account_id`) REFERENCES `accounts` (`id`)
);
