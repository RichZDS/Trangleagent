export interface LekeApiDepartmentV1Department {
    /** 自增主键 */
    id?: number;
    /** 所属用户ID（对应 users.id） */
    userId?: number;
    /** 分部名称 */
    branchName?: string;
    /** 分部散逸端的数量 */
    terminalCount?: number;
    /** 分部当前天气/气候描述 */
    weather?: string;
    /** 分部经理名称 */
    managerName?: string;
    /** 分部地址 */
    location?: string;
    /** 创建时间 */
    createdAt?: string;
    /** 更新时间 */
    updatedAt?: string;
}
