import request from "../../http.js";
import { type TrangleAgentApiForumV1ForumBoardsDeleteRes, type DeepRequired } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 删除版块
 * /api/forum/boards/delete
 */
export function deleteApiForumBoardsDelete(params: DeleteApiForumBoardsDeleteParams, config?: AxiosRequestConfig) {
    const paramsInput = {
        id: params.id,
    };
    return request.delete<DeepRequired<TrangleAgentApiForumV1ForumBoardsDeleteRes>>(`/api/forum/boards/delete`, {
        params: paramsInput,
        ...config,
    });
}

export interface DeleteApiForumBoardsDeleteParams {
    /** 版块ID */
    id?: number;
}
