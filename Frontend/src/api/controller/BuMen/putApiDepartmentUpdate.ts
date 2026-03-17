import request from "../../http.js";
import { type TrangleAgentApiDepartmentV1DepartmentUpdateRes, type DeepRequired, type TrangleAgentApiDepartmentV1DepartmentUpdateReq } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 更新部门
 * /api/department/update
 */
export function putApiDepartmentUpdate(input?: TrangleAgentApiDepartmentV1DepartmentUpdateReq, config?: AxiosRequestConfig) {
    return request.put<DeepRequired<TrangleAgentApiDepartmentV1DepartmentUpdateRes>>(`/api/department/update`, input, config);
}
