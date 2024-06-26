CREATE DATABASE IF NOT EXISTS referral_db;

USE referral_db;

CREATE TABLE IF NOT EXISTS referrals (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user VARCHAR(50) NOT NULL,
    hash VARCHAR(255) UNIQUE,
    promo_code VARCHAR(50) UNIQUE,
    count INT DEFAULT 0,
    date_generated DATETIME,
    date_registered DATETIME
);
