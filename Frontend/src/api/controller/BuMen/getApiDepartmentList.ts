import request from "../../http.js";
import { type TrangleAgentApiDepartmentV1DepartmentListRes, type DeepRequired } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 部门列表
 * /api/department/list
 */
export function getApiDepartmentList(params: GetApiDepartmentListParams, config?: AxiosRequestConfig) {
    const paramsInput = {
        page: params.page,
        pageSize: params.pageSize,
        total: params.total,
        branchName: params.branchName,
        managerName: params.managerName,
        userId: params.userId,
    };
    return request.get<DeepRequired<TrangleAgentApiDepartmentV1DepartmentListRes>>(`/api/department/list`, {
        params: paramsInput,
        ...config,
    });
}

export interface GetApiDepartmentListParams {
    page?: number;
    pageSize?: number;
    total?: number;
    /** 分部名称 */
    branchName?: string;
    /** 分部经理名称 */
    managerName?: string;
    /** 所属用户ID */
    userId?: number;
}
