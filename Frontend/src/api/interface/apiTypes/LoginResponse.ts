import { type User } from "../../interface";

export interface LoginResponse {
    token?: string;
    user?: User;
}
