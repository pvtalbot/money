BEGIN;


CREATE TABLE IF NOT EXISTS expenses_categories(
  id INT NOT NULL UNIQUE AUTO_INCREMENT,
  name VARCHAR (127),
  user_id INT NOT NULL,
  PRIMARY KEY(id),
  CONSTRAINT `fk_expense_category_user`
    FOREIGN KEY (user_id) REFERENCES users(id)
);

ALTER TABLE expenses
ADD expense_category_id INT,
ADD CONSTRAINT `fk_expense_expense_category`
  FOREIGN KEY (expense_category_id) REFERENCES expenses_categories(id);


COMMIT;