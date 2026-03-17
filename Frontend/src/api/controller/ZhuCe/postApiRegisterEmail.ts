import request from "../../http.js";
import { type TrangleAgentApiLoginV1RegisterByEmailRes, type DeepRequired, type TrangleAgentApiLoginV1RegisterByEmailReq } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 用户邮箱注册
 * /api/register/email
 */
export function postApiRegisterEmail(input?: TrangleAgentApiLoginV1RegisterByEmailReq, config?: AxiosRequestConfig) {
    return request.post<DeepRequired<TrangleAgentApiLoginV1RegisterByEmailRes>>(`/api/register/email`, input, config);
}
