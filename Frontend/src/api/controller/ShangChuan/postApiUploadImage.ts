import request from "../../http.js";
import { type AxiosRequestConfig } from "axios";

/**
 * 上传图片到 MinIO
 * POST /api/upload/image
 * @param file File 对象
 * @returns 返回 { url: string }
 */
export function postApiUploadImage(file: File, config?: AxiosRequestConfig) {
  const form = new FormData();
  form.append("file", file);
  return request.post<{ url: string }>("/api/upload/image", form, {
    ...config,
    headers: {
      "Content-Type": "multipart/form-data",
      ...config?.headers,
    },
  });
}
