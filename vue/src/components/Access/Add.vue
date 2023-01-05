<template>
  <a-modal
      v-model:visible="visible"
      title="Add"
      :maskClosable="false"
      :closable="true"
      destroyOnClose
      style="top: 20px"
      width="1024px"
  >
  <div style="">
    <div class="flex">
      <div class="flex flex-wrap table-left">
        <div class="header">table</div>
        <div class="p-2" style="height: 620px; overflow-y: auto">
          <a-checkbox-group
              v-model:value="tableNameList"
              name="tag"
              :options="tableOptions"
          />
        </div>
      </div>

      <div v-if="tableArray?.length" class="flex-1 ml-2">
        <div style="height: 650px; overflow-y: auto">
          <div class="mb-2 mr-2" v-for="(table, index) in tableArray" :key="table">
            <a-card :title="table.tableName" size="small">
              <template #extra>
                <span class="cursor-pointer text-1890ff" @click="remove(table.tableName)">移除</span>
              </template>
              <a-form :model="table" ref="tableFormRef" :label-col="{ span: 2 }" :wrapper-col="{ span: 22 }">

                <a-form-item label="alias" :name="['access', 'alias']"
                             :rules="{ required: true, message: 'alias不能为空' }">
                  <a-input
                      v-model:value="table['access'].alias"
                      placeholder="请输入tableAlias"
                  />
                </a-form-item>

                <a-form-item label="tag" :name="['request', 'tag']" :rules="{ required: true, message: 'tag不能为空' }">
                  <a-input
                      v-model:value="table['request'].tag"
                      placeholder="请输入tableAlias"
                  />
                </a-form-item>

                <a-form-item
                    label="rowKey"
                    :name="['access', 'rowKey']"
                    :rules="{ required: true, message: 'rowKey不能为空' }">
                  <a-select
                      v-model:value="table['access'].rowKey"
                      :options="table['access'].rowKeyOptions"
                      mode="multiple"
                      :maxTagCount="4"
                      placeholder="请选择rowKey"
                  />
                </a-form-item>

                <a-form-item label="request" name="request">
                  <a-checkbox-group
                      v-model:value="table['request'].request"
                      name="tag"
                      :options="requestOptions"
                  />
                </a-form-item>

                <a-form-item label="fieldsGet" name="fieldsGet">
                  <a-card size="small">
                    <template #title>
                      <div style="font-size: 14px;">操作设置</div>
                    </template>
                    <template #extra>
                      <a-checkbox-group
                          v-model:value="table['access'].roles"
                          name="tag"
                          :options="roleOptions"
                          @change="roleChange(index, $event)"
                      />
                    </template>

                    <div v-if="table['access'].roles.length">
                      <FieldsGetRoles
                          v-for="role in table['access'].roles"
                          :key="role"
                          :inOptions="table['access'].rowKeyOptions"
                          :outOptions="table['access'].rowKeyOptions"
                          :role="role"
                          :fields-get="table['access'].fieldsGet"
                          @change="fieldsGetChange"
                      />
                    </div>
                    <div v-else>null</div>
                  </a-card>
                </a-form-item>
              </a-form>
            </a-card>
          </div>
        </div>
      </div>
      <div v-else class="m-auto text-gray-400">
        请选择表
      </div>
    </div>
  </div>
    <template #footer>
      <div class="box flex justify-end pr-4 z-99">
        <a-button @click="visible = false" class="mr-2">取消</a-button>
        <a-button type="primary" @click="submit">提交</a-button>
      </div>
    </template>
  </a-modal>
</template>


<script lang="ts" setup>
import {ref, watch} from "vue";
import api from '../../api'
import {useRouter} from 'vue-router'
import {message} from "ant-design-vue";
import {useCommonData} from '../../hook'
import FieldsGetRoles from './FieldsGetRoles.vue'
import camelCase from 'lodash-es/camelCase'
import upperFirst from 'lodash-es/upperFirst'

const visible = ref(false)



const {loadRoleList, roleOptions,loadTableFields, tableFields} = useCommonData()
const loadTableList = () => {
  api.access.tableList({}).then((data: any) => {
    if (data?.List?.length) {
      tableOptions.value = data.List
      tableData.value = data.List
    }
  })
}


const requestOptions = [
  {
    label:"POST",value:"POST"
  },
  {
    label:"PUT",value:"PUT"
  },
  {
    label: "DELETE", value: "DELETE"
  }
]


function open() {

  loadTableList() // 获取左侧表名


  visible.value = true
}
defineExpose({
  open
})

// 初始化

const tableNameList = ref<any[]>([])
const tableOptions = ref([])
const tableData = ref<any>([])

const tableArray = ref<any[]>([])


watch(() => tableNameList.value, function (newVal, oldVal) {
  const arr = newVal.filter((item: any) => !oldVal.includes(item))
  if (arr.length) {

    let tableName = arr[0]

    const rowKeyOptions = ref([])

    loadTableFields(tableName).then(data=>{
      rowKeyOptions.value = data.List
    })

    tableArray.value.push({
      tableName: tableName,
      request: {
        tag: upperFirst(camelCase(tableName)),
        request: ['POST', 'PUT', 'DELETE']
      },
      access: {
        alias: upperFirst(camelCase(tableName)),
        rowKey: ref([]),
        rowKeyOptions,
        fieldsGet: {},
        roles: []
      }
    })
  } else {
    tableArray.value = tableArray.value.filter((item) => newVal.includes(item.tableName))
  }
})

// 移除方法
const remove = (name: string) => {
  tableNameList.value = tableNameList.value.filter(item => item !== name)
}

// roles @change 事件触发 反选当前点击的多选框
const roles = ref<any>([])
const current = ref<number>(0)
const roleChange = (index: number, checkedValue: any) => {
  // index => 当前操作的表索引(number)
  // checkedValue => 已选中的role(Array[])
  roles.value = checkedValue
  current.value = index
}

// 监听roles列表的变化， 对fieldsGet进行初始化和增减操作
watch([() => roles.value, () => current.value], ([newVal, newCur], [oldVal, oldCur]) => {
  let key = ''

  if (oldVal.length && newCur === oldCur) {
    key = oldVal.filter((item: any) => !newVal.includes(item))[0]
  }

  if (key) {
    delete tableArray.value[current.value].access.fieldsGet[key]
  } else {
    newVal.forEach((item: string) => {
      if (!Object.keys(tableArray.value[current.value].access.fieldsGet).includes(item)) {
        tableArray.value[current.value].access.fieldsGet[item] = {
          in: {},
          out: {},
        }
      }
    })
  }
})

// 调用fieldsGetRoles子组件抛出的事件，并对相应的参数进行处理
const fieldsGetChange = (opt: any) => {
  let handlers: any = {
    "in": () => {
      if (opt.flag) {
        tableArray.value[current.value].access.fieldsGet[opt.key].in[opt.item.value] = ['*']
      } else {
        tableArray.value[current.value].access.fieldsGet[opt.key].in[opt.item.value] = opt.valueArr.filter((item: any) => item !== '*')
      }
    },
    "out": () => {
      tableArray.value[current.value].access.fieldsGet[opt.key].out = opt.out
    }
  }

  let h = handlers[opt.type.toLowerCase()]

  if (h) {
    h()
  }
}

// 提交方法
const submit = () => {
  if (!tableNameList.value.length) return message.warn('请先选择table')
  let list = JSON.parse(JSON.stringify(tableArray.value))

  list.map((item: any) => {
    if (item?.request?.request?.length) {
      item.request.request = item.request.request.map((v: any) => {
        return {
          method: v
        }
      })
    }
    delete item.access.roles
    delete item.access.rowKeyOptions
    return item
  })
  if (list.length) {
    api.access.add(list).then(() => {
      message.success('添加成功')
    })
  }
}


</script>

<style scoped lang="less">
.table-left {
  width: 180px;
  border: 1px solid #f0f0f0;

  .header {
    width: 100%;
    padding: 4px 10px;
    border-bottom: 1px solid #f0f0f0;
  }
}


.content {
  border: 1px solid #f0f0f0;
  display: flex;

  .in {
    width: 50%;
    border-right: 1px solid #f0f0f0;
  }

  .out {
    flex: 1;
  }

  .title {
    min-height: 40px;
    line-height: 2;
    font-size: 16px;
    border-bottom: 1px solid #f0f0f0;
    padding: 0 10px;

    span {
      font-weight: bold;
      font-size: 18px;
      margin-right: 10px;
    }
  }

}
</style>
