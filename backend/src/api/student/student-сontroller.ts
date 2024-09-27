import type { Request, RequestHandler, Response } from "express";

import { studentService } from "@/api/student/student-service";
import { handleServiceResponse } from "@/common/utils/httpHandlers";

class StudentController {
    // Получение студента по ID
    public getStudent: RequestHandler = async (req: Request, res: Response) => {
        try {
            const id = Number.parseInt(req.params.id as string, 10);
            if (isNaN(id)) {
                return res.status(400).json({ message: "Invalid student ID" });
            }

            const serviceResponse = await studentService.findById(id);
            return handleServiceResponse(serviceResponse, res);
        } catch (error) {
            return res.status(500).json({ message: "Internal Server Error" });
        }
    };

    // Создание нового студента
    public createStudent: RequestHandler = async (req: Request, res: Response) => {
        try {
            const { first_name, middle_name, last_name } = req.body;

            if (!first_name || !last_name) {
                return res
                    .status(400)
                    .json({ message: "First name and Last name are required" });
            }

            const serviceResponse = await studentService.create({
                first_name,
                middle_name,
                last_name,
            });

            return handleServiceResponse(serviceResponse, res);
        } catch (error) {
            return res.status(500).json({ message: "Internal Server Error" });
        }
    };

    // Обновление данных студента по ID
    public updateStudent: RequestHandler = async (req: Request, res: Response) => {
        try {
            const id = Number.parseInt(req.params.id as string, 10);
            if (isNaN(id)) {
                return res.status(400).json({ message: "Invalid student ID" });
            }

            const { first_name, middle_name, last_name } = req.body;

            if (!first_name || !last_name) {
                return res
                    .status(400)
                    .json({ message: "First name and Last name are required" });
            }

            const serviceResponse = await studentService.update(id, {
                first_name,
                middle_name,
                last_name,
            });

            return handleServiceResponse(serviceResponse, res);
        } catch (error) {
            return res.status(500).json({ message: "Internal Server Error" });
        }
    };

    // Удаление студента по ID
    public deleteStudent: RequestHandler = async (req: Request, res: Response) => {
        try {
            const id = Number.parseInt(req.params.id as string, 10);
            if (isNaN(id)) {
                return res.status(400).json({ message: "Invalid student ID" });
            }

            const serviceResponse = await studentService.delete(id);
            return handleServiceResponse(serviceResponse, res);
        } catch (error) {
            return res.status(500).json({ message: "Internal Server Error" });
        }
    };
}

export const studentController = new StudentController();
