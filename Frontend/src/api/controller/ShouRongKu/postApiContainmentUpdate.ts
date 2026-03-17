import request from "../../http.js";
import { type TrangleAgentApiContainmentV1ContainmentRepoUpdateRes, type DeepRequired, type TrangleAgentApiContainmentV1ContainmentRepoUpdateReq } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 创建或更新收容库
 * /api/containment/update
 */
export function postApiContainmentUpdate(input?: TrangleAgentApiContainmentV1ContainmentRepoUpdateReq, config?: AxiosRequestConfig) {
    return request.post<DeepRequired<TrangleAgentApiContainmentV1ContainmentRepoUpdateRes>>(`/api/containment/update`, input, config);
}
