export interface TrangleAgentApiUserV1SubscribeCreateReq {
    userId: number;
    followId: number;
    /** 备注名（user_id 对 follow_id 的备注） */
    remark?: string;
}
