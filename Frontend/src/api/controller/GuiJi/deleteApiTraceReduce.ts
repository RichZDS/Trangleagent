import request from "../../http.js";
import { type TrangleAgentApiUserV1TraceReduceRes, type DeepRequired } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 删除轨迹
 * /api/trace/Reduce
 */
export function deleteApiTraceReduce(params: DeleteApiTraceReduceParams, config?: AxiosRequestConfig) {
    const paramsInput = {
        roleId: params.roleId,
        redTrace: params.redTrace,
        yellowTrace: params.yellowTrace,
        blueTrace: params.blueTrace,
    };
    return request.delete<DeepRequired<TrangleAgentApiUserV1TraceReduceRes>>(`/api/trace/Reduce`, {
        params: paramsInput,
        ...config,
    });
}

export interface DeleteApiTraceReduceParams {
    /** 角色ID */
    roleId?: number;
    /** 红轨 */
    redTrace?: number;
    /** 黄轨 */
    yellowTrace?: number;
    /** 蓝轨 */
    blueTrace?: number;
}
