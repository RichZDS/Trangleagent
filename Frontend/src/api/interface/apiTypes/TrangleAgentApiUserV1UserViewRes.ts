export interface TrangleAgentApiUserV1UserViewRes {
    /** 用户ID */
    id?: number;
    /** 账号 */
    account?: string;
    /** 昵称 */
    nickname?: string;
    /** 性别：0未知 1男 2女 */
    gender?: number;
    /** 生日 */
    birthDate?: string;
    /** 用户类型：normal普通用户，vip为VIP用户 */
    userType?: string;
    /** VIP开始时间 */
    vipStartAt?: string;
    /** VIP结束时间 */
    vipEndAt?: string;
    /** 创建时间 */
    createdAt?: string;
    /** 现实身份/角色 */
    realityRole?: string;
    /** 异常身份/角色 */
    abnormalRole?: string;
    /** 职位 */
    jobTitle?: string;
}
