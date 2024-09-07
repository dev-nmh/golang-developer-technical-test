SET FOREIGN_KEY_CHECKS=0;
ALTER TABLE ms_user DROP CONSTRAINT  account_ms_user;
ALTER TABLE ms_user DROP CONSTRAINT  approval_status_ms_user;

ALTER TABLE ms_user DROP COLUMN `fk_ms_account`;
ALTER TABLE ms_user DROP COLUMN `fk_ms_approval_status`;
DROP TABLE if exists ms_account;
DROP TABLE if exists ms_approval_status;
DROP TABLE if exists ms_role;

SET FOREIGN_KEY_CHECKS=1;
