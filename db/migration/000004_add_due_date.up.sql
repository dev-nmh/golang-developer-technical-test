ALTER TABLE tr_loan_detail
ADD COLUMN `due_date`  datetime NOT NULL AFTER `transaction_date`;