import { type LekeInternalModelContainmentRepo } from "../../interface";

export interface LekeApiContainmentV1ContainmentRepoListRes {
    page?: number;
    pageSize?: number;
    total?: number;
    list?: LekeInternalModelContainmentRepo[];
}
