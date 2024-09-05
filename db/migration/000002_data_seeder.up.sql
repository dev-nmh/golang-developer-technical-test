INSERT INTO ms_billing_status(title,description)
VALUES ("PENDING","Payment is yet to be received"),
       ("PAID","Biling Belum Dibayar"),
       ("OVERDUE","Payment is overdue and pending"),
       ("CANCELED","Payment was canceled");


INSERT INTO ms_payment_status(title,description)
VALUES ("ACTIVE","Loan is active and in repayment"),
       ("PAID","Loan has been fully repaid"),
       ("DEFAULTED","Loan payments are overdue"),
       ("CANCELED","Payment was canceled");


INSERT INTO ms_source(pk_ms_source,title,description,admin_fee)
VALUES (uuid(),"Toxopedia","Ecommerce",2000.0),
       (uuid(),"Carmud","Ecommerce",30000),
       ( uuid(),"WEB PT XYZ","Web PT XYZ",5000);


INSERT INTO ms_tenor(pk_ms_tenor,tenor_months,interest_rate_percent)
VALUES (uuid(),1,1),
        (uuid(),2,2.5),
        (uuid(),3,4),
        (uuid(),4,5.5);
