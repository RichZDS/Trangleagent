export interface LekeApiRoomV1RoomUpdateReq {
    /** 房间号(作为更新定位字段) */
    room_code?: string;
    /** 自增主键 */
    id?: number;
    /** 房间全局唯一ID（建议UUID） */
    roomId?: string;
    /** 房间名称 */
    roomName?: string;
    /** 主持人用户ID（GM/KP/经理，对应用户表id） */
    hostId?: number;
    /** 主持人昵称（冗余字段，可选） */
    hostNickname?: string;
    /** 最大玩家人数（不含主持人，可按需要约定） */
    maxPlayers?: number;
    /** 当前玩家人数（不含主持人） */
    currentPlayers?: number;
    /** 是否有密码：0无 1有 */
    hasPassword?: number;
    /** 房间密码（建议存加密/哈希后的密码） */
    roomPassword?: string;
    /** 是否私密房：0公开 1私密 */
    isPrivate?: number;
    /** 房间状态：0未开始 1进行中 2已结束 3已关闭 */
    status?: number;
    /** 规则系统：如 COC、DND5E 等 */
    systemName?: string;
    /** 模组/剧本名称 */
    scenarioName?: string;
    /** 房间简介/招募说明 */
    description?: string;
    /** 房间创建时间 */
    createdAt?: string;
    /** 开团时间 */
    startedAt?: string;
    /** 结束时间 */
    endedAt?: string;
    /** 信息最近更新时间 */
    updatedAt?: string;
}
