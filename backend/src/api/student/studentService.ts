import { StatusCodes } from "http-status-codes";

import type { Student } from "@/api/student/studentModel";
import { StudentRepository } from "@/api/student/studentRepository";
import { ServiceResponse } from "@/common/models/serviceResponse";
import { logger } from "@/server";

export class StudentService {
  private studentRepository: StudentRepository;

  constructor(repository: StudentRepository = new StudentRepository()) {
    this.studentRepository = repository;
  }

  // Retrieves all users from the database
  async findAll(): Promise<ServiceResponse<Student[] | null>> {
    try {
      const users = await this.studentRepository.findAllAsync();
      if (!users || users.length === 0) {
        return ServiceResponse.failure("No Users found", null, StatusCodes.NOT_FOUND);
      }
      return ServiceResponse.success<Student[]>("Users found", users);
    } catch (ex) {
      const errorMessage = `Error finding all users: $${(ex as Error).message}`;
      logger.error(errorMessage);
      return ServiceResponse.failure(
        "An error occurred while retrieving users.",
        null,
        StatusCodes.INTERNAL_SERVER_ERROR,
      );
    }
  }

  // Retrieves a single user by their ID
  async findById(id: number): Promise<ServiceResponse<Student | null>> {
    try {
      const user = await this.studentRepository.findByIdAsync(id);
      if (!user) {
        return ServiceResponse.failure("Student not found", null, StatusCodes.NOT_FOUND);
      }
      return ServiceResponse.success<Student>("Student found", user);
    } catch (ex) {
      const errorMessage = `Error finding user with id ${id}:, ${(ex as Error).message}`;
      logger.error(errorMessage);
      return ServiceResponse.failure("An error occurred while finding user.", null, StatusCodes.INTERNAL_SERVER_ERROR);
    }
  }
}

export const studentService = new StudentService();
