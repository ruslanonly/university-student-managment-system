import type {
    IStudent,
    IStudentRepository,
    TStudentCreateRequest,
    TStudentUpdateRequest,
} from "@/api/student/student-model";
import { Pool } from "pg";
import { createClient as createRedisClient } from "redis";
import { MongoClient } from "mongodb";
import { Client as ElasticSearchClient } from "@elastic/elasticsearch";
import neo4j from "neo4j-driver";

// PostgreSQL initialization
const pgPool = new Pool({
    user: "user",
    host: "localhost",
    database: "my_database",
    password: "password",
    port: 5432, // стандартный порт PostgreSQL
});

pgPool.on("error", (err: Error) => {
    console.error("Unexpected error on idle PostgreSQL client", err);
    process.exit(-1);
});

// Redis initialization
const redisClient = createRedisClient({
    url: "redis://localhost:6379",
});

redisClient.on("error", (err: Error) => {
    console.error("Redis error", err);
});

const mongoClient = new MongoClient('mongodb://user:password@localhost:27017');

mongoClient.connect().then((err) => {
    if (err) {
        console.error("MongoDB connection error:", err);
    } else {
        console.log("Connected to MongoDB");
    }
});

// Elasticsearch initialization
const elasticClient = new ElasticSearchClient({
    node: "http://localhost:9200",
});

elasticClient.ping({}).then((err) => {
    if (err) {
        console.error("Elasticsearch cluster is down!", err);
    } else {
        console.log("Connected to Elasticsearch");
    }
});

// Neo4j initialization
const neo4jDriver = neo4j.driver(
    "bolt://localhost:7687",
    neo4j.auth.basic("neo4j", "password")
);

neo4jDriver
    .verifyConnectivity()
    .then(() => {
        console.log("Connected to Neo4j");
    })
    .catch((err) => {
        console.error("Neo4j connection error", err);
    });

export class PostgresStudentRepository implements IStudentRepository {
    async create(data: TStudentCreateRequest): Promise<void> {
        const { first_name, middle_name, last_name } = data;
        const client = await pgPool.connect();

        try {
            await client.query(
                `INSERT INTO students (first_name, middle_name, last_name) VALUES ($1, $2, $3)`,
                [first_name, middle_name, last_name]
            );
        } catch (err) {
            console.error("Error creating student:", err);
            throw err;
        } finally {
            client.release();
        }
    }

    async findById(id: number): Promise<any> {
        const client = await pgPool.connect();

        try {
            const result = await client.query(
                `SELECT * FROM students WHERE id = $1`,
                [id]
            );

            return result.rows[0] || null;
        } catch (err) {
            console.error("Error finding student by ID:", err);
            throw err;
        } finally {
            client.release();
        }
    }

    async update(id: number, data: TStudentUpdateRequest): Promise<void> {
        const { first_name, middle_name, last_name } = data;
        const client = await pgPool.connect();

        try {
            await client.query(
                `UPDATE students SET first_name = $1, middle_name = $2, last_name = $3 WHERE id = $4`,
                [first_name, middle_name, last_name, id]
            );
        } catch (err) {
            console.error("Error updating student:", err);
            throw err;
        } finally {
            client.release();
        }
    }

    async delete(id: number): Promise<void> {
        const client = await pgPool.connect();

        try {
            await client.query(`DELETE FROM students WHERE id = $1`, [id]);
        } catch (err) {
            console.error("Error deleting student:", err);
            throw err;
        } finally {
            client.release();
        }
    }
}

export class MongoStudentRepository implements IStudentRepository {
    private db = mongoClient.db("university").collection("students");

    async create(data: TStudentCreateRequest): Promise<void> {
        await this.db.insertOne(data);
    }

    async findById(id: number): Promise<any> {
        return this.db.findOne({ id });
    }

    async update(id: number, data: TStudentUpdateRequest): Promise<void> {
        await this.db.updateOne({ id }, { $set: data });
    }

    async delete(id: number): Promise<void> {
        await this.db.deleteOne({ id });
    }
}

export class RedisStudentRepository implements IStudentRepository {
    async create(data: TStudentCreateRequest): Promise<void> {
        await redisClient.set(`student:${5}`, JSON.stringify(data));
    }

    async findById(id: number): Promise<any> {
        const result = await redisClient.get(`student:${id}`);
        return result ? JSON.parse(result) : null;
    }

    async update(id: number, data: TStudentUpdateRequest): Promise<void> {
        await redisClient.set(`student:${id}`, JSON.stringify(data));
    }

    async delete(id: number): Promise<void> {
        await redisClient.del(`student:${id}`);
    }
}

export class ElasticStudentRepository implements IStudentRepository {
    async create(data: TStudentCreateRequest): Promise<void> {
        await elasticClient.index({
            index: "students",
            id: String(5),
            document: data,
        });
    }

    async findById(id: number): Promise<any> {
        return elasticClient.get({ index: "students", id: String(id) });
    }

    async update(id: number, data: TStudentUpdateRequest): Promise<void> {
        await elasticClient.update({
            index: "students",
            id: String(id),
            doc: data,
        });
    }

    async delete(id: number): Promise<void> {
        await elasticClient.delete({ index: "students", id: String(id) });
    }
}

export class Neo4jStudentRepository implements IStudentRepository {
    private session = neo4jDriver.session();

    async create(data: TStudentCreateRequest): Promise<void> {
        await this.session.run(
            `CREATE (s:Student {id: $id, first_name: $first_name, middle_name: $middle_name, last_name: $last_name})`,
            data
        );
    }

    async findById(id: number): Promise<any> {
        const result = await this.session.run(
            `MATCH (s:Student {id: $id}) RETURN s`,
            { id }
        );
        return result.records[0]?.get("s") || null;
    }

    async update(id: number, data: TStudentUpdateRequest): Promise<void> {
        await this.session.run(
            `MATCH (s:Student {id: $id}) SET s.first_name = $first_name, s.middle_name = $middle_name, s.last_name = $last_name`,
            { id, ...data }
        );
    }

    async delete(id: number): Promise<void> {
        await this.session.run(`MATCH (s:Student {id: $id}) DELETE s`, { id });
    }
}
