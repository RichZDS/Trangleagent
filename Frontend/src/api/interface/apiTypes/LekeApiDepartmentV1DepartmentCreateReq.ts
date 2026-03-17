export interface LekeApiDepartmentV1DepartmentCreateReq {
    userId: number;
    branchName: string;
    /** 分部散逸端的数量 */
    terminalCount?: number;
    /** 分部当前天气/气候描述 */
    weather?: string;
    /** 分部经理名称 */
    managerName?: string;
    /** 分部地址 */
    location?: string;
}
