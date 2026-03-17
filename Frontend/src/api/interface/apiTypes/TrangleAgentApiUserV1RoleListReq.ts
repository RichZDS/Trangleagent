export interface TrangleAgentApiUserV1RoleListReq {
    page?: number;
    pageSize?: number;
    total?: number;
    /** 用户ID */
    userId?: number;
    /** 部门ID */
    departmentId?: number;
}
