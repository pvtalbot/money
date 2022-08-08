BEGIN;


CREATE TABLE IF NOT EXISTS expenses_categories(
  id INT NOT NULL UNIQUE AUTO_INCREMENT,
  name VARCHAR (127),
  PRIMARY KEY(id)
);

ALTER TABLE expenses
ADD expense_category_id INT NOT NULL,
ADD CONSTRAINT `fk_expense_expense_category`
  FOREIGN KEY (expense_category_id) REFERENCES expenses_categories(id);


COMMIT;