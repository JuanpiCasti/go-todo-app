INSERT INTO app_roles (id, name) 
VALUES (1, 'admin'), (2, 'user');

INSERT INTO app_users (username, password, role_id, active)
VALUES ('admin', 'admin', 1, true), ('user', 'user', 2, true);
