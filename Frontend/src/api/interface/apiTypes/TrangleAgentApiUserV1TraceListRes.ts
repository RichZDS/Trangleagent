import { type TrangleAgentApiUserV1TraceViewParams } from "../../interface";

export interface TrangleAgentApiUserV1TraceListRes {
    page?: number;
    pageSize?: number;
    total?: number;
    list?: TrangleAgentApiUserV1TraceViewParams[];
}
