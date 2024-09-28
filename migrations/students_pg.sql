CREATE TABLE students (
  id SERIAL PRIMARY KEY,
  first_name VARCHAR(50) NOT NULL,
  middle_name VARCHAR(50),
  last_name VARCHAR(50) NOT NULL
);

INSERT INTO students (first_name, middle_name, last_name) VALUES
('John', 'A', 'Doe'),
('Jane', 'B', 'Smith'),
('Alice', 'C', 'Johnson');