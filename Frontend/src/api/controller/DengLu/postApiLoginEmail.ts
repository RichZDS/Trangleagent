import request from "../../http.js";
import { type TrangleAgentApiLoginV1LoginByEmailRes, type DeepRequired, type TrangleAgentApiLoginV1LoginByEmailReq } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 用户邮箱登录
 * /api/login/email
 */
export function postApiLoginEmail(input?: TrangleAgentApiLoginV1LoginByEmailReq, config?: AxiosRequestConfig) {
    return request.post<DeepRequired<TrangleAgentApiLoginV1LoginByEmailRes>>(`/api/login/email`, input, config);
}
