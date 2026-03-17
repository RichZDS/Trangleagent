import { type LekeApiUserV1FansListItem } from "../../interface";

export interface LekeApiUserV1FansListRes {
    page?: number;
    pageSize?: number;
    total?: number;
    list?: LekeApiUserV1FansListItem[];
}
