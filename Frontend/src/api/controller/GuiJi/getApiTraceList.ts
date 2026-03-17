import request from "../../http.js";
import { type TrangleAgentApiUserV1TraceListRes, type DeepRequired } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 轨迹列表
 * /api/trace/list
 */
export function getApiTraceList(params: GetApiTraceListParams, config?: AxiosRequestConfig) {
    const paramsInput = {
        page: params.page,
        pageSize: params.pageSize,
        total: params.total,
        userId: params.userId,
        roleId: params.roleId,
        departmentId: params.departmentId,
        redTrace: params.redTrace,
        yellowTrace: params.yellowTrace,
        blueTrace: params.blueTrace,
    };
    return request.get<DeepRequired<TrangleAgentApiUserV1TraceListRes>>(`/api/trace/list`, {
        params: paramsInput,
        ...config,
    });
}

export interface GetApiTraceListParams {
    page?: number;
    pageSize?: number;
    total?: number;
    /** 用户ID */
    userId?: number;
    /** 角色ID */
    roleId?: number;
    /** 部门ID */
    departmentId?: number;
    /** 红轨 */
    redTrace?: number;
    /** 黄轨 */
    yellowTrace?: number;
    /** 蓝轨 */
    blueTrace?: number;
}
