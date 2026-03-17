import request from "../../http.js";
import { type TrangleAgentApiRoomV1RoomListRes, type DeepRequired } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 房间列表
 * /api/room/list
 */
export function getApiRoomList(params: GetApiRoomListParams, config?: AxiosRequestConfig) {
    const paramsInput = {
        page: params.page,
        pageSize: params.pageSize,
        total: params.total,
        room_code: params.room_code,
        room_name: params.room_name,
        host_id: params.host_id,
        status: params.status,
        system_name: params.system_name,
    };
    return request.get<DeepRequired<TrangleAgentApiRoomV1RoomListRes>>(`/api/room/list`, {
        params: paramsInput,
        ...config,
    });
}

export interface GetApiRoomListParams {
    page?: number;
    pageSize?: number;
    total?: number;
    /** 房间号 */
    room_code?: string;
    /** 房间名称 */
    room_name?: string;
    /** 主持人用户ID */
    host_id?: number;
    /** 房间状态：0未开始 1进行中 2已结束 3已关闭 */
    status?: number;
    /** 规则系统，如COC、DND5E */
    system_name?: string;
}
