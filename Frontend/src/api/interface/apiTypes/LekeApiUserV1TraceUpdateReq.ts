export interface LekeApiUserV1TraceUpdateReq {
    userId?: number;
    /** 角色ID */
    roleId?: number;
    departmentId?: number;
    /** 红轨 */
    redTrace?: number;
    /** 黄轨 */
    yellowTrace?: number;
    /** 蓝轨 */
    blueTrace?: number;
}
