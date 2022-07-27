BEGIN;


ALTER TABLE Users 
ADD first_name VARCHAR (127),
ADD last_name VARCHAR (127); 

UPDATE Users
SET first_name = "paul", last_name = "talbot"
WHERE id = 2;

UPDATE Users
SET first_name = "paul", last_name = "talbot"
WHERE id = 3;

ALTER TABLE Users
MODIFY COLUMN first_name VARCHAR (127) NOT NULL,
MODIFY COLUMN last_name VARCHAR (127) NOT NULL;


COMMIT;