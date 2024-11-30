-- Заполнение таблицы units
INSERT INTO units (name) VALUES
('HR Department'),
('IT Department'),
('Sales Department'),
('Marketing Department'),
('Finance Department'),
('Support Team'),
('Development Team'),
('Logistics Team'),
('Management'),
('Administration');

-- Заполнение таблицы units_relations
INSERT INTO units_relations (parent_id, child_id) VALUES
(9, 1), -- HR Department подчиняется Management
(9, 2), -- IT Department подчиняется Management
(9, 3), -- Sales Department подчиняется Management
(9, 4), -- Marketing Department подчиняется Management
(9, 5), -- Finance Department подчиняется Management
(6, 7), -- Support Team подчиняется Development Team
(9, 10), -- Administration подчиняется Management
(8, 6), -- Logistics Team подчиняется Support Team
(3, 8), -- Logistics Team подчиняется Sales Department
(4, 3); -- Marketing Department подчиняется Sales Department

-- Заполнение таблицы employees
INSERT INTO employees (unit_id, name, role_name, family_name, middle_name, phone, city, project, office_address, position, birth_date, is_general) VALUES
(1, 'John', 'role_1', 'Doe', 'Michael', '+1234567890', 'New York', 'Onboarding', '123 Main St', 'HR Specialist', '1985-03-15', TRUE),
(2, 'Alice', 'role_2', 'Smith', 'Mary', '+1234567891', 'San Francisco', 'IT Infrastructure', '456 Elm St', 'System Administrator', '1990-07-22', FALSE),
(2, 'Bob', 'role_3', 'Johnson', 'Andrew', '+1234567892', 'Chicago', 'Regional Sales', '789 Pine St', 'Sales Manager', '1988-11-09', FALSE),
(2, 'Eve', 'role_4', 'Brown', 'Anna', '+1234567893', 'Boston', 'Campaign Launch', '321 Maple St', 'Marketing Coordinator', '1992-05-30', TRUE),
(3, 'Genry', 'role_99', 'Kavil', 'Bob', '+1234567893', 'Boston', 'Campaign Launch', '321 Maple St', 'Marketing Coordinator', '1992-05-30', TRUE),
(4, 'Martin', 'role_4', 'Iden', 'Kayel', '+1234567893', 'Boston', 'Campaign Launch', '321 Maple St', 'Marketing Coordinator', '1992-05-30', TRUE),
(5, 'Charlie', 'role_5', 'Davis', 'Patrick', '+1234567894', 'Seattle', 'Annual Budget', '654 Oak St', 'Finance Analyst', '1986-02-12', TRUE),
(6, 'Grace', 'role_6', 'Wilson', 'Elizabeth', '+1234567895', 'Austin', 'Customer Support', '987 Cedar St', 'Support Specialist', '1994-09-17', TRUE),
(7, 'David', 'role_7', 'Martinez', 'Victor', '+1234567896', 'Denver', 'App Development', '123 Spruce St', 'Software Engineer', '1987-04-25', TRUE),
(8, 'Hannah', 'role_8', 'Taylor', 'Emily', '+1234567897', 'Los Angeles', 'Logistics Optimization', '456 Birch St', 'Logistics Manager', '1991-10-19', TRUE),
(9, 'William', 'role_9', 'Anderson', 'Jacob', '+1234567898', 'Houston', 'Company Strategy', '789 Walnut St', 'CEO', '1975-01-05', TRUE),
(10, 'Emma', 'role_10', 'Thomas', 'Sophia', '+1234567899', 'Dallas', 'Office Operations', '321 Cherry St', 'Administrator', '1989-08-08', TRUE);