export interface LekeApiRoomV1RoomListReq {
    page?: number;
    pageSize?: number;
    total?: number;
    /** 房间号 */
    room_code?: string;
    /** 房间名称 */
    room_name?: string;
    /** 主持人用户ID */
    host_id?: number;
    /** 房间状态：0未开始 1进行中 2已结束 3已关闭 */
    status?: number;
    /** 规则系统，如COC、DND5E */
    system_name?: string;
}
