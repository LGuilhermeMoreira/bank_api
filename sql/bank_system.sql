CREATE TABLE IF NOT EXISTS `login_accounts` (
  `id` VARCHAR(36) PRIMARY KEY UNIQUE,  -- assuming UUIDs are 36 characters long
  `user_mail` VARCHAR (255) NOT NULL UNIQUE,
  `user_password` VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS `accounts` (
  `id` VARCHAR(36) PRIMARY KEY UNIQUE,
  `login_account_id` VARCHAR(36) NOT NULL,
  `owner` VARCHAR(255) NOT NULL,
  `balance` FLOAT NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (`login_account_id`) REFERENCES `login_accounts` (`id`) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS `entries` (
    `id` VARCHAR(36) PRIMARY KEY UNIQUE,
    `account_id` VARCHAR(36) NOT NULL,
    `amount` FLOAT NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS `transfers` (
    `id` VARCHAR(36) PRIMARY KEY UNIQUE,
    `from_account_id` VARCHAR(36) NOT NULL,
    `to_account_id` VARCHAR(36) NOT NULL,
    `value` FLOAT NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`to_account_id`) REFERENCES `accounts` (`id`),
    FOREIGN KEY (`from_account_id`) REFERENCES `accounts` (`id`)
);
