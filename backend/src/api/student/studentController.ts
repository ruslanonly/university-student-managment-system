import type { Request, RequestHandler, Response } from "express";

import { studentService } from "@/api/student/studentService";
import { handleServiceResponse } from "@/common/utils/httpHandlers";

class StudentController {
  public getUsers: RequestHandler = async (_req: Request, res: Response) => {
    const serviceResponse = await studentService.findAll();
    return handleServiceResponse(serviceResponse, res);
  };

  public getUser: RequestHandler = async (req: Request, res: Response) => {
    const id = Number.parseInt(req.params.id as string, 10);
    const serviceResponse = await studentService.findById(id);
    return handleServiceResponse(serviceResponse, res);
  };
}

export const studentController = new StudentController();
