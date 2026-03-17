import request from "../../http.js";
import { type TrangleAgentApiLoginV1LoginRes, type DeepRequired, type TrangleAgentApiLoginV1LoginReq } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 用户登录
 * /api/login
 */
export function postApiLogin(input?: TrangleAgentApiLoginV1LoginReq, config?: AxiosRequestConfig) {
    return request.post<DeepRequired<TrangleAgentApiLoginV1LoginRes>>(`/api/login`, input, config);
}
