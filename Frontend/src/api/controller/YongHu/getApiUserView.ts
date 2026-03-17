import request from "../../http.js";
import { type TrangleAgentApiUserV1UserViewRes, type DeepRequired } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 用户详情
 * /api/user/view
 */
export function getApiUserView(params: GetApiUserViewParams, config?: AxiosRequestConfig) {
    const paramsInput = {
        account: params.account,
        nickname: params.nickname,
        id: params.id,
    };
    return request.get<DeepRequired<TrangleAgentApiUserV1UserViewRes>>(`/api/user/view`, {
        params: paramsInput,
        ...config,
    });
}

export interface GetApiUserViewParams {
    /** 账号 */
    account?: string;
    /** 昵称 */
    nickname?: string;
    /** 用户ID */
    id?: number;
}
