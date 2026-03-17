export interface LekeApiUserV1SubscribeUpdateReq {
    id: number;
    /** 状态：1=关注中 0=取消关注 */
    status?: number;
    /** 备注名（user_id 对 follow_id 的备注） */
    remark?: string;
}
