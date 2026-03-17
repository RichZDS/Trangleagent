import { type TrangleAgentApiUserV1RoleViewParams } from "../../interface";

export interface TrangleAgentApiUserV1RoleListRes {
    page?: number;
    pageSize?: number;
    total?: number;
    /** 角色列表 */
    list?: TrangleAgentApiUserV1RoleViewParams[];
}
