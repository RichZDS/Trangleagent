export interface TrangleAgentApiForumV1ForumBoardsListReq {
    page?: number;
    pageSize?: number;
    total?: number;
    /** 所属频道ID */
    sectionId?: number;
    /** 版块名称 */
    name?: string;
    /** 版块状态 */
    status?: string;
}
