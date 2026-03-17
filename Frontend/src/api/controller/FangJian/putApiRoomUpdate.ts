import request from "../../http.js";
import { type TrangleAgentApiRoomV1RoomUpdateRes, type DeepRequired, type TrangleAgentApiRoomV1RoomUpdateReq } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 更新房间
 * /api/room/update
 */
export function putApiRoomUpdate(input?: TrangleAgentApiRoomV1RoomUpdateReq, config?: AxiosRequestConfig) {
    return request.put<DeepRequired<TrangleAgentApiRoomV1RoomUpdateRes>>(`/api/room/update`, input, config);
}
