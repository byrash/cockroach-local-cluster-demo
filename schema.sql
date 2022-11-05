CREATE TABLE parent (
	id serial PRIMARY KEY,
	name VARCHAR ( 10 ) NOT NULL
);

CREATE TABLE child (
	id serial PRIMARY KEY,
	name VARCHAR ( 10 ) NOT NULL,
	parent_id INT NOT NULL,
	FOREIGN KEY (parent_id) REFERENCES parent (id)
);

-- TODO: https://www.cockroachlabs.com/docs/v22.1/serial.html