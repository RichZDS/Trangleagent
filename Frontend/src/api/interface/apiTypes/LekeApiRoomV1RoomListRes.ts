import { type LekeInternalModelRoomParams } from "../../interface";

export interface LekeApiRoomV1RoomListRes {
    page?: number;
    pageSize?: number;
    total?: number;
    rooms?: LekeInternalModelRoomParams[];
}
