export interface LekeApiUserV1SubscribeListItem {
    /** 自增主键 */
    id?: number;
    /** 用户ID（关注者，本人） */
    userId?: number;
    /** 被关注的用户ID */
    followId?: number;
    /** 状态：1=关注中 0=取消关注 */
    status?: number;
    /** 备注名（user_id 对 follow_id 的备注） */
    remark?: string;
    /** 创建时间 */
    createdAt?: string;
    /** 更新时间 */
    updatedAt?: string;
}
