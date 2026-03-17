import request from "../../http.js";
import { type TrangleAgentApiUserV1UserDeleteRes, type DeepRequired } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 删除用户
 * /api/user/delete
 */
export function deleteApiUserDelete(params: DeleteApiUserDeleteParams, config?: AxiosRequestConfig) {
    const paramsInput = {
        account: params.account,
    };
    return request.delete<DeepRequired<TrangleAgentApiUserV1UserDeleteRes>>(`/api/user/delete`, {
        params: paramsInput,
        ...config,
    });
}

export interface DeleteApiUserDeleteParams {
    /** 账号 */
    account?: string;
}
