import { type LekeApiUserV1RoleViewParams } from "../../interface";

export interface LekeApiUserV1RoleListRes {
    page?: number;
    pageSize?: number;
    total?: number;
    /** 角色列表 */
    list?: LekeApiUserV1RoleViewParams[];
}
