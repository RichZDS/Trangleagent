import { type TrangleAgentApiUserV1SubscribeListItem } from "../../interface";

export interface TrangleAgentApiUserV1SubscribeListRes {
    page?: number;
    pageSize?: number;
    total?: number;
    list?: TrangleAgentApiUserV1SubscribeListItem[];
}
