import type { Student } from "@/api/student/studentModel";
import { MongoClient } from "mongodb";

export const students: Student[] = [
    {
        id: 1,
        name: "Alice",
        email: "alice@example.com",
        age: 42,
    },
    {
        id: 2,
        name: "Robert",
        email: "Robert@example.com",
        age: 21,
    },
];

export class StudentRepository {
    async findAllAsync(): Promise<Student[]> {
        const uri = "mongodb://user:password@localhost:27017";
        const client = new MongoClient(uri);

        try {
            await client.connect();
            const db = client.db("my_database");
            const collection = db.collection("users");

            await collection.insertOne({ name: "Alice", age: 30 });

            const user = await collection.findOne({ name: "Alice" });
            console.log("MongoDB - User:", user);

            await collection.updateOne(
                { name: "Alice" },
                { $set: { age: 31 } }
            );

            await collection.deleteOne({ name: "Alice" });
        } catch (err) {
            console.error("MongoDB Error:", err);
        } finally {
            await client.close();
        }
        return students;
    }

    async findByIdAsync(id: number): Promise<Student | null> {
        return students.find((user) => user.id === id) || null;
    }
}
