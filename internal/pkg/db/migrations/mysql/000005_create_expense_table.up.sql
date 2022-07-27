CREATE TABLE IF NOT EXISTS expenses(
    ID INT NOT NULL UNIQUE AUTO_INCREMENT,
    amount INT NOT NULL,
    user_id INT NOT NULL,
    CONSTRAINT `fk_expense_user`
        FOREIGN KEY (user_id) REFERENCES Users(id)
        ON DELETE CASCADE
);