import request from "../../http.js";
import { type TrangleAgentApiDepartmentV1DepartmentDeleteRes, type DeepRequired } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 删除部门
 * /api/department/delete
 */
export function deleteApiDepartmentDelete(params: DeleteApiDepartmentDeleteParams, config?: AxiosRequestConfig) {
    const paramsInput = {
        id: params.id,
    };
    return request.delete<DeepRequired<TrangleAgentApiDepartmentV1DepartmentDeleteRes>>(`/api/department/delete`, {
        params: paramsInput,
        ...config,
    });
}

export interface DeleteApiDepartmentDeleteParams {
    id: number;
}
