import axios, {AxiosRequestConfig, AxiosResponse} from 'axios'
import {message} from "ant-design-vue/es";

let baseURL = "./"

if (import.meta.env.DEV){
    baseURL = 'http://127.0.0.1:8090/ui'
}

const instance = axios.create({
    baseURL: baseURL,
    timeout: 5000
})


instance.interceptors.request.use(
    (config: AxiosRequestConfig) => {
        return config
    },
    (error) => {
        error.data = {};
        error.data.msg = "服务器异常";
        message.error(error.data.msg);
        return Promise.reject(error);
    }
)



instance.interceptors.response.use(
    (res: AxiosResponse) => {
        if (res.data.code !== 200) {
            message.error(res.data.msg)
            throw new Error(res.data.msg)
        }
        return res.data.data
    },
    (error) => {
        message.error("服务器异常")
        throw new Error("未知错误")
    }
)

export default {
    access:{
        list: (p:any) => instance.get("/access",{params:p}),
        add: (p:any) => instance.post("/access",{data:p}),
        update: (p:any) => instance.put("/access",{data:p, id:p.id}),
        optionsList: (p:any) => instance.get("/access/options",{params:p}),
        tableFields: (p:any) => instance.get("/access/table-fields",{params:p}),
        tableList: (p:any) => instance.get("/access/table",{params:p}),
    },
    request:{
        list: (p:any) => instance.get("/request",{params:p}),
        update: (p:any) => instance.put("/request",{data:p}),
    }
}