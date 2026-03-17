import request from "../../http.js";
import { type TrangleAgentApiForumV1ForumPostsUpdateRes, type DeepRequired, type TrangleAgentApiForumV1ForumPostsUpdateReq } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 更新帖子
 * /api/forum/posts/update
 */
export function putApiForumPostsUpdate(input?: TrangleAgentApiForumV1ForumPostsUpdateReq, config?: AxiosRequestConfig) {
    return request.put<DeepRequired<TrangleAgentApiForumV1ForumPostsUpdateRes>>(`/api/forum/posts/update`, input, config);
}
