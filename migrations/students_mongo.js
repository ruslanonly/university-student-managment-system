db = db.getSiblingDB('my_database');

db.students.insertMany([
    { first_name: 'John', middle_name: 'A', last_name: 'Doe' },
    { first_name: 'Jane', middle_name: 'B', last_name: 'Smith' },
    { first_name: 'Alice', middle_name: 'C', last_name: 'Johnson' }
]);
