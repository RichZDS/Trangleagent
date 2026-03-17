import request from "../../http.js";
import { type TrangleAgentApiContainmentV1ContainmentRepoDeleteRes, type DeepRequired } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 收容库删除
 * /api/containment/delete
 */
export function deleteApiContainmentDelete(params: DeleteApiContainmentDeleteParams, config?: AxiosRequestConfig) {
    const paramsInput = {
        id: params.id,
    };
    return request.delete<DeepRequired<TrangleAgentApiContainmentV1ContainmentRepoDeleteRes>>(`/api/containment/delete`, {
        params: paramsInput,
        ...config,
    });
}

export interface DeleteApiContainmentDeleteParams {
    /** primary key */
    id?: number;
}
