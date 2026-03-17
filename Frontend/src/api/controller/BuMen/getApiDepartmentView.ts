import request from "../../http.js";
import { type TrangleAgentApiDepartmentV1DepartmentViewRes, type DeepRequired } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 部门详情
 * /api/department/view
 */
export function getApiDepartmentView(params: GetApiDepartmentViewParams, config?: AxiosRequestConfig) {
    const paramsInput = {
        id: params.id,
    };
    return request.get<DeepRequired<TrangleAgentApiDepartmentV1DepartmentViewRes>>(`/api/department/view`, {
        params: paramsInput,
        ...config,
    });
}

export interface GetApiDepartmentViewParams {
    id: number;
}
