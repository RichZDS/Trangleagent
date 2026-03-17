import { type LekeApiUserV1TraceViewParams } from "../../interface";

export interface LekeApiUserV1TraceListRes {
    page?: number;
    pageSize?: number;
    total?: number;
    list?: LekeApiUserV1TraceViewParams[];
}
