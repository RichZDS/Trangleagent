export interface TrangleAgentApiForumV1ForumCommentsListReq {
    page?: number;
    pageSize?: number;
    total?: number;
    /** 评论发布者ID */
    userId?: number;
    /** 所属帖子ID */
    postId?: number;
    /** 父评论ID */
    parentId?: number;
    /** 评论状态 */
    status?: string;
}
