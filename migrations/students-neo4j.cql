// ============================
// 1. Create Constraints
// ============================

// Specialization
CREATE CONSTRAINT specialization_id IF NOT EXISTS 
FOR (s:Specialization) 
REQUIRE s.id IS UNIQUE;

// Course
CREATE CONSTRAINT course_id IF NOT EXISTS 
FOR (c:Course) 
REQUIRE c.id IS UNIQUE;

// Class
CREATE CONSTRAINT class_id IF NOT EXISTS 
FOR (cl:Class) 
REQUIRE cl.id IS UNIQUE;

// Material
CREATE CONSTRAINT material_id IF NOT EXISTS 
FOR (m:Material) 
REQUIRE m.id IS UNIQUE;

// Schedule
CREATE CONSTRAINT schedule_id IF NOT EXISTS 
FOR (sc:Schedule) 
REQUIRE sc.id IS UNIQUE;

// Student
CREATE CONSTRAINT student_id IF NOT EXISTS 
FOR (st:Student) 
REQUIRE st.id IS UNIQUE;

// Group
CREATE CONSTRAINT group_id IF NOT EXISTS 
FOR (g:Group) 
REQUIRE g.id IS UNIQUE;

// ============================
// 2. Create Nodes
// ============================

// 2.4 Specializations
CREATE (:Specialization {id: 1, name: 'Прикладная математика'});
CREATE (:Specialization {id: 2, name: 'Прикладная математика'});
CREATE (:Specialization {id: 3, name: 'Прикладная информатика'});
CREATE (:Specialization {id: 4, name: 'Информационные системы и технологии'});
CREATE (:Specialization {id: 5, name: 'Программная инженерия'});

// 2.5 Courses
CREATE (:Course {id: 1, name: 'Основы программирования'});
CREATE (:Course {id: 2, name: 'Системы связи'});
CREATE (:Course {id: 3, name: 'Теория автоматического управления'});
CREATE (:Course {id: 4, name: 'Математические модели и методы безопасного функционирования компонент программного обеспечения'});
CREATE (:Course {id: 5, name: 'Методы и средства сборки и интеграции программных модулей, сервисов и компонентов'});
CREATE (:Course {id: 6, name: 'Методы оценки эффективности информационных систем'});
CREATE (:Course {id: 7, name: 'Облачные технологии'});
CREATE (:Course {id: 8, name: 'Проектирование архитектуры программного обеспечения'});
CREATE (:Course {id: 9, name: 'Разработка мобильных компонент анализа безопасности программного обеспечения'});
CREATE (:Course {id: 10, name: 'Разработка программного решения для стартапа'});
CREATE (:Course {id: 11, name: 'Создание инструментальных средств разработки программного обеспечения'});
CREATE (:Course {id: 12, name: 'Управление информационно-технологическими проектами'});
CREATE (:Course {id: 13, name: 'Разработка безопасного программного обеспечения'});

// 2.6 Groups
CREATE (:Group {id: 1, name: 'БСБО-01-21'});
CREATE (:Group {id: 2, name: 'БСБО-02-21'});
CREATE (:Group {id: 3, name: 'БСБО-04-21'});

// 2.7 Students
CREATE (:Student {id: 1, full_name: 'Петров Алексей Сергеевич', group_id: 1});
CREATE (:Student {id: 2, full_name: 'Смирнов Дмитрий Валерьевич', group_id: 1});
CREATE (:Student {id: 3, full_name: 'Кузнецова Мария Александровна', group_id: 1});
CREATE (:Student {id: 4, full_name: 'Васильев Николай Игоревич', group_id: 1});
CREATE (:Student {id: 5, full_name: 'Соколова Ольга Вячеславовна', group_id: 1});
CREATE (:Student {id: 6, full_name: 'Григорьев Андрей Викторович', group_id: 1});
CREATE (:Student {id: 7, full_name: 'Зайцева Екатерина Анатольевна', group_id: 1});
CREATE (:Student {id: 8, full_name: 'Михайлов Артем Сергеевич', group_id: 1});
CREATE (:Student {id: 9, full_name: 'Попова Валерия Евгеньевна', group_id: 1});
CREATE (:Student {id: 10, full_name: 'Кравцов Кирилл Александрович', group_id: 1});
CREATE (:Student {id: 11, full_name: 'Иванов Иван Иванович', group_id: 1});
CREATE (:Student {id: 12, full_name: 'Фадеев Всеволод Вадимович', group_id: 2});
CREATE (:Student {id: 13, full_name: 'Усанкин Александр Александрович', group_id: 2});
CREATE (:Student {id: 14, full_name: 'Рзаев Руслан Халидович', group_id: 2});
CREATE (:Student {id: 15, full_name: 'Тимофеев Максим Павлович', group_id: 3});
CREATE (:Student {id: 16, full_name: 'Лебедев Артем Юрьевич', group_id: 3});
CREATE (:Student {id: 17, full_name: 'Фролова Дарина Васильевна', group_id: 3});
CREATE (:Student {id: 18, full_name: 'Ильин Иван Игоревич', group_id: 3});
CREATE (:Student {id: 19, full_name: 'Чернова Светлана Павловна', group_id: 3});
CREATE (:Student {id: 20, full_name: 'Андреев Павел Владиславович', group_id: 3});
CREATE (:Student {id: 21, full_name: 'Егорова Виктория Николаевна', group_id: 3});
CREATE (:Student {id: 22, full_name: 'Морозов Михаил Евгеньевич', group_id: 3});
CREATE (:Student {id: 23, full_name: 'Сидорова Ирина Борисовна', group_id: 3});
CREATE (:Student {id: 24, full_name: 'Борисов Сергей Юрьевич', group_id: 3});

// 2.8 Classes
CREATE (:Class {id: 1, course_id: 1, type: 'Лекторная'});
CREATE (:Class {id: 2, course_id: 2, type: 'Лабораторная'});
CREATE (:Class {id: 3, course_id: 3, type: 'Семинарная'});

// 2.9 Materials
CREATE (:Material {id: 1, content: 'Теоретические основы программирования', class_id: 1});
CREATE (:Material {id: 2, content: 'Лабораторная работа по радиотехнике', class_id: 2});
CREATE (:Material {id: 3, content: 'Практическое занятие по автоматике', class_id: 3});

// 2.10 Schedules
CREATE (:Schedule {id: 1, date: datetime('2024-12-06T09:00:00'), class_id: 1, group_id: 1});
CREATE (:Schedule {id: 2, date: datetime('2024-12-06T11:00:00'), class_id: 2, group_id: 2});
CREATE (:Schedule {id: 3, date: datetime('2024-12-07T14:00:00'), class_id: 3, group_id: 2});
CREATE (:Schedule {id: 4, date: datetime('2024-12-07T16:00:00'), class_id: 1, group_id: 3});
CREATE (:Schedule {id: 5, date: datetime('2024-12-08T10:00:00'), class_id: 2, group_id: 1});
CREATE (:Schedule {id: 6, date: datetime('2024-12-08T12:00:00'), class_id: 3, group_id: 3});
CREATE (:Schedule {id: 7, date: datetime('2024-12-09T09:00:00'), class_id: 1, group_id: 2});
CREATE (:Schedule {id: 8, date: datetime('2024-12-09T11:00:00'), class_id: 2, group_id: 1});
CREATE (:Schedule {id: 9, date: datetime('2024-12-10T13:00:00'), class_id: 3, group_id: 3});
CREATE (:Schedule {id: 10, date: datetime('2024-12-10T15:00:00'), class_id: 1, group_id: 2});

// ============================
// 3. Create Relationships
// ============================

// 3.4 Specialization -> Course (Specialization_Course)
MATCH (s:Specialization), (c:Course)
WHERE 
    (s.id = 1 AND c.id = 1) OR
    (s.id = 2 AND c.id = 2) OR
    (s.id = 3 AND c.id = 3) OR
    (s.id = 4 AND c.id IN [4,5,6,7,8,9,10,11,12,13])
CREATE (s)-[:INCLUDES_COURSE]->(c);

// 3.5 Course -> Class
MATCH (c:Course), (cl:Class)
WHERE cl.course_id = c.id
CREATE (c)-[:HAS_CLASS]->(cl);

// 3.6 Class -> Material
MATCH (cl:Class), (m:Material)
WHERE m.class_id = cl.id
CREATE (cl)-[:HAS_MATERIAL]->(m);

// 3.7 Class -> Schedule
MATCH (cl:Class), (sc:Schedule)
WHERE sc.class_id = cl.id
CREATE (cl)-[:SCHEDULED_AT]->(sc);

// 3.8 Schedule -> Group
MATCH (sc:Schedule), (g:Group)
WHERE sc.group_id = g.id
CREATE (sc)-[:FOR_GROUP]->(g);

// 3.9 Group -> Student
MATCH (g:Group), (st:Student)
WHERE st.group_id = g.id
CREATE (g)-[:HAS_STUDENT]->(st);

// 3.10 Student -> Group
MATCH (st:Student), (g:Group)
WHERE st.group_id = g.id
CREATE (st)-[:BELONGS_TO]->(g);

// 3.11 Create ATTENDS Relationships Based on Attendance Status '+'
UNWIND [
    {status: '+', date: datetime('2024-12-06T09:00:00'), student_id: 1, schedule_id: 1},
    {status: '+', date: datetime('2024-12-06T11:00:00'), student_id: 4, schedule_id: 2},
    {status: '+', date: datetime('2024-12-07T14:00:00'), student_id: 7, schedule_id: 3},
    {status: '+', date: datetime('2024-12-07T16:00:00'), student_id: 10, schedule_id: 4},
    {status: '+', date: datetime('2024-12-08T12:00:00'), student_id: 13, schedule_id: 6},
    {status: '+', date: datetime('2024-12-09T09:00:00'), student_id: 16, schedule_id: 7},
    {status: '+', date: datetime('2024-12-10T13:00:00'), student_id: 19, schedule_id: 9},
    {status: '+', date: datetime('2024-12-10T15:00:00'), student_id: 21, schedule_id: 10}
] AS attendance
MATCH (st:Student {id: attendance.student_id}), (sc:Schedule {id: attendance.schedule_id})
CREATE (st)-[:ATTENDS {date: attendance.date}]->(sc);