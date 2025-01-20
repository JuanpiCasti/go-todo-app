CREATE SCHEMA IF NOT EXISTS todo_app;

-- Users, Roles and Permissions

CREATE TABLE IF NOT EXISTS todo_app.app_users (
  id SERIAL,
  username VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  role_id SMALLINT NOT NULL,
  active BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS todo_app.app_roles (
  id smallint NOT NULL,
  name VARCHAR(255) NOT NULL
);

-- Domain

CREATE TABLE IF NOT EXISTS todo_app.todos (
  id SERIAL,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  completed BOOLEAN DEFAULT FALSE,
  user_id INTEGER NOT NULL
);

-- Primary key constraints

ALTER TABLE todo_app.app_users ADD PRIMARY KEY (id);

ALTER TABLE todo_app.app_roles ADD PRIMARY KEY (id);

ALTER TABLE todo_app.todos ADD PRIMARY KEY (id);

-- Foreign key constraints

ALTER TABLE todo_app.app_users ADD FOREIGN KEY (role_id) REFERENCES todo_app.app_roles(id);

Alter TABLE todo_app.todos ADD FOREIGN KEY (user_id) REFERENCES todo_app.app_users(id);

