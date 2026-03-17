import request from "../../http.js";
import { type TrangleAgentApiForumV1ForumPostsCreateRes, type DeepRequired, type TrangleAgentApiForumV1ForumPostsCreateReq } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 创建帖子
 * /api/forum/posts/create
 */
export function postApiForumPostsCreate(input?: TrangleAgentApiForumV1ForumPostsCreateReq, config?: AxiosRequestConfig) {
    return request.post<DeepRequired<TrangleAgentApiForumV1ForumPostsCreateRes>>(`/api/forum/posts/create`, input, config);
}
