export interface TrangleAgentApiUserV1RolePermissionCheckReq {
    /** 角色ID */
    roleId: number;
    /** 文件需要的轨道值（0-40） */
    fileValue: number;
    /** 轨道类型：blue(蓝轨),yellow(黄轨),red(红轨) */
    trackType: 'blue' | 'yellow' | 'red';
}
