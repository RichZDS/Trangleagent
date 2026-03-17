import { type TrangleAgentInternalModelContainmentRepo } from "../../interface";

export interface TrangleAgentApiContainmentV1ContainmentRepoListRes {
    page?: number;
    pageSize?: number;
    total?: number;
    list?: TrangleAgentInternalModelContainmentRepo[];
}
