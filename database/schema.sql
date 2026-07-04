-- Run this on a fresh Aiven MySQL if AutoMigrate isn't available
-- Note: GORM AutoMigrate handles this automatically on first run,
-- but you can pre-create if you want to inspect the schema.

USE defaultdb;

CREATE TABLE IF NOT EXISTS donors (
    id              BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    created_at      DATETIME(3) NULL,
    updated_at      DATETIME(3) NULL,
    deleted_at      DATETIME(3) NULL,
    name            VARCHAR(255) NOT NULL,
    phone           VARCHAR(255) NOT NULL,
    email           VARCHAR(255),
    blood_group     VARCHAR(10)  NOT NULL,
    age             INT,
    gender          VARCHAR(20),
    area            VARCHAR(255) NOT NULL,
    city            VARCHAR(255),
    last_donation   DATETIME(3),
    available       BOOLEAN DEFAULT TRUE,
    notes           TEXT,
    INDEX idx_blood_group (blood_group),
    INDEX idx_area (area),
    INDEX idx_deleted_at (deleted_at)
);

CREATE TABLE IF NOT EXISTS blood_requests (
    id              BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    created_at      DATETIME(3) NULL,
    updated_at      DATETIME(3) NULL,
    deleted_at      DATETIME(3) NULL,
    patient_name    VARCHAR(255) NOT NULL,
    blood_group     VARCHAR(10)  NOT NULL,
    units           INT NOT NULL,
    hospital        VARCHAR(255) NOT NULL,
    contact         VARCHAR(255) NOT NULL,
    area            VARCHAR(255) NOT NULL,
    needed_by       VARCHAR(255),
    urgency         VARCHAR(20) DEFAULT 'normal',
    status          VARCHAR(20) DEFAULT 'open',
    requested_by    VARCHAR(255),
    INDEX idx_req_blood (blood_group),
    INDEX idx_req_status (status),
    INDEX idx_req_deleted (deleted_at)
);

-- Demo Data for donors
INSERT INTO donors
(name, phone, email, blood_group, age, gender, area, city, last_donation, available, notes, created_at)
VALUES
('Rahim Uddin', '01711111111', 'rahim@gmail.com', 'A+', 28, 'Male', 'Dhanmondi', 'Dhaka', '2025-12-10 10:00:00', TRUE, 'Regular donor', NOW()),
('Karima Akter', '01822222222', 'karima@gmail.com', 'B+', 32, 'Female', 'Mirpur', 'Dhaka', '2025-11-05 14:30:00', TRUE, 'Can donate anytime', NOW()),
('Sabbir Hossain', '01933333333', NULL, 'O+', 24, 'Male', 'Uttara', 'Dhaka', '2025-10-01 09:15:00', FALSE, 'Recently donated', NOW()),
('Nusrat Jahan', '01644444444', 'nusrat@gmail.com', 'AB+', 29, 'Female', 'Chawkbazar', 'Chittagong', '2025-09-20 16:45:00', TRUE, 'Emergency donor', NOW()),
('Arif Hasan', '01555555555', 'arif@gmail.com', 'O-', 35, 'Male', 'Rajpara', 'Rajshahi', '2025-08-12 11:20:00', TRUE, 'Rare blood group', NOW());

-- Demo Data for blood_requests
INSERT INTO blood_requests
(patient_name, blood_group, units, hospital, contact, area, needed_by, urgency, status, requested_by, created_at)
VALUES
('Abdul Karim', 'A+', 2, 'Dhaka Medical College Hospital', '01799999991', 'Dhaka', '2026-01-10', 'urgent', 'open', 'Family', NOW()),
('Rina Begum', 'B+', 1, 'Square Hospital', '01799999992', 'Dhaka', '2026-01-12', 'normal', 'open', 'Friend', NOW()),
('Hasan Ali', 'O+', 3, 'United Hospital', '01799999993', 'Dhaka', '2026-01-15', 'urgent', 'open', 'Brother', NOW()),
('Mehedi Hasan', 'AB+', 2, 'Chittagong Medical College', '01799999994', 'Chittagong', '2026-01-18', 'normal', 'fulfilled', 'Hospital', NOW()),
('Salma Khatun', 'O-', 1, 'Rajshahi Medical College', '01799999995', 'Rajshahi', '2026-01-20', 'urgent', 'open', 'Relative', NOW());

SELECT * FROM donors;
SELECT * FROM blood_requests;

