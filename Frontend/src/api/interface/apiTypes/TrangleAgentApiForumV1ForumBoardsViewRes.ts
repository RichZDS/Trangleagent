export interface TrangleAgentApiForumV1ForumBoardsViewRes {
    /** 版块ID */
    id?: number;
    /** 所属频道ID */
    sectionId?: number;
    /** 版块名称 */
    name?: string;
    /** 版块简介 */
    description?: string;
    /** 版块封面图URL */
    coverImage?: string;
    /** 版块状态 */
    status?: string;
    /** 排序值 */
    displayOrder?: number;
    /** 帖子总数 */
    postCount?: number;
    /** 今日发帖数 */
    todayPostCount?: number;
    /** 最后一篇帖子ID */
    lastPostId?: number;
    /** 最后发帖时间 */
    lastPostAt?: string;
    /** 创建时间 */
    createdAt?: string;
    /** 更新时间 */
    updatedAt?: string;
}
