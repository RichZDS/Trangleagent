export interface LekeApiDepartmentV1DepartmentListReq {
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
