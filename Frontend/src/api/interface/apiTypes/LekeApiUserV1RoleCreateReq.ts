export interface LekeApiUserV1RoleCreateReq {
    /** 所属用户ID */
    userId: number;
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
}
