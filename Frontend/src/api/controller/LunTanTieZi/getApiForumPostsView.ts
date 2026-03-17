import request from "../../http.js";
import { type TrangleAgentApiForumV1ForumPostsViewRes, type DeepRequired } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 查看帖子
 * /api/forum/posts/view
 */
export function getApiForumPostsView(params: GetApiForumPostsViewParams, config?: AxiosRequestConfig) {
    const paramsInput = {
        id: params.id,
    };
    return request.get<DeepRequired<TrangleAgentApiForumV1ForumPostsViewRes>>(`/api/forum/posts/view`, {
        params: paramsInput,
        ...config,
    });
}

export interface GetApiForumPostsViewParams {
    /** 帖子ID */
    id?: number;
}
