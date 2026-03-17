import request from "../../http.js";
import { type TrangleAgentApiForumV1ForumCommentsListRes, type DeepRequired } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 评论列表
 * /api/forum/comments/list
 */
export function getApiForumCommentsList(params: GetApiForumCommentsListParams, config?: AxiosRequestConfig) {
    const paramsInput = {
        page: params.page,
        pageSize: params.pageSize,
        total: params.total,
        userId: params.userId,
        postId: params.postId,
        parentId: params.parentId,
        status: params.status,
    };
    return request.get<DeepRequired<TrangleAgentApiForumV1ForumCommentsListRes>>(`/api/forum/comments/list`, {
        params: paramsInput,
        ...config,
    });
}

export interface GetApiForumCommentsListParams {
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
