BEGIN;

CREATE TABLE IF NOT EXISTS users (
    id                  bigint              NOT NULL GENERATED ALWAYS AS IDENTITY,
    full_name           text                NOT NULL,
    primary_email       text                NOT NULL,
    passwd              text,
    created_at          timestamptz         NOT NULL DEFAULT NOW(),
    updated_at          timestamptz         NOT NULL DEFAULT NOW(),
    deleted_at          timestamptz,
    CONSTRAINT pk_users PRIMARY KEY (id)
);

------------------------ Accounts

DROP TYPE IF EXISTS t_account_type RESTRICT;
CREATE TYPE t_account_type AS ENUM ('savings', 'checking', 'investment', 'debt');

CREATE TABLE IF NOT EXISTS financial_accounts (
    id                  bigint              NOT NULL GENERATED ALWAYS AS IDENTITY,
    title               text                NOT NULL,
    cents               bigint,
    account_type        t_account_type      NOT NULL,
    user_id             bigint              NOT NULL, -- foreign key
    created_at          timestamptz         NOT NULL DEFAULT NOW(),
    updated_at          timestamptz         NOT NULL DEFAULT NOW(),
    deleted_at          timestamptz,
    CONSTRAINT pk_financial_accounts PRIMARY KEY (id),
    CONSTRAINT fk_financial_accounts__users
        FOREIGN KEY(user_id)
	    REFERENCES users(id)
	    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS financial_accounts_history (
    id                  bigint              NOT NULL GENERATED ALWAYS AS IDENTITY,
    cents               bigint,
    financial_account_id bigint             NOT NULL, -- foreign key
    user_id             bigint              NOT NULL, -- foreign key
    created_at          timestamptz         NOT NULL DEFAULT NOW(),
    updated_at          timestamptz         NOT NULL DEFAULT NOW(),
    deleted_at          timestamptz,
    CONSTRAINT pk_financial_accounts_history PRIMARY KEY (id),
    CONSTRAINT fk_financial_accounts_history__users
        FOREIGN KEY(user_id)
	    REFERENCES users(id)
	    ON DELETE CASCADE,
    CONSTRAINT fk_financial_accounts_history__financial_accounts
        FOREIGN KEY(financial_account_id)
	    REFERENCES financial_accounts(id)
	    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS assets (
    id                  bigint              NOT NULL GENERATED ALWAYS AS IDENTITY,
    title               text                NOT NULL,
    cents_value         bigint,
    annual_appreciation_pct smallint,
    user_id             bigint              NOT NULL, -- foreign key
    created_at          timestamptz         NOT NULL DEFAULT NOW(),
    updated_at          timestamptz         NOT NULL DEFAULT NOW(),
    deleted_at          timestamptz,
    CONSTRAINT pk_assets PRIMARY KEY (id),
    CONSTRAINT fk_assets__users
        FOREIGN KEY(user_id)
	    REFERENCES users(id)
	    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS assets_history (
    id                  bigint              NOT NULL GENERATED ALWAYS AS IDENTITY,
    cents_value         bigint,
    annual_appreciation_pct smallint,
    user_id             bigint              NOT NULL, -- foreign key
    asset_id            bigint              NOT NULL, -- foreign key
    created_at          timestamptz         NOT NULL DEFAULT NOW(),
    updated_at          timestamptz         NOT NULL DEFAULT NOW(),
    deleted_at          timestamptz,
    CONSTRAINT pk_assets_history PRIMARY KEY (id),
    CONSTRAINT fk_assets_history__users
        FOREIGN KEY(user_id)
	    REFERENCES users(id)
	    ON DELETE CASCADE,
    CONSTRAINT fk_assets_history__assets
        FOREIGN KEY(asset_id)
        REFERENCES assets(id)
        ON DELETE CASCADE
);

------------------------ Income

DROP TYPE IF EXISTS t_income_type RESTRICT;
CREATE TYPE t_income_type AS ENUM ('job', 'side_gig', 'investment', 'gift');
DROP TYPE IF EXISTS t_income_freq RESTRICT;
CREATE TYPE t_income_freq AS ENUM ('regular', 'irregular', 'once');

CREATE TABLE IF NOT EXISTS incomes (
    id                  bigint              NOT NULL GENERATED ALWAYS AS IDENTITY,
    title               text                NOT NULL,
    cents               bigint,
    user_id             bigint              NOT NULL, -- foreign key
    income_type         t_income_type       NOT NULL,
    income_freq         t_income_freq       NOT NULL,
    start_date          timestamptz,
    end_date            timestamptz,
    destination_account_id bigint, -- foreign key
    created_at          timestamptz         NOT NULL DEFAULT NOW(),
    updated_at          timestamptz         NOT NULL DEFAULT NOW(),
    deleted_at          timestamptz,
    CONSTRAINT pk_incomes PRIMARY KEY (id),
    CONSTRAINT fk_incomes__users
        FOREIGN KEY(user_id)
	    REFERENCES users(id)
	    ON DELETE CASCADE,
    CONSTRAINT fk_incomes__destination_account
        FOREIGN KEY(destination_account_id)
        REFERENCES financial_accounts(id)
        ON DELETE SET NULL
);

------------------------ Budgets

DROP TYPE IF EXISTS t_money_category RESTRICT;
CREATE TYPE t_money_category AS ENUM ('expense', 'investment');

CREATE TABLE IF NOT EXISTS money_categories (
    id                  bigint              NOT NULL GENERATED ALWAYS AS IDENTITY,
    title               text                NOT NULL,
    category_type       t_money_category    NOT NULL,
    user_id             bigint              NOT NULL, -- foreign key
    created_at          timestamptz         NOT NULL DEFAULT NOW(),
    updated_at          timestamptz         NOT NULL DEFAULT NOW(),
    deleted_at          timestamptz,
    CONSTRAINT pk_money_categories PRIMARY KEY (id),
    CONSTRAINT fk_money_categories__users
        FOREIGN KEY(user_id)
	    REFERENCES users(id)
	    ON DELETE CASCADE
);

DROP TYPE IF EXISTS t_expense_type RESTRICT;
CREATE TYPE t_expense_type AS ENUM ('variable', 'fixed', 'future');

CREATE TABLE IF NOT EXISTS expense_items (
    id                  bigint              NOT NULL GENERATED ALWAYS AS IDENTITY,
    title               text                NOT NULL,
    expense_type        t_expense_type      NOT NULL,
    money_category_id   bigint              NOT NULL, -- foreign key
    user_id             bigint              NOT NULL, -- foreign key
    cents               bigint,
    start_date          timestamptz,
    end_date            timestamptz,
    pct_extra_income    smallint,
    destination_account_id bigint, -- foreign key
    created_at          timestamptz         NOT NULL DEFAULT NOW(),
    updated_at          timestamptz         NOT NULL DEFAULT NOW(),
    deleted_at          timestamptz,
    CONSTRAINT pk_expense_items PRIMARY KEY (id),
    CONSTRAINT fk_expense_items__money_categories
        FOREIGN KEY(money_category_id)
	    REFERENCES money_categories(id)
	    ON DELETE CASCADE,
    CONSTRAINT fk_expense_items__users
        FOREIGN KEY(user_id)
	    REFERENCES users(id)
	    ON DELETE CASCADE,
    CONSTRAINT fk_expense_items__destination_account
        FOREIGN KEY(destination_account_id)
        REFERENCES financial_accounts(id)
        ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS actual_expenses (
    id                  bigint              NOT NULL GENERATED ALWAYS AS IDENTITY,
    money_category_id   bigint              NOT NULL, -- foreign key
    expense_item_id     bigint              NOT NULL, -- foreign key
    user_id             bigint              NOT NULL, -- foreign key
    cents               bigint,
    created_at          timestamptz         NOT NULL DEFAULT NOW(),
    updated_at          timestamptz         NOT NULL DEFAULT NOW(),
    deleted_at          timestamptz,
    CONSTRAINT pk_actual_expenses PRIMARY KEY (id),
    CONSTRAINT fk_actual_expenses__money_categories
        FOREIGN KEY(money_category_id)
	    REFERENCES money_categories(id)
	    ON DELETE CASCADE,
    CONSTRAINT fk_actual_expenses__expense_items
        FOREIGN KEY(expense_item_id)
	    REFERENCES expense_items(id)
	    ON DELETE CASCADE,
    CONSTRAINT fk_actual_expenses__users
        FOREIGN KEY(user_id)
	    REFERENCES users(id)
	    ON DELETE CASCADE
);

COMMIT;
