CREATE TABLE universities (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE institutes (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    university_id INT NOT NULL REFERENCES universities(id)
);

CREATE TABLE departments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    institute_id INT NOT NULL REFERENCES institutes(id)
);

CREATE TABLE specializations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE department_specialization (
    department_id INT NOT NULL REFERENCES departments(id),
    specialization_id INT NOT NULL REFERENCES specializations(id),
    PRIMARY KEY (department_id, specialization_id)
);

CREATE TABLE courses (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE specialization_course (
    specialization_id INT NOT NULL REFERENCES specializations(id),
    course_id INT NOT NULL REFERENCES courses(id),
    PRIMARY KEY (specialization_id, course_id)
);

CREATE TABLE "groups" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE students (
    grade_book_id SERIAL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    group_id INT NOT NULL REFERENCES "groups"(id)
);

CREATE TABLE classes (
    id SERIAL PRIMARY KEY,
    course_id INT NOT NULL REFERENCES courses(id),
    type VARCHAR(50) NOT NULL
);

CREATE TABLE materials (
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    class_id INT NOT NULL REFERENCES classes(id)
);

CREATE TABLE schedules (
    id SERIAL PRIMARY KEY,
    date TIMESTAMP NOT NULL,
    class_id INT NOT NULL REFERENCES classes(id),
    group_id INT NOT NULL REFERENCES "groups"(id)
);

CREATE TABLE attendances (
    id SERIAL PRIMARY KEY,
    status VARCHAR(50) NOT NULL,
    date TIMESTAMP NOT NULL,
    student_id INT NOT NULL REFERENCES students(grade_book_id),
    schedule_id INT NOT NULL REFERENCES schedules(id)
);

INSERT INTO universities (name)
VALUES ('Московский институт радиотехники, электроники и автоматики (МИРЭА)');

INSERT INTO institutes (name, university_id)
VALUES 
('Институт информационных технологий', 1),
('Институт искусственного интеллекта', 1),
('Институт радиоэлектроники и информатики', 1),
('Институт кибербезопасноти и цифровых технологий', 1),
('Институт технологий управления', 1),
('Институт тонких химических технологий им. М.В. Ломоносова', 1),
('Институт перспективных технологий и индустриального программирования', 1);

INSERT INTO departments (name, institute_id)
VALUES
('Кафедра программного обеспечения', 1),
('Кафедра радиотехники', 2),
('Кафедра автоматического управления', 3),
('Кафедра КБ-3 «Разработка программных решений и системного программирования» ', 4);


INSERT INTO specializations (name)
VALUES
('Прикладная математика'),
('Прикладная математика'),
('Прикладная информатика'),
('Информационные системы и технологии'),
('Программная инженерия');

INSERT INTO department_specialization(department_id, specialization_id)
VALUES
(1, 1),
(2, 2),
(3, 3),
(4, 4);

INSERT INTO courses (name)
VALUES
('Основы программирования'),
('Системы связи'),
('Теория автоматического управления'),
('Математические модели и методы безопасного функционирования компонент программного обеспечения'),
('Методы и средства сборки и интеграции программных модулей, сервисов и компонентов'),
('Методы оценки эффективности информационных систем'),
('Облачные технологии'),
('Проектирование архитектуры программного обеспечения'),
('Разработка мобильных компонент анализа безопасности программного обеспечения'),
('Разработка программного решения для стартапа'),
('Создание инструментальных средств разработки программного обеспечения'),
('Управление информационно-технологическими проектами'),
('Разработка безопасного программного обеспечения');

INSERT INTO specialization_course (specialization_id, course_id)
VALUES
(1, 1),
(2, 2),
(3, 3),
(4, 4),
(4, 5),
(4, 6),
(4, 7),
(4, 8),
(4, 9),
(4, 10),
(4, 11),
(4, 12),
(4, 13);

INSERT INTO "groups" (name)
VALUES
('БСБО-01-21'),
('БСБО-02-21'),
('БСБО-04-21');

INSERT INTO students(full_name, group_id)
VALUES
('Петров Алексей Сергеевич', 1),
('Смирнов Дмитрий Валерьевич', 1),
('Кузнецова Мария Александровна', 1),
('Васильев Николай Игоревич', 1),
('Соколова Ольга Вячеславовна', 1),
('Григорьев Андрей Викторович', 1),
('Зайцева Екатерина Анатольевна', 1),
('Михайлов Артем Сергеевич', 1),
('Попова Валерия Евгеньевна', 1),
('Кравцов Кирилл Александрович', 1),
('Иванов Иван Иванович', 1),
('Фадеев Всеволод Вадимович', 2),
('Усанкин Александр Александрович', 2),
('Рзаев Руслан Халидович', 2),
('Тимофеев Максим Павлович', 3),
('Лебедев Артем Юрьевич', 3),
('Фролова Дарина Васильевна', 3),
('Ильин Иван Игоревич', 3),
('Чернова Светлана Павловна', 3),
('Андреев Павел Владиславович', 3),
('Егорова Виктория Николаевна', 3),
('Морозов Михаил Евгеньевич', 3),
('Сидорова Ирина Борисовна', 3),
('Борисов Сергей Юрьевич', 3);

INSERT INTO classes (course_id, type)
VALUES
(1, 'Лекторная'),
(2, 'Лабораторная'),
(3, 'Семинарная');

INSERT INTO materials (content, class_id)
VALUES
('Теоретические основы программирования', 1),
('Лабораторная работа по радиотехнике', 2),
('Практическое занятие по автоматике', 3);

INSERT INTO schedules (date, class_id, group_id)
VALUES
('2024-12-06 09:00:00', 1, 1),
('2024-12-06 11:00:00', 2, 2),
('2024-12-07 14:00:00', 3, 2),
('2024-12-07 16:00:00', 1, 3),
('2024-12-08 10:00:00', 2, 1),
('2024-12-08 12:00:00', 3, 3),
('2024-12-09 09:00:00', 1, 2),
('2024-12-09 11:00:00', 2, 1),
('2024-12-10 13:00:00', 3, 3),
('2024-12-10 15:00:00', 1, 2);

INSERT INTO attendances (status, date, student_id, schedule_id)
VALUES
('+', '2024-12-06 09:00:00', 1, 1),
('У', '2024-12-06 09:00:00', 2, 1),
('Н', '2024-12-06 09:00:00', 3, 1),
('+', '2024-12-06 11:00:00', 4, 2),
('У', '2024-12-06 11:00:00', 5, 2),
('Н', '2024-12-06 11:00:00', 6, 2),
('+', '2024-12-07 14:00:00', 7, 3),
('У', '2024-12-07 14:00:00', 8, 3),
('Н', '2024-12-07 14:00:00', 9, 3),
('+', '2024-12-07 16:00:00', 10, 4),
('У', '2024-12-07 16:00:00', 11, 4),
('Н', '2024-12-08 10:00:00', 12, 5),
('+', '2024-12-08 12:00:00', 13, 6),
('У', '2024-12-08 12:00:00', 14, 6),
('Н', '2024-12-09 09:00:00', 15, 7),
('+', '2024-12-09 09:00:00', 16, 7),
('У', '2024-12-09 11:00:00', 17, 8),
('Н', '2024-12-09 11:00:00', 18, 8),
('+', '2024-12-10 13:00:00', 19, 9),
('У', '2024-12-10 13:00:00', 20, 9),
('+', '2024-12-10 15:00:00', 21, 10);