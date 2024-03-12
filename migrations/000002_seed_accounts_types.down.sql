-- DROP SEED DATA FROM OPERATION_TYPES TABLE

BEGIN;
Delete from operation_types where description = 'Normal Purchase';
Delete from operation_types where description = 'Purchase with installments';
Delete from operation_types where description = 'Withdrawal';
Delete from operation_types where description = 'Credit Voucher';
COMMIT;