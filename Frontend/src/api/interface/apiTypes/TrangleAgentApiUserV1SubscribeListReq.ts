export interface TrangleAgentApiUserV1SubscribeListReq {
    page?: number;
    pageSize?: number;
    total?: number;
    userId: number;
    /** 状态：1=关注中 0=取消关注 */
    status?: number;
    /** 被关注的用户ID */
    followId?: number;
}
