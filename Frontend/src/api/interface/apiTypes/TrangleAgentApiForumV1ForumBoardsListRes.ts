import { type TrangleAgentInternalModelForumBoardViewParams } from "../../interface";

export interface TrangleAgentApiForumV1ForumBoardsListRes {
    page?: number;
    pageSize?: number;
    total?: number;
    /** 版块列表 */
    list?: TrangleAgentInternalModelForumBoardViewParams[];
}
