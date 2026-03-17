import { type TrangleAgentInternalModelRoomParams } from "../../interface";

export interface TrangleAgentApiRoomV1RoomListRes {
    page?: number;
    pageSize?: number;
    total?: number;
    rooms?: TrangleAgentInternalModelRoomParams[];
}
