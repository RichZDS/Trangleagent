import request from "../../http.js";
import { type TrangleAgentApiLoginV1SendVerificationCodeRes, type DeepRequired, type TrangleAgentApiLoginV1SendVerificationCodeReq } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 发送验证码
 * /api/email/send-code
 */
export function postApiEmailSendCode(input?: TrangleAgentApiLoginV1SendVerificationCodeReq, config?: AxiosRequestConfig) {
    return request.post<DeepRequired<TrangleAgentApiLoginV1SendVerificationCodeRes>>(`/api/email/send-code`, input, config);
}
