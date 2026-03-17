export interface LekeApiContainmentV1ContainmentRepoUpdateReq {
    /** primary key，大于0为修改，其他为新增 */
    id?: number;
    /** terminal (散逸端) */
    terminalId?: number;
    /** number of contained anomalies */
    abnormal?: string;
    /** anomaly name */
    anomalyName?: string;
    /** agent name */
    agentName?: string;
    /** containment repo name */
    repoName?: string;
}
