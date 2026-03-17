export interface LekeApiUserV1FansCreateReq {
    userId: number;
    fanId: number;
    /** 备注名（user_id 对 fan_id 的备注/分组） */
    remark?: string;
}
