import type {baseResponse, listDataType, paramsType} from "@/api/index";
import {useAxios} from "@/api/index";

export interface searchDataType {
    id: string;
    title: string;
}

export function articleSearchByTitle(key: string): Promise<baseResponse<listDataType<searchDataType>>> {
    return useAxios.get("/api/article/search", {
        params: {
            key: key
        }
    })
}

export interface articleListType {
    id: string
    created_at: string
    updated_at: string
    title: string
    abstract: string
    content: string
    look_count: number
    comment_count: number
    digg_count: number
    collects_count: number
    user_id: number
    user_name: string
    user_avatar: string
    category: string
    cover_id: number
    cover_url: string
}

export function articleList(params: paramsType): Promise<baseResponse<listDataType<articleListType>>> {
    return useAxios.get("/api/article/list", {params: params})
}

export interface articleDetailType {
    id: string
    created_at: string
    updated_at: string
    title: string
    abstract: string
    content: string
    look_count: number
    comment_count: number
    digg_count: number
    collects_count: number
    user_id: number
    user_name: string
    user_avatar: string
    category: string
    cover_id: number
    cover_url: string
}

export function articleDetail(id: string): Promise<baseResponse<articleDetailType>> {
    return useAxios.get("/api/article/detail" + id)
}
