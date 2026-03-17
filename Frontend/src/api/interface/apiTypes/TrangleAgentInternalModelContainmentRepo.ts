export interface TrangleAgentInternalModelContainmentRepo {
    /** number of contained anomalies (收容异常的数量) */
    abnormal?: number;
    /** agent (特工) */
    agentName?: string;
    /** name of the anomaly (异常体的名字) */
    anomalyName?: string;
    /** primary key(大于0为修改，其他为新增) */
    id?: number;
    /** containment repository name or code (收容库) */
    repoName?: string;
    /** terminal (散逸端) */
    terminalId?: number;
}
