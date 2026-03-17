export interface TrangleAgentApiRoomV1RoomJoinReq {
    /** 房间ID */
    roomId: number;
    /** 用户ID（当前登录用户） */
    userId: number;
    /** 使用的角色卡ID（三角机构：每个角色拥有不同素质保障） */
    roleCardId: number;
}
