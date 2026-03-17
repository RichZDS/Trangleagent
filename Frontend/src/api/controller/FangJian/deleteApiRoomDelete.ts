import request from "../../http.js";
import { type TrangleAgentApiRoomV1RoomDeleteRes, type DeepRequired } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 删除房间
 * /api/room/delete
 */
export function deleteApiRoomDelete(params: DeleteApiRoomDeleteParams, config?: AxiosRequestConfig) {
    const paramsInput = {
        room_code: params.room_code,
    };
    return request.delete<DeepRequired<TrangleAgentApiRoomV1RoomDeleteRes>>(`/api/room/delete`, {
        params: paramsInput,
        ...config,
    });
}

export interface DeleteApiRoomDeleteParams {
    /** 房间号 */
    room_code?: string;
}
