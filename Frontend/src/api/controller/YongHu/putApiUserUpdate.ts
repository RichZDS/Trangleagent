import request from "../../http.js";
import { type TrangleAgentApiUserV1UserUpdateRes, type DeepRequired, type TrangleAgentApiUserV1UserUpdateReq } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 更新用户
 * /api/user/update
 */
export function putApiUserUpdate(input?: TrangleAgentApiUserV1UserUpdateReq, config?: AxiosRequestConfig) {
    return request.put<DeepRequired<TrangleAgentApiUserV1UserUpdateRes>>(`/api/user/update`, input, config);
}
