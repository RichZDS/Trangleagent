import request from "../../http.js";
import { type TrangleAgentApiUserV1RoleUpdateRes, type DeepRequired, type TrangleAgentApiUserV1RoleUpdateReq } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 更新角色
 * /api/role/update
 */
export function putApiRoleUpdate(input?: TrangleAgentApiUserV1RoleUpdateReq, config?: AxiosRequestConfig) {
    return request.put<DeepRequired<TrangleAgentApiUserV1RoleUpdateRes>>(`/api/role/update`, input, config);
}
