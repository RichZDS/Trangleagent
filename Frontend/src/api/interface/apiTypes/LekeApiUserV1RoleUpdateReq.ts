export interface LekeApiUserV1RoleUpdateReq {
    /** 角色ID */
    id: number;
    /** 所属部门ID */
    departmentId?: number;
    /** 特工名字 */
    agentName: string;
    /** 代号 */
    codeName: string;
    /** 性别 */
    gender: 'male' | 'female' | 'other';
    /** ARC：异常 */
    arcAbnormal?: string;
    /** ARC：现实 */
    arcReality?: string;
    /** ARC：职位 */
    arcPosition?: string;
    /** 嘉奖次数 */
    commendation?: number;
    /** 申戒次数 */
    reprimand?: number;
    /** 蓝轨（0-40） */
    blueTrack?: number;
    /** 黄轨（0-40） */
    yellowTrack?: number;
    /** 红轨（0-40） */
    redTrack?: number;
}
