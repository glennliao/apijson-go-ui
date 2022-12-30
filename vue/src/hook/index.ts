import api from '../api'
import {reactive, ref} from 'vue'

interface OptionItem {
    label: string
    value: string
}


export const useCommonData = () => {


    // 权限角色列表
    const roleOptions = ref<any>([])
    const loadRoleList = () => {
        api.access.roleList({}).then((data: any) => {
            if (data?.List?.length) {
                roleOptions.value = data.List.map((item: any) => {
                    return {
                        value: item,
                        label: item
                    }
                })
            }
        })
    }


    const tableFields = ref<any>([])

    function loadTableFields(table: string) {
       return  api.access.tableFields({table}).then((data: any) => {
            if (data?.List?.length) {
                tableFields.value = data.List
            }
        })
    }

    return {
        loadRoleList, roleOptions,
        loadTableFields,tableFields
    }
}