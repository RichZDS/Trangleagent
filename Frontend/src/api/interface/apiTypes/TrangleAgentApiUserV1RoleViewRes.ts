export interface TrangleAgentApiUserV1RoleViewRes {
    /** 角色卡ID */
    id?: number;
    /** 所属用户ID */
    userId?: number;
    /** 所属部门ID */
    departmentId?: number;
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
    /** ARC：异常 */
    arcAbnormal?: string;
    /** ARC：现实 */
    arcReality?: string;
    /** ARC：职位 */
    arcPosition?: string;
    /** 特工名字 */
    agentName?: string;
    /** 代号 */
    codeName?: string;
    /** 性别 */
    gender?: string;
    /** Meticulousness (0-100, QA) */
    qaMeticulous?: number;
    /** Deception (0-100, QA) */
    qaDeception?: number;
    /** Vigor / Drive (0-100, QA) */
    qaVigor?: number;
    /** Empathy (0-100, QA) */
    qaEmpathy?: number;
    /** Initiative (0-100, QA) */
    qaInitiative?: number;
    /** Resilience / Persistence (0-100, QA) */
    qaResilience?: number;
    /** Presence / Charisma (0-100, QA) */
    qaPresence?: number;
    /** Professionalism (0-100, QA) */
    qaProfessional?: number;
    /** Discretion / Low profile (0-100, QA) */
    qaDiscretion?: number;
}
