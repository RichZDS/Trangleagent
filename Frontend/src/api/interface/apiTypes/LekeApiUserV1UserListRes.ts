import { type LekeInternalModelUserViewParams } from "../../interface";

export interface LekeApiUserV1UserListRes {
    page?: number;
    pageSize?: number;
    total?: number;
    Users?: LekeInternalModelUserViewParams[];
}
