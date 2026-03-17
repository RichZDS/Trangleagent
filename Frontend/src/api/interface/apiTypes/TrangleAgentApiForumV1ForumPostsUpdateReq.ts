export interface TrangleAgentApiForumV1ForumPostsUpdateReq {
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
    /** 是否置顶 */
    isPinned?: number;
    /** 是否精华 */
    isEssence?: number;
    /** 浏览量 */
    viewCount?: number;
    /** 点赞数 */
    likeCount?: number;
    /** 评论数 */
    commentCount?: number;
    /** 收藏数 */
    collectCount?: number;
    /** 最后一条评论ID */
    lastCommentId?: number;
    /** 最后评论时间 */
    lastCommentAt?: string;
    /** 发帖时间 */
    createdAt?: string;
    /** 更新时间 */
    updatedAt?: string;
}
