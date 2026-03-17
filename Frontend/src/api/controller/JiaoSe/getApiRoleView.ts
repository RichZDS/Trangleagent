import request from "../../http.js";
import { type TrangleAgentApiUserV1RoleViewRes, type DeepRequired } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 查看角色
 * /api/role/view
 */
export function getApiRoleView(params: GetApiRoleViewParams, config?: AxiosRequestConfig) {
    const paramsInput = {
        id: params.id,
    };
    return request.get<DeepRequired<TrangleAgentApiUserV1RoleViewRes>>(`/api/role/view`, {
        params: paramsInput,
        ...config,
    });
}

export interface GetApiRoleViewParams {
    /** 角色ID */
    id: number;
}
