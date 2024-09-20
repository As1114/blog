import {type baseResponse, useAxios} from "@/api/index";

export interface captchaType {
    pic_path: string;
    captcha_id: string;
}

export function captchaGen(): Promise<baseResponse<captchaType>> {
    return useAxios.get("/api/captcha")
}
