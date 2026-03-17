import request from "../../http.js";
import { type TrangleAgentApiUserV1UserListRes, type DeepRequired } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 用户列表
 * /api/user/list
 */
export function getApiUserList(params: GetApiUserListParams, config?: AxiosRequestConfig) {
    const paramsInput = {
        page: params.page,
        pageSize: params.pageSize,
        total: params.total,
        account: params.account,
        nickname: params.nickname,
    };
    return request.get<DeepRequired<TrangleAgentApiUserV1UserListRes>>(`/api/user/list`, {
        params: paramsInput,
        ...config,
    });
}

export interface GetApiUserListParams {
    page?: number;
    pageSize?: number;
    total?: number;
    /** 账号 */
    account?: string;
    /** 昵称 */
    nickname?: string;
}
