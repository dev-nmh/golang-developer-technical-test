create table ms_user_account
(
    pk_ms_user_account                  varchar(100) not null,
    username                            varchar(100) not null,
    email                               varchar(100) not null,
    password                            varchar(100) not null,
    password_salt                       varchar(100) null,
    `is_active`         tinyint(1) NOT NULL DEFAULT 1,
    `created_by`        varchar(50) NOT NULL DEFAULT 'sysadmin',
    `created_at`        datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_by`        varchar(50) NOT NULL DEFAULT 'sysadmin',
    `updated_at`        datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
    primary key (id)
) engine = InnoDB;