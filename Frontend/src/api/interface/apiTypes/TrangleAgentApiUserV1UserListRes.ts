import { type TrangleAgentInternalModelUserViewParams } from "../../interface";

export interface TrangleAgentApiUserV1UserListRes {
    page?: number;
    pageSize?: number;
    total?: number;
    Users?: TrangleAgentInternalModelUserViewParams[];
}
