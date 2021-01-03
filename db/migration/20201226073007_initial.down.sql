BEGIN;

--------

DROP TABLE IF EXISTS actual_expenses;

DROP TABLE IF EXISTS expense_items;
DROP TYPE IF EXISTS t_expense_type RESTRICT;

DROP TABLE IF EXISTS money_categories;
DROP TYPE IF EXISTS t_money_category RESTRICT;

--------

DROP TABLE IF EXISTS incomes;
DROP TYPE IF EXISTS t_income_freq RESTRICT;
DROP TYPE IF EXISTS t_income_type RESTRICT;

--------

DROP TABLE IF EXISTS assets_history;

DROP TABLE IF EXISTS assets;

DROP TABLE IF EXISTS financial_accounts_history;

DROP TABLE IF EXISTS financial_accounts;
DROP TYPE IF EXISTS t_account_type RESTRICT;

--------

DROP TABLE IF EXISTS users;

--------

COMMIT;