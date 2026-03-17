import request from "../../http.js";
import { type TrangleAgentApiUserV1TraceViewRes, type DeepRequired } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 轨迹详情
 * /api/trace/view
 */
export function getApiTraceView(params: GetApiTraceViewParams, config?: AxiosRequestConfig) {
    const paramsInput = {
        roleId: params.roleId,
        redTrace: params.redTrace,
        yellowTrace: params.yellowTrace,
        blueTrace: params.blueTrace,
    };
    return request.get<DeepRequired<TrangleAgentApiUserV1TraceViewRes>>(`/api/trace/view`, {
        params: paramsInput,
        ...config,
    });
}

export interface GetApiTraceViewParams {
    /** 角色ID */
    roleId?: number;
    /** 红轨 */
    redTrace?: number;
    /** 黄轨 */
    yellowTrace?: number;
    /** 蓝轨 */
    blueTrace?: number;
}
