export interface LekeApiUserV1FansListItem {
    /** 自增主键 */
    id?: number;
    /** 用户ID（被关注者，本人） */
    userId?: number;
    /** 粉丝用户ID */
    fanId?: number;
    /** 状态：1=粉丝 0=取关/无效 */
    status?: number;
    /** 备注名（user_id 对 fan_id 的备注/分组） */
    remark?: string;
    /** 创建时间 */
    createdAt?: string;
    /** 更新时间 */
    updatedAt?: string;
}
