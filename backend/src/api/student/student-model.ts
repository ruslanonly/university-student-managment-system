export interface IStudent {
    id: number;
    first_name: string;
    middle_name?: string;
    last_name: string;
}

export type TStudentUpdateRequest = Omit<IStudent, "id">;
export type TStudentCreateRequest = Omit<IStudent, "id">;

export interface IStudentRepository {
    create(data: TStudentCreateRequest): Promise<void>;
    findById(id: number): Promise<IStudent | null>;
    update(id: number, data: TStudentUpdateRequest): Promise<void>;
    delete(id: number): Promise<void>;
}
