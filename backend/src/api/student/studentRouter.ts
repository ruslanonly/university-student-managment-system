import express, { type Router } from "express";
import { z } from "zod";

import { GetStudentSchema } from "@/api/student/studentModel";
import { validateRequest } from "@/common/utils/httpHandlers";
import { studentController } from "./studentController";

export const studentRouter: Router = express.Router();

studentRouter.get("/", studentController.getUsers);
studentRouter.get("/:id", validateRequest(GetStudentSchema), studentController.getUser);
