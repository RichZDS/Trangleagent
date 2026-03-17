export interface LekeApiUserV1UserListReq {
    page?: number;
    pageSize?: number;
    total?: number;
    /** 账号 */
    account?: string;
    /** 昵称 */
    nickname?: string;
}
