export interface TrangleAgentApiForumV1ForumPostsCreateReq {
    /** 帖子ID */
    id?: number;
    /** 所属版块ID */
    boardId?: number;
    /** 发帖用户ID */
    userId?: number;
    /** 帖子标题 */
    title?: string;
    /** 帖子正文 */
    content?: string;
    /** 帖子封面图URL */
    coverImage?: string;
    /** 帖子状态 */
    status?: string;
}
