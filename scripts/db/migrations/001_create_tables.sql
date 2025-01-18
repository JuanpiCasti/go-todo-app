CREATE DATABASE todo_app;

-- Users, Roles and Permissions

CREATE TABLE app_users (
  id SERIAL,
  username VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  role_id SMALLINT NOT NULL,
  active BOOLEAN NOT NULL
);

CREATE TABLE app_roles (
  id smallint NOT NULL,
  name VARCHAR(255) NOT NULL
);

-- Domain

CREATE TABLE todos (
  id SERIAL,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  completed BOOLEAN DEFAULT FALSE,
  user_id INTEGER NOT NULL
);

-- Primary key constraints

ALTER TABLE app_users ADD PRIMARY KEY (id);

ALTER TABLE app_roles ADD PRIMARY KEY (id);

ALTER TABLE todos ADD PRIMARY KEY (id);

-- Foreign key constraints

ALTER TABLE app_users ADD FOREIGN KEY (role_id) REFERENCES app_roles(id);

Alter TABLE todos ADD FOREIGN KEY (user_id) REFERENCES app_users(id);

