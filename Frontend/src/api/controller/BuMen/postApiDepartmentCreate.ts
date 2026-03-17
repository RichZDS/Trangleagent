import request from "../../http.js";
import { type TrangleAgentApiDepartmentV1DepartmentCreateRes, type DeepRequired, type TrangleAgentApiDepartmentV1DepartmentCreateReq } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 创建部门
 * /api/department/create
 */
export function postApiDepartmentCreate(input?: TrangleAgentApiDepartmentV1DepartmentCreateReq, config?: AxiosRequestConfig) {
    return request.post<DeepRequired<TrangleAgentApiDepartmentV1DepartmentCreateRes>>(`/api/department/create`, input, config);
}
