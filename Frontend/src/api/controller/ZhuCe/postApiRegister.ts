import request from "../../http.js";
import { type TrangleAgentApiLoginV1RegisterRes, type DeepRequired, type TrangleAgentApiLoginV1RegisterReq } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 用户注册
 * /api/register
 */
export function postApiRegister(input?: TrangleAgentApiLoginV1RegisterReq, config?: AxiosRequestConfig) {
    return request.post<DeepRequired<TrangleAgentApiLoginV1RegisterRes>>(`/api/register`, input, config);
}
