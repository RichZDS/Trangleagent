import request from "../../http.js";
import { type TrangleAgentApiContainmentV1ContainmentRepoViewRes, type DeepRequired } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 收容库详情
 * /api/containment/view
 */
export function getApiContainmentView(params: GetApiContainmentViewParams, config?: AxiosRequestConfig) {
    const paramsInput = {
        id: params.id,
    };
    return request.get<DeepRequired<TrangleAgentApiContainmentV1ContainmentRepoViewRes>>(`/api/containment/view`, {
        params: paramsInput,
        ...config,
    });
}

export interface GetApiContainmentViewParams {
    /** primary key */
    id?: number;
}
