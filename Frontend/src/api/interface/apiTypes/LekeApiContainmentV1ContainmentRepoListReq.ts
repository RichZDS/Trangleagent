export interface LekeApiContainmentV1ContainmentRepoListReq {
    page?: number;
    pageSize?: number;
    total?: number;
    /** terminal (散逸端) */
    terminalId?: number;
    /** anomaly name */
    anomalyName?: string;
    /** agent name */
    agentName?: string;
    /** containment repo name */
    repoName?: string;
}
