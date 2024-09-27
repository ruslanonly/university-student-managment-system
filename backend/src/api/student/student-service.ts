import type {
    IStudent,
    TStudentCreateRequest,
    TStudentUpdateRequest,
} from "@/api/student/student-model";
import {
    ElasticStudentRepository,
    MongoStudentRepository,
    Neo4jStudentRepository,
    PostgresStudentRepository,
    RedisStudentRepository,
} from "@/api/student/student-repository";
import { ServiceResponse } from "@/common/models/serviceResponse";

class StudentService {
    private repositories = [
        new PostgresStudentRepository(),
        new MongoStudentRepository(),
        new RedisStudentRepository(),
        new ElasticStudentRepository(),
        new Neo4jStudentRepository(),
    ];

    public async create(
        data: TStudentCreateRequest
    ): Promise<ServiceResponse<void>> {
        try {
            await Promise.all(
                this.repositories.map((repo) => repo.create(data))
            );

            return ServiceResponse.success<void>(
                "Студент успешно создан",
                undefined,
                200
            );
        } catch (error) {
            return ServiceResponse.success<void>(
                "Не удалось создать студента",
                undefined,
                200
            );
        }
    }

    public async findById(id: number): Promise<ServiceResponse<any[]>> {
        try {
            const results = await Promise.all(
                this.repositories.map((repo) => repo.findById(id))
            );

            return ServiceResponse.success(
                "Не удалось создать студента",
                results.find((result) => result !== null),
                200
            );
        } catch (error) {
            return ServiceResponse.failure(
                "Не удалось создать студента",
                [],
                400
            );
        }
    }

    public async update(
        id: number,
        data: TStudentUpdateRequest
    ): Promise<ServiceResponse<boolean>> {
        try {
            await Promise.all(
                this.repositories.map((repo) => repo.update(id, data))
            );

            return ServiceResponse.success(
                "Студент успешно изменен",
                true,
                200
            );
        } catch (error) {
            return ServiceResponse.failure(
                "Не удалось измененить студента",
                true,
                400
            );
        }
    }

    public async delete(id: number): Promise<ServiceResponse<boolean>> {
        try {
          await Promise.all(this.repositories.map((repo) => repo.delete(id)));

          return ServiceResponse.success(
              "Студент успешно удален",
              true,
              200
          );
      } catch (error) {
          return ServiceResponse.failure(
              "Не удалось удалить студента",
              false,
              400
          );
      }
    }
}

export const studentService = new StudentService();
