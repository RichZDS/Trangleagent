export interface LekeApiUserV1TraceViewRes {
    /** 用户ID */
    userId?: number;
    /** 角色ID */
    roleId?: number;
    /** 部门ID */
    departmentId?: number;
    /** 红轨 */
    redTrace?: number;
    /** 黄轨 */
    yellowTrace?: number;
    /** 蓝轨 */
    blueTrace?: number;
}
