import { type TrangleAgentInternalModelForumCommentViewParams } from "../../interface";

export interface TrangleAgentApiForumV1ForumCommentsListRes {
    page?: number;
    pageSize?: number;
    total?: number;
    /** 评论列表 */
    list?: TrangleAgentInternalModelForumCommentViewParams[];
}
