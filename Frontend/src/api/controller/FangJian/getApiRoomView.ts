import request from "../../http.js";
import { type TrangleAgentApiRoomV1RoomViewRes, type DeepRequired } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 房间详情
 * /api/room/view
 */
export function getApiRoomView(params: GetApiRoomViewParams, config?: AxiosRequestConfig) {
    const paramsInput = {
        id: params.id,
        room_code: params.room_code,
    };
    return request.get<DeepRequired<TrangleAgentApiRoomV1RoomViewRes>>(`/api/room/view`, {
        params: paramsInput,
        ...config,
    });
}

export interface GetApiRoomViewParams {
    /** 房间ID */
    id?: number;
    /** 房间号 */
    room_code?: string;
}
