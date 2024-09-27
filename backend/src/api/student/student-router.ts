import express, { type Router } from "express";

import { studentController } from "./student-сontroller";

export const studentRouter: Router = express.Router();

studentRouter.get("/", studentController.getStudent);
studentRouter.get("/:id", studentController.getStudent);
