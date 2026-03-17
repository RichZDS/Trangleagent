import request from "../../http.js";
import { type TrangleAgentApiForumV1ForumCommentsCreateRes, type DeepRequired, type TrangleAgentApiForumV1ForumCommentsCreateReq } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 创建评论
 * /api/forum/comments/create
 */
export function postApiForumCommentsCreate(input?: TrangleAgentApiForumV1ForumCommentsCreateReq, config?: AxiosRequestConfig) {
    return request.post<DeepRequired<TrangleAgentApiForumV1ForumCommentsCreateRes>>(`/api/forum/comments/create`, input, config);
}
