import request from "../../http.js";
import { type TrangleAgentApiForumV1ForumBoardsUpdateRes, type DeepRequired, type TrangleAgentApiForumV1ForumBoardsUpdateReq } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 更新版块
 * /api/forum/boards/update
 */
export function putApiForumBoardsUpdate(input?: TrangleAgentApiForumV1ForumBoardsUpdateReq, config?: AxiosRequestConfig) {
    return request.put<DeepRequired<TrangleAgentApiForumV1ForumBoardsUpdateRes>>(`/api/forum/boards/update`, input, config);
}
