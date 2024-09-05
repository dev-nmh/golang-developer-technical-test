CREATE DATABASE IF NOT EXISTS golang_developer_technical_test;
USE golang_developer_technical_test;

CREATE TABLE ms_billing_status (
    `pk_ms_billing_status` int  NOT  NULL AUTO_INCREMENT PRIMARY KEY,
    `title` varchar(50) NOT NULL,
    `description` varchar(255),
    `is_active` tinyint(1) NOT NULL DEFAULT 1,
    `created_by` varchar(50) NOT NULL DEFAULT 'sysadmin',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_by` varchar(50) NOT NULL DEFAULT 'sysadmin',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
)ENGINE=InnoDB;

CREATE TABLE ms_payment_status (
    `pk_ms_payment_status` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `title` varchar(50) NOT NULL,
    `description` varchar(255),
    `is_active` tinyint(1) NOT NULL DEFAULT 1,
    `created_by` varchar(50) NOT NULL DEFAULT 'sysadmin',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_by` varchar(50) NOT NULL DEFAULT 'sysadmin',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
)ENGINE=InnoDB;

CREATE TABLE ms_item_type (
    `pk_ms_item_type` VARCHAR(36) UNIQUE NOT NULL PRIMARY KEY,
    `title` varchar(50) UNIQUE NOT NULL,
    `description` varchar(255),
    `is_active` tinyint(1) NOT NULL DEFAULT 1,
    `created_by` varchar(50) NOT NULL DEFAULT 'sysadmin',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_by` varchar(50) NOT NULL DEFAULT 'sysadmin',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
)ENGINE=InnoDB;

CREATE TABLE ms_user(
    `pk_ms_user` varchar(36) UNIQUE NOT NULL PRIMARY KEY,
    `NIK` varchar (16) UNIQUE NOT NULL,
    `full_name` varchar(100) NOT NULL,
    `legal_name` varchar(60) NOT NULL,
    `birth_place` varchar(50) NOT NULL,
    `birth_date` datetime NOT NULL,
    `salary` int,
    `image_ktp` varchar(36) NOT NULL,
    `image_selfie` varchar(36) NOT NULL,
    `is_active` tinyint(1) NOT NULL DEFAULT 1,
    `created_by` varchar(50) NOT NULL DEFAULT 'sysadmin',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_by` varchar(50) NOT NULL DEFAULT 'sysadmin',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
)ENGINE=InnoDB;

CREATE TABLE ms_source(
    `pk_ms_source` varchar(36) UNIQUE NOT NULL PRIMARY KEY,
    `title` varchar(50) UNIQUE NOT NULL,
    `description` varchar(100) NOT NULL,
    `admin_fee` varchar(60) NOT NULL,
    `is_active` tinyint(1) NOT NULL DEFAULT 1,
    `created_by` varchar(50) NOT NULL DEFAULT 'sysadmin',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_by` varchar(50) NOT NULL DEFAULT 'sysadmin',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
)ENGINE=InnoDB;

CREATE TABLE ms_tenor(
    `pk_ms_tenor` varchar(36) UNIQUE NOT null PRIMARY KEY,
    `tenor_months` int NOT NULL,
    `interest_rate_percent` decimal(7,4) NOT NULL,
    `is_active` tinyint(1) NOT NULL DEFAULT 1,
    `created_by` varchar(50) NOT NULL DEFAULT 'sysadmin',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_by` varchar(50) NOT NULL DEFAULT 'sysadmin',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
)ENGINE=InnoDB;


CREATE TABLE map_user_tenor(
    `pk_map_user_tenor` varchar(36) UNIQUE NOT NULL PRIMARY KEY,
    `fk_ms_user` varchar(36) UNIQUE NOT NULL,
    `fk_ms_tenor` varchar(36) UNIQUE NOT NULL,
    `amount` decimal(20,2) NOT NULL,
    `is_active` tinyint(1) NOT NULL DEFAULT 1,
    `created_by` varchar(50) NOT NULL DEFAULT 'sysadmin',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_by` varchar(50) NOT NULL DEFAULT 'sysadmin',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
)ENGINE=InnoDB;


CREATE TABLE tr_loan_header(
    `pk_tr_loan_header` varchar(36) UNIQUE NOT NULL PRIMARY KEY,
    `fk_ms_user` varchar(36) UNIQUE NOT NULL,
    `fk_ms_payment_status` int NOT NULL,
    `fk_ms_item_type`varchar(36) UNIQUE NOT NULL,
    `contract_number` varchar(60) UNIQUE NOT NULL,
    `asset_name` varchar(60) NOT NULL,
    `is_active` tinyint(1) NOT NULL DEFAULT 1,
    `created_by` varchar(50) NOT NULL DEFAULT 'sysadmin',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_by` varchar(50) NOT NULL DEFAULT 'sysadmin',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(contract_number,fk_ms_user)
)ENGINE=InnoDB;

CREATE TABLE tr_loan_detail(
    `pk_tr_loan_detail` varchar(36) UNIQUE NOT NULL PRIMARY KEY,
    `fk_tr_loan_header` varchar(36) NOT NULL,
    `fk_ms_source` varchar(36) NOT NULL,
    `fk_map_user_tenor` varchar(36) NOT NULL,
    `otr_amount` decimal(20,2) NOT NULL,
    `loan_balance` decimal(20,2) NOT NULL,
    `transaction_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `is_active` tinyint(1) NOT NULL DEFAULT 1,
    `created_by` varchar(50) NOT NULL DEFAULT 'sysadmin',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_by` varchar(50) NOT NULL DEFAULT 'sysadmin',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
)ENGINE=InnoDB;


CREATE TABLE tr_loan_billing(
    `pk_tr_loan_billing` varchar(36) UNIQUE NOT NULL PRIMARY KEY,
    `fk_tr_loan_header` varchar(36) NOT NULL,
    `fk_ms_billing_status` int NOT NULL,
    `sort_order` int NOT NULL,
    `payoff_balance` decimal(20,2) NOT NULL,
    `expired_date` datetime NOT NULL,
    `is_active` tinyint(1) NOT NULL DEFAULT 1,
    `created_by` varchar(50) NOT NULL DEFAULT 'sysadmin',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_by` varchar(50) NOT NULL DEFAULT 'sysadmin',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(sort_order,fk_tr_loan_header)
)ENGINE=InnoDB;

/*Constraint Foreign Key*/
/** map_user_tenor**/
ALTER TABLE map_user_tenor 
ADD CONSTRAINT user_map_user_tenor FOREIGN KEY (fk_ms_user) REFERENCES ms_user(pk_ms_user);
ALTER TABLE map_user_tenor 
ADD CONSTRAINT tenor_map_user_tenor FOREIGN KEY (fk_ms_tenor) REFERENCES ms_tenor(pk_ms_tenor);


/** tr_loan_header**/
ALTER TABLE tr_loan_header 
ADD CONSTRAINT user_tr_loan_header FOREIGN KEY (fk_ms_user) REFERENCES ms_user(pk_ms_user);
ALTER TABLE tr_loan_header 
ADD CONSTRAINT status_tr_loan_header FOREIGN KEY (fk_ms_payment_status) REFERENCES ms_payment_status(pk_ms_payment_status);
ALTER TABLE tr_loan_header 
ADD CONSTRAINT item_type_tr_loan_header FOREIGN KEY (fk_ms_item_type) REFERENCES ms_item_type(pk_ms_item_type);


/** tr_loan_header**/
ALTER TABLE tr_loan_detail
ADD CONSTRAINT header_tr_loan_detail FOREIGN KEY (fk_tr_loan_header) REFERENCES tr_loan_header(pk_tr_loan_header);
ALTER TABLE tr_loan_detail
ADD CONSTRAINT source_tr_loan_detail FOREIGN KEY (fk_ms_source) REFERENCES ms_source(pk_ms_source);
ALTER TABLE tr_loan_detail
ADD CONSTRAINT user_tenor_tr_loan_detail FOREIGN KEY (fk_map_user_tenor) REFERENCES map_user_tenor(pk_map_user_tenor);


/**tr_loan_billing**/
ALTER TABLE tr_loan_billing
ADD CONSTRAINT header_tr_loan_billing FOREIGN KEY (fk_tr_loan_header) REFERENCES tr_loan_header(pk_tr_loan_header);
ALTER TABLE tr_loan_billing
ADD CONSTRAINT status_tr_loan_billing FOREIGN KEY (fk_ms_billing_status) REFERENCES ms_billing_status(pk_ms_billing_status);

