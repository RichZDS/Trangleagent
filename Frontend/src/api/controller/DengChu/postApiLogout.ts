import request from "../../http.js";
import { type TrangleAgentApiLoginV1LogoutRes, type DeepRequired, type TrangleAgentApiLoginV1LogoutReq } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 用户登出
 * /api/logout
 */
export function postApiLogout(input?: TrangleAgentApiLoginV1LogoutReq, config?: AxiosRequestConfig) {
    return request.post<DeepRequired<TrangleAgentApiLoginV1LogoutRes>>(`/api/logout`, input, config);
}
