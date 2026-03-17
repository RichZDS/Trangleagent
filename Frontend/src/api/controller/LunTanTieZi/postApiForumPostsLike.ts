import request from "../../http.js";
import { type AxiosRequestConfig } from "axios";

/**
 * 点赞帖子
 * /api/forum/posts/like
 */
export function postApiForumPostsLike(id: number, config?: AxiosRequestConfig) {
  return request.post(`/api/forum/posts/like`, { id: Number(id) }, config);
}
