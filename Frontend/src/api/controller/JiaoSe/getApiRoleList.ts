import request from "../../http.js";
import { type TrangleAgentApiUserV1RoleListRes, type DeepRequired } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 角色列表
 * /api/role/list
 */
export function getApiRoleList(params: GetApiRoleListParams, config?: AxiosRequestConfig) {
    const paramsInput = {
        page: params.page,
        pageSize: params.pageSize,
        total: params.total,
        userId: params.userId,
        departmentId: params.departmentId,
    };
    return request.get<DeepRequired<TrangleAgentApiUserV1RoleListRes>>(`/api/role/list`, {
        params: paramsInput,
        ...config,
    });
}

export interface GetApiRoleListParams {
    page?: number;
    pageSize?: number;
    total?: number;
    /** 用户ID */
    userId?: number;
    /** 部门ID */
    departmentId?: number;
}
