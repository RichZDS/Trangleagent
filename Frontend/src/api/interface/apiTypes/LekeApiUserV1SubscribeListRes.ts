import { type LekeApiUserV1SubscribeListItem } from "../../interface";

export interface LekeApiUserV1SubscribeListRes {
    page?: number;
    pageSize?: number;
    total?: number;
    list?: LekeApiUserV1SubscribeListItem[];
}
