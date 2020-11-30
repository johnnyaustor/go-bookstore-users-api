CREATE DATABASE go_bookstore_users;
CREATE TABLE go_bookstore_users.users (
	id INTEGER auto_increment NOT NULL,
	first_name varchar(100) NULL,
	last_name varchar(100) NULL,
	email varchar(100) NOT NULL,
	date_created DATETIME NOT NULL,
	status varchar(100) NOT NULL,
	password varchar(100) NOT NULL,
	CONSTRAINT users_PK PRIMARY KEY (id),
	CONSTRAINT users_UN UNIQUE KEY (email)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;