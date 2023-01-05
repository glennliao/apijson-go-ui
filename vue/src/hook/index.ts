import api from '../api'
import {reactive, ref} from 'vue'

interface OptionItem {
    label: string
    value: string
}


export const useCommonData = () => {


    // 权限角色列表
    const roleOptions = ref<any>([])
    const rowKeyGenOptions = ref<any>([])
    const queryExecutorOptions = ref<any>([])
    const actionExecutorOptions = ref<any>([])
    const loadRoleList = () => {
        api.access.optionsList({}).then((data: any) => {
            roleOptions.value = data.RoleList.map((item: any) => {
                return {
                    value: item,
                    label: item
                }
            })

            rowKeyGenOptions.value = data.RowKeyGenList.map((item: any) => {
                return {
                    value: item,
                    label: item
                }
            })

            queryExecutorOptions.value = data.QueryExecutorList.map((item: any) => {
                return {
                    value: item,
                    label: item
                }
            })

            actionExecutorOptions.value = data.ActionExecutorList.map((item: any) => {
                return {
                    value: item,
                    label: item
                }
            })
        })
    }


    const tableFields = ref<any>([])

    function loadTableFields(table: string) {
       return  api.access.tableFields({table}).then((data: any) => {
            if (data?.List?.length) {
                tableFields.value = data.List
            }
            return data
        })
    }

    return {
        loadRoleList, roleOptions, queryExecutorOptions, actionExecutorOptions,rowKeyGenOptions,
        loadTableFields,tableFields
    }
}