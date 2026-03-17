import request from "../../http.js";
import { type TrangleAgentApiRoomV1RoomJoinRes, type DeepRequired, type TrangleAgentApiRoomV1RoomJoinReq } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 加入房间
 * /api/room/join
 */
export function postApiRoomJoin(input?: TrangleAgentApiRoomV1RoomJoinReq, config?: AxiosRequestConfig) {
    return request.post<DeepRequired<TrangleAgentApiRoomV1RoomJoinRes>>(`/api/room/join`, input, config);
}
