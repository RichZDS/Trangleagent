import request from "../../http.js";
import { type TrangleAgentApiForumV1ForumBoardsViewRes, type DeepRequired } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 查看版块
 * /api/forum/boards/view
 */
export function getApiForumBoardsView(params: GetApiForumBoardsViewParams, config?: AxiosRequestConfig) {
    const paramsInput = {
        id: params.id,
    };
    return request.get<DeepRequired<TrangleAgentApiForumV1ForumBoardsViewRes>>(`/api/forum/boards/view`, {
        params: paramsInput,
        ...config,
    });
}

export interface GetApiForumBoardsViewParams {
    /** 版块ID */
    id?: number;
}
