import request from "../../http.js";
import { type TrangleAgentApiRoomV1RoomCreateRes, type DeepRequired, type TrangleAgentApiRoomV1RoomCreateReq } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 创建房间
 * /api/room/create
 */
export function postApiRoomCreate(input?: TrangleAgentApiRoomV1RoomCreateReq, config?: AxiosRequestConfig) {
    return request.post<DeepRequired<TrangleAgentApiRoomV1RoomCreateRes>>(`/api/room/create`, input, config);
}
