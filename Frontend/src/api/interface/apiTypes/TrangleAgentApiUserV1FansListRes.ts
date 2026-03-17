import { type TrangleAgentApiUserV1FansListItem } from "../../interface";

export interface TrangleAgentApiUserV1FansListRes {
    page?: number;
    pageSize?: number;
    total?: number;
    list?: TrangleAgentApiUserV1FansListItem[];
}
