export interface TrangleAgentApiForumV1ForumBoardsCreateReq {
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
}
