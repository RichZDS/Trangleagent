export interface TrangleAgentApiUserV1FansListReq {
    page?: number;
    pageSize?: number;
    total?: number;
    userId: number;
    /** 状态：1=粉丝 0=取关/无效 */
    status?: number;
}
