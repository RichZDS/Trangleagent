import request from "../../http.js";
import { type TrangleAgentApiForumV1ForumBoardsCreateRes, type DeepRequired, type TrangleAgentApiForumV1ForumBoardsCreateReq } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 创建版块
 * /api/forum/boards/create
 */
export function postApiForumBoardsCreate(input?: TrangleAgentApiForumV1ForumBoardsCreateReq, config?: AxiosRequestConfig) {
    return request.post<DeepRequired<TrangleAgentApiForumV1ForumBoardsCreateRes>>(`/api/forum/boards/create`, input, config);
}
