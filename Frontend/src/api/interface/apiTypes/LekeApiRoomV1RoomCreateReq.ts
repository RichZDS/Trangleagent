export interface LekeApiRoomV1RoomCreateReq {
    /** 房间号 */
    room_code?: string;
    /** 房间名称 */
    room_name?: string;
    /** 主持人用户ID */
    host_id?: number;
    /** 最大玩家人数 */
    max_players?: number;
    /** 是否有密码：0无 1有 */
    has_password?: number;
    /** 房间密码(建议哈希) */
    room_password?: string;
    /** 是否私密房：0公开 1私密 */
    is_private?: number;
    /** 规则系统，如COC、DND5E */
    system_name?: string;
    /** 模组/剧本名称 */
    scenario_name?: string;
    /** 房间简介/招募说明 */
    description?: string;
}
