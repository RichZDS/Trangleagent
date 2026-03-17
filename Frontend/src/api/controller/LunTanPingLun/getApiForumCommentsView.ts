import request from "../../http.js";
import { type TrangleAgentApiForumV1ForumCommentsViewRes, type DeepRequired } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 查看评论
 * /api/forum/comments/view
 */
export function getApiForumCommentsView(params: GetApiForumCommentsViewParams, config?: AxiosRequestConfig) {
    const paramsInput = {
        id: params.id,
    };
    return request.get<DeepRequired<TrangleAgentApiForumV1ForumCommentsViewRes>>(`/api/forum/comments/view`, {
        params: paramsInput,
        ...config,
    });
}

export interface GetApiForumCommentsViewParams {
    /** 评论ID */
    id?: number;
}
