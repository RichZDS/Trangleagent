import request from "../../http.js";
import { type AxiosRequestConfig } from "axios";

export interface CheckInReq {
  userId: number;
}

export interface CheckInRes {
  exp: number;
  level: number;
  addedExp: number;
  message: string;
}

/**
 * 每日签到
 * /api/user/checkin
 */
export function postApiUserCheckin(input: CheckInReq, config?: AxiosRequestConfig) {
  return request.post<CheckInRes>(`/api/user/checkin`, input, config);
}
