import request from "../../http.js";
import { type TrangleAgentApiForumV1ForumCommentsDeleteRes, type DeepRequired } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 删除评论
 * /api/forum/comments/delete
 */
export function deleteApiForumCommentsDelete(params: DeleteApiForumCommentsDeleteParams, config?: AxiosRequestConfig) {
    const paramsInput = {
        id: params.id,
    };
    return request.delete<DeepRequired<TrangleAgentApiForumV1ForumCommentsDeleteRes>>(`/api/forum/comments/delete`, {
        params: paramsInput,
        ...config,
    });
}

export interface DeleteApiForumCommentsDeleteParams {
    /** 评论ID */
    id?: number;
}
