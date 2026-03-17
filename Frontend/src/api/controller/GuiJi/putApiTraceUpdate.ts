import request from "../../http.js";
import { type TrangleAgentApiUserV1TraceUpdateRes, type DeepRequired, type TrangleAgentApiUserV1TraceUpdateReq } from "../../interface";
import { type AxiosRequestConfig } from "axios";

/**
 * 更新轨迹
 * /api/trace/update
 */
export function putApiTraceUpdate(input?: TrangleAgentApiUserV1TraceUpdateReq, config?: AxiosRequestConfig) {
    return request.put<DeepRequired<TrangleAgentApiUserV1TraceUpdateRes>>(`/api/trace/update`, input, config);
}
