import request from "../../http.js";
import { type TrangleAgentApiForumV1ForumCommentsUpdateRes, type DeepRequired, type TrangleAgentApiForumV1ForumCommentsUpdateReq } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 更新评论
 * /api/forum/comments/update
 */
export function putApiForumCommentsUpdate(input?: TrangleAgentApiForumV1ForumCommentsUpdateReq, config?: AxiosRequestConfig) {
    return request.put<DeepRequired<TrangleAgentApiForumV1ForumCommentsUpdateRes>>(`/api/forum/comments/update`, input, config);
}
