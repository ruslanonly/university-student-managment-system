import { z } from "zod";

import { commonValidations } from "@/common/utils/commonValidation";

export type Student = z.infer<typeof StudentSchema>;
export const StudentSchema = z.object({
  id: z.number(),
  name: z.string(),
  email: z.string().email(),
  age: z.number(),
});

// Input Validation for 'GET students/:id' endpoint
export const GetStudentSchema = z.object({
  params: z.object({ id: commonValidations.id }),
});
