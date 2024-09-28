const redis = require('redis');
const client = redis.createClient();

client.on('error', (err) => {
    console.error('Redis error:', err);
});

// Данные студентов
const students = [
    { first_name: 'John', middle_name: 'A', last_name: 'Doe' },
    { first_name: 'Jane', middle_name: 'B', last_name: 'Smith' },
    { first_name: 'Alice', middle_name: 'C', last_name: 'Johnson' }
];

// Добавляем студентов в Redis
students.forEach(student => {
    const key = `student:${student.first_name}:${student.last_name}`;
    client.hmset(key, student, (err, res) => {
        if (err) {
            console.error('Error adding student to Redis:', err);
        } else {
            console.log('Added to Redis:', res);
        }
    });
});

client.quit();
