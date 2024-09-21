import {type baseResponse, useAxios} from "@/api/index";

export interface userLoginType {
    account: string
    password: string
    captcha: string
    captcha_id: string
}

export function userLogin(req: userLoginType): Promise<baseResponse<string>> {
    return useAxios.post("/api/user/login", req)
}

export interface userInfoResponse {
    id: number
    created_at: string
    nick_name: string
    address: string
    avatar: string
    email: string
    role: number
}

export function userInfo(): Promise<baseResponse<userInfoResponse>> {
    return useAxios.get("/api/user/info")
}

export interface userInfoUpdateType {
    nick_name: string
    avatar: string
    email: string
}

export function userInfoUpdate(req: userInfoUpdateType): Promise<baseResponse<string>> {
    return useAxios.get("/api/user/info")
}

export function userLogout(): Promise<baseResponse<string>> {
    return useAxios.post("/api/user/logout")
}







