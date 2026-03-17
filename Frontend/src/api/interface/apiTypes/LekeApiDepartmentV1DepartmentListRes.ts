import { type LekeApiDepartmentV1Department } from "../../interface";

export interface LekeApiDepartmentV1DepartmentListRes {
    page?: number;
    pageSize?: number;
    total?: number;
    list?: LekeApiDepartmentV1Department[];
}
