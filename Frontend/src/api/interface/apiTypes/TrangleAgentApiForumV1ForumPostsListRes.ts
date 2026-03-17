import { type TrangleAgentInternalModelForumPostViewParams } from "../../interface";

export interface TrangleAgentApiForumV1ForumPostsListRes {
    page?: number;
    pageSize?: number;
    total?: number;
    /** 帖子列表 */
    list?: TrangleAgentInternalModelForumPostViewParams[];
}
