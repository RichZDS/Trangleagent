import request from "../../http.js";
import { type TrangleAgentApiUserV1RolePermissionCheckRes, type DeepRequired, type TrangleAgentApiUserV1RolePermissionCheckReq } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 检查角色权限
 * /api/role/permission/check
 */
export function postApiRolePermissionCheck(input?: TrangleAgentApiUserV1RolePermissionCheckReq, config?: AxiosRequestConfig) {
    return request.post<DeepRequired<TrangleAgentApiUserV1RolePermissionCheckRes>>(`/api/role/permission/check`, input, config);
}
