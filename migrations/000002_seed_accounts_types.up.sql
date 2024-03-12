-- INSERT SEED DATA IN OPERATION_TYPES TABLE
BEGIN;

INSERT INTO operation_types (description) VALUES ('Normal Purchase');
INSERT INTO operation_types (description) VALUES ('Purchase with installments');
INSERT INTO operation_types (description) VALUES ('Withdrawal');
INSERT INTO operation_types (description) VALUES ('Credit Voucher');

COMMIT;