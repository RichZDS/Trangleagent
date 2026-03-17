import request from "../../http.js";
import { type TrangleAgentApiContainmentV1ContainmentRepoListRes, type DeepRequired } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 收容库列表
 * /api/containment/list
 */
export function getApiContainmentList(params: GetApiContainmentListParams, config?: AxiosRequestConfig) {
    const paramsInput = {
        page: params.page,
        pageSize: params.pageSize,
        total: params.total,
        terminalId: params.terminalId,
        anomalyName: params.anomalyName,
        agentName: params.agentName,
        repoName: params.repoName,
    };
    return request.get<DeepRequired<TrangleAgentApiContainmentV1ContainmentRepoListRes>>(`/api/containment/list`, {
        params: paramsInput,
        ...config,
    });
}

export interface GetApiContainmentListParams {
    page?: number;
    pageSize?: number;
    total?: number;
    /** terminal (散逸端) */
    terminalId?: number;
    /** anomaly name */
    anomalyName?: string;
    /** agent name */
    agentName?: string;
    /** containment repo name */
    repoName?: string;
}
