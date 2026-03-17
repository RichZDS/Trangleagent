import request from "../../http.js";
import { type TrangleAgentApiUserV1RoleCreateRes, type DeepRequired, type TrangleAgentApiUserV1RoleCreateReq } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 创建角色
 * /api/role/create
 */
export function postApiRoleCreate(input?: TrangleAgentApiUserV1RoleCreateReq, config?: AxiosRequestConfig) {
    return request.post<DeepRequired<TrangleAgentApiUserV1RoleCreateRes>>(`/api/role/create`, input, config);
}
