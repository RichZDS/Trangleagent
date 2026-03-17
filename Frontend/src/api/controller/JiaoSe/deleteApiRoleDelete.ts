import request from "../../http.js";
import { type TrangleAgentApiUserV1RoleDeleteRes, type DeepRequired } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 删除角色
 * /api/role/delete
 */
export function deleteApiRoleDelete(params: DeleteApiRoleDeleteParams, config?: AxiosRequestConfig) {
    const paramsInput = {
        id: params.id,
    };
    return request.delete<DeepRequired<TrangleAgentApiUserV1RoleDeleteRes>>(`/api/role/delete`, {
        params: paramsInput,
        ...config,
    });
}

export interface DeleteApiRoleDeleteParams {
    /** 角色ID */
    id: number;
}
