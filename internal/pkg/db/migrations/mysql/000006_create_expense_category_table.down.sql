BEGIN;


ALTER TABLE expenses
DROP FOREIGN KEY `fk_expense_expense_category`;

ALTER TABLE expenses
DROP COLUMN IF EXISTS expense_category_id;

DROP TABLE IF EXISTS expenses_categories;


COMMIT;