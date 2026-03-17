export interface TrangleAgentApiForumV1ForumCommentsCreateReq {
    /** 评论ID */
    id?: number;
    /** 所属帖子ID */
    postId?: number;
    /** 评论发布者ID */
    userId?: number;
    /** 父评论ID */
    parentId?: number;
    /** 回复的用户ID */
    replyToUserId?: number;
    /** 评论内容 */
    content?: string;
    /** 评论状态 */
    status?: string;
    /** 点赞数 */
    likeCount?: number;
    /** 创建时间 */
    createdAt?: string;
    /** 更新时间 */
    updatedAt?: string;
}
