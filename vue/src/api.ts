import axios, {AxiosRequestConfig, AxiosResponse} from 'axios'
import {message} from "ant-design-vue/es";

// axios.defaults.baseURL = "http://127.0.0.1:8088"



const instance = axios.create({
    baseURL: 'http://127.0.0.1:8088',
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
        update: (p:any) => instance.put("/access",{data:p, id:p.id}),
        roleList: (p:any) => instance.get("/access/role",{params:p}),
        tableFields: (p:any) => instance.get("/access/table-fields",{params:p}),
    },
    request:{
        list: (p:any) => instance.get("/request",{params:p}),
        update: (p:any) => instance.put("/request",{data:p}),
    }
}