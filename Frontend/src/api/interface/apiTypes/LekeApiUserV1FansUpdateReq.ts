export interface LekeApiUserV1FansUpdateReq {
    id: number;
    /** 状态：1=粉丝 0=取关/无效 */
    status?: number;
    /** 备注名（user_id 对 fan_id 的备注/分组） */
    remark?: string;
}
