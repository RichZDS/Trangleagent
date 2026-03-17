import request from "../../http.js";
import { type TrangleAgentApiForumV1ForumPostsListRes, type DeepRequired } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 帖子列表
 * /api/forum/posts/list
 */
export function getApiForumPostsList(params: GetApiForumPostsListParams, config?: AxiosRequestConfig) {
    const paramsInput = {
        page: params.page,
        pageSize: params.pageSize,
        total: params.total,
        boardId: params.boardId,
        userId: params.userId,
        title: params.title,
        status: params.status,
        isPinned: params.isPinned,
        isEssence: params.isEssence,
    };
    return request.get<DeepRequired<TrangleAgentApiForumV1ForumPostsListRes>>(`/api/forum/posts/list`, {
        params: paramsInput,
        ...config,
    });
}

export interface GetApiForumPostsListParams {
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
