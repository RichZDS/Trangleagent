import { type TrangleAgentApiDepartmentV1Department } from "../../interface";

export interface TrangleAgentApiDepartmentV1DepartmentListRes {
    page?: number;
    pageSize?: number;
    total?: number;
    list?: TrangleAgentApiDepartmentV1Department[];
}
