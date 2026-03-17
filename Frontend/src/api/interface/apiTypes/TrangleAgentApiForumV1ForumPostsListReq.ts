export interface TrangleAgentApiForumV1ForumPostsListReq {
    page?: number;
    pageSize?: number;
    total?: number;
    /** 所属版块ID */
    boardId?: number;
    /** 发帖用户ID */
    userId?: number;
    /** 帖子标题 */
    title?: string;
    /** 帖子状态 */
    status?: string;
    /** 是否置顶 */
    isPinned?: number;
    /** 是否精华 */
    isEssence?: number;
}
