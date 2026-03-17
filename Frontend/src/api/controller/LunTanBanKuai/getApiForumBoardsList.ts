import request from "../../http.js";
import { type TrangleAgentApiForumV1ForumBoardsListRes, type DeepRequired } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 版块列表
 * /api/forum/boards/list
 */
export function getApiForumBoardsList(params: GetApiForumBoardsListParams, config?: AxiosRequestConfig) {
    const paramsInput = {
        page: params.page,
        pageSize: params.pageSize,
        total: params.total,
        sectionId: params.sectionId,
        name: params.name,
        status: params.status,
    };
    return request.get<DeepRequired<TrangleAgentApiForumV1ForumBoardsListRes>>(`/api/forum/boards/list`, {
        params: paramsInput,
        ...config,
    });
}

export interface GetApiForumBoardsListParams {
    page?: number;
    pageSize?: number;
    total?: number;
    /** 所属频道ID */
    sectionId?: number;
    /** 版块名称 */
    name?: string;
    /** 版块状态 */
    status?: string;
}
