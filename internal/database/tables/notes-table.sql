CREATE TABLE notes (
                       id SERIAL PRIMARY KEY,
                       title VARCHAR(100),
                       content TEXT,
                       user_id INT,
                       FOREIGN KEY (user_id) REFERENCES users(id)
);