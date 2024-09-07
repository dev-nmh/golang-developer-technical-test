CREATE TABLE ms_approval_status (
    `pk_ms_approval_status` int  NOT  NULL AUTO_INCREMENT PRIMARY KEY,
    `title` varchar(50) NOT NULL,
    `description` varchar(255),
    `is_active` tinyint(1) NOT NULL DEFAULT 1,
    `created_by` varchar(50) NOT NULL DEFAULT 'sysadmin',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_by` varchar(50) NOT NULL DEFAULT 'sysadmin',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
)ENGINE=InnoDB;

CREATE TABLE ms_account
(
    `pk_ms_account`       varchar(36) UNIQUE not null PRIMARY KEY,
    `fk_ms_role`           int not null,
    `email`               varchar(100) UNIQUE not null,
    `password`          varchar(100) not null,
    `password_salt`     varchar(255) not null,
    `is_active` tinyint(1) NOT NULL DEFAULT 1,
    `created_by` varchar(50) NOT NULL DEFAULT 'sysadmin',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_by` varchar(50) NOT NULL DEFAULT 'sysadmin',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) engine = InnoDB;

CREATE TABLE ms_role (
    `pk_ms_role` int  NOT  NULL AUTO_INCREMENT PRIMARY KEY,
    `title` varchar(50) NOT NULL,
    `description` varchar(255),
    `is_active` tinyint(1) NOT NULL DEFAULT 1,
    `created_by` varchar(50) NOT NULL DEFAULT 'sysadmin',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_by` varchar(50) NOT NULL DEFAULT 'sysadmin',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
)ENGINE=InnoDB;


ALTER TABLE ms_user
ADD COLUMN `fk_ms_account` VARCHAR(36) UNIQUE NOT NULL AFTER `pk_ms_user`;

ALTER TABLE ms_user
ADD COLUMN `fk_ms_approval_status`  int NOT NULL AFTER `fk_ms_account` ON DELETE CASCADE;



ALTER TABLE ms_account
ADD CONSTRAINT role_ms_account FOREIGN KEY (fk_ms_role) REFERENCES ms_role(pk_ms_role);

ALTER TABLE ms_user 
ADD CONSTRAINT approval_status_ms_user FOREIGN KEY (fk_ms_approval_status) REFERENCES ms_approval_status(pk_ms_approval_status);
ALTER TABLE ms_user
ADD CONSTRAINT account_ms_user FOREIGN KEY (fk_ms_account) REFERENCES ms_account(pk_ms_account);


INSERT INTO ms_approval_status(title,description)
VALUES ("REQUESTED","user is sending request account loan"),
       ("UNDER_REVIEWED","users is under review"),
       ("APPROVED","user request is approved"),
       ("REJECTED","user request is rejected");


INSERT INTO ms_role(title,description)
VALUES ("Admin","Admin Role"),
       ("User","User Role");
