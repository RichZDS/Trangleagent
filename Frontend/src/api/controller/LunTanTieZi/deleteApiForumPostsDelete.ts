import request from "../../http.js";
import { type TrangleAgentApiForumV1ForumPostsDeleteRes, type DeepRequired } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 删除帖子
 * /api/forum/posts/delete
 */
export function deleteApiForumPostsDelete(params: DeleteApiForumPostsDeleteParams, config?: AxiosRequestConfig) {
    const paramsInput = {
        id: params.id,
    };
    return request.delete<DeepRequired<TrangleAgentApiForumV1ForumPostsDeleteRes>>(`/api/forum/posts/delete`, {
        params: paramsInput,
        ...config,
    });
}

export interface DeleteApiForumPostsDeleteParams {
    /** 帖子ID */
    id?: number;
}
