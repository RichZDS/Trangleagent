export interface TrangleAgentApiContainmentV1ContainmentRepoViewRes {
    /** primary key(大于0为修改，其他为新增) */
    id?: number;
    /** terminal (散逸端) */
    terminalId?: number;
    /** number of contained anomalies (收容异常的数量) */
    abnormal?: string;
    /** name of the anomaly (异常体的名字) */
    anomalyName?: string;
    /** agent (特工) */
    agentName?: string;
    /** containment repository name or code (收容库) */
    repoName?: string;
}
