<template>
  <a-modal
      v-model:visible="visible"
      title="edit"
      :maskClosable="false"
      :closable="true"
      destroyOnClose
      style="top: 20px"
      width="880px"
  >
    <a-form
        :model="formInfo"
        :label-col="{ span: 3 }"
        :wrapper-col="{ span: 21 }"
        name="request"
        ref="formRef"
        layout="vertical"
    >
      <a-tabs
          tab-position="left"
          :style="{ height: '100%' }"
          size="small"
      >
        <a-tab-pane key="basic" tab="basic">
          <a-row :gutter="8">
            <a-col :span="12">
              <a-form-item label="name" :label-col="{col:2}" name="name"
                           :rules="{ required: true, message: 'name不能为空' }">
                <a-input
                    v-model:value="formInfo['name']"
                    placeholder="name"
                    style="width: 100%"
                />
              </a-form-item>
            </a-col>
            <a-col :span="12">
              <a-form-item label="alias" name="alias">
                <a-input
                    v-model:value="formInfo['alias']"
                    placeholder="alias"
                    style="width: 100%"
                />
              </a-form-item>
            </a-col>
          </a-row>

          <a-form-item label="debug" name="debug">
            <a-switch
                v-model:checked="formInfo['debug']"
                :checkedValue="1"
                :unCheckedValue="0"
                checked-children="是"
                un-checked-children="否"
            />
          </a-form-item>

          <a-form-item label="row_key" name="row_key" :rules="{ required: true, message: 'row_key不能为空' }">
            <a-select
                v-model:value="formInfo['row_key']"
                :options="tableFields"
                mode="multiple"
                :maxTagCount="4"
                placeholder="请选择rowKey"
                style="width: 100%"
            />
          </a-form-item>

          <a-form-item label="detail" name="detail">
            <a-textarea
                v-model:value="formInfo['detail']"
                placeholder="请输入detail"
                :maxlength="140"
                show-count
                :rows="3"
            />
          </a-form-item>


        </a-tab-pane>
        <a-tab-pane key="role" tab="role">
          <a-form-item label="get" name="get">
            <a-select
                v-model:value="formInfo['get']"
                :options="roleOptions"
                mode="multiple"
                :maxTagCount="4"
                placeholder="请选择角色"
            />
          </a-form-item>

          <a-form-item label="head" name="head">
            <a-select
                v-model:value="formInfo['head']"
                :options="roleOptions"
                mode="multiple"
                :maxTagCount="4"
                placeholder="请选择角色"
            />
          </a-form-item>

          <a-form-item label="gets" name="gets">
            <a-select
                v-model:value="formInfo['gets']"
                :options="roleOptions"
                mode="multiple"
                :maxTagCount="4"
                placeholder="请选择角色"
            />
          </a-form-item>

          <a-form-item label="heads" name="heads">
            <a-select
                v-model:value="formInfo['heads']"
                :options="roleOptions"
                mode="multiple"
                :maxTagCount="4"
                placeholder="请选择角色"
            />
          </a-form-item>

          <a-form-item label="post" name="post">
            <a-select
                v-model:value="formInfo['post']"
                :options="roleOptions"
                mode="multiple"
                :maxTagCount="4"
                placeholder="请选择角色"
            />
          </a-form-item>

          <a-form-item label="put" name="put">
            <a-select
                v-model:value="formInfo['put']"
                :options="roleOptions"
                mode="multiple"
                :maxTagCount="4"
                placeholder="请选择角色"
            />
          </a-form-item>

          <a-form-item label="delete" name="delete">
            <a-select
                v-model:value="formInfo['delete']"
                :options="roleOptions"
                mode="multiple"
                :maxTagCount="4"
                placeholder="请选择角色"
            />
          </a-form-item>


        </a-tab-pane>
        <a-tab-pane key="fields_get" tab="fields_get">
          <div>

            <div>
              <FieldsGetRoles
                  role="default"
                  :fields-get="formInfo.fields_get"
                  :inOptions="formInfo.fields_get['default'].inOptions"
                  :outOptions="formInfo.fields_get['default'].outOptions"
                  @change="fieldsGetChange"
              />
            </div>

            <div>
              <a-checkbox-group
                  v-model:value="fieldsGetRoleList"
                  name="tag"
                  :options="roleOptions"
              />
            </div>



            <a-card size="small">
              <div v-if="fieldsGetRoleList.length">
                            <FieldsGetRoles
                                v-for="role in fieldsGetRoleList"
                                :key="role"
                                :role="role"
                                :fields-get="formInfo.fields_get"
                                :inOptions="formInfo.fields_get[role].inOptions"
                                :outOptions="formInfo.fields_get[role].outOptions"
                                @change="fieldsGetChange"
                            />
              </div>
            </a-card>
          </div>
        </a-tab-pane>
      </a-tabs>


    </a-form>
    <template #footer>
      <a-button style="margin-right: 8px" @click="visible = false">取消</a-button>
      <a-button type="primary" @click="submit">提交</a-button>
    </template>
  </a-modal>
</template>

<script lang="ts">
export default {
  name:"AccessEdit"
}
</script>

<script lang="ts" setup>
import {ref, watch} from "vue";
import api from '../../api'
import {message} from "ant-design-vue";
import {useCommonData} from '../../hook'
import FieldsGetRoles from './FieldsGetRoles.vue'
import {methods} from "../../consts";


const visible = ref(false)

const emit = defineEmits(['update:visible', 'refresh'])

// 初始化
const {loadRoleList, roleOptions,loadTableFields, tableFields} = useCommonData()
const formInfo = ref<any>({})


// 监听roles列表的变化， 对fieldsGet进行初始化和增减操作
const fieldsGetRoleList = ref<any>([])
watch(() => fieldsGetRoleList.value, (newVal, oldVal) => {
  let key = ''

  if (oldVal.length) {
    key = oldVal.filter((item: any) => !newVal.includes(item))[0] // 此处为每次只点击一个
  }

  if (key) {
    delete formInfo.value.fields_get[key]
  } else {
    newVal.forEach((item: string) => {
      if (!Object.keys(formInfo.value.fields_get).includes(item)) {
        formInfo.value.fields_get[item] = {
          in: {},
          out: {},
          inOptions: tableFields.value.map((item: any) => {
            return {
              ...item,
              checked: true
            }
          }),
          outOptions: tableFields.value.map((item: any) => {
            return {
              ...item,
              checked: true
            }
          })
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
        formInfo.value.fields_get[opt.key].in[opt.item.value] = ['*']
      } else {
        formInfo.value.fields_get[opt.key].in[opt.item.value] = opt.valueArr.filter((item: any) => item !== '*')
      }
    },
    "out": () => {
      formInfo.value.fields_get[opt.key].out = opt.out
    }
  }

  let h = handlers[opt.type.toLowerCase()]

  if (h) {
    h()
  }
}



// 提交方法
const formRef = ref<any>(null)
const submit = () => {

  let data = JSON.parse(JSON.stringify(formInfo.value))
  formRef.value.validate().then(() => {

    data.fields_get = data.fields_get || {}

    Object.keys(data.fields_get).forEach((key: string) => {
      delete data.fields_get[key].inOptions
      delete data.fields_get[key].outOptions
    })

    methods.forEach(method => {
      data[method] = data[method].map((roleItem: string) => {
          // @ts-ignore
        if (roleItem.value) {
          // @ts-ignore
          return roleItem.value
        }
        return roleItem
      })
    })



    data["row_key"] = data["row_key"].map((item:any)=>item.value?item.value:item).join(",")

    delete data["created_at"]

    api.access.update(data).then(() => {
      message.success('编辑成功')
      emit('refresh')
      visible.value = false
    })
  })
}


function open(info: any) {
  visible.value = true
  info = JSON.parse(JSON.stringify(info))
  info.row_key = info.row_key.split(",").map((item: string) => {
    return {
      value: item
    }
  })

  info.fields_get = info.fields_get || "{}"
  info.fields_get = JSON.parse(info.fields_get)

  if(Object.keys(info.fields_get).length === 0){
    info.fields_get = {
      default:{
        in:{},
        out:{},
      }
    }
  }

  loadRoleList()
  loadTableFields(info.name).then(()=>{

    Object.keys(info.fields_get).forEach((role: string) => {
      info.fields_get[role].inOptions = tableFields.value.map((item: any) => {
        return {
          value: item.value,
          label: item.value,
          checked: !(info.fields_get[role].in[item.value] !== undefined)
        }
      })


      info.fields_get[role].outOptions = tableFields.value.map((item: any) => {
        return {
          value: item.value,
          label: item.value,
          checked: !(Object.keys(info.fields_get[role].out).includes(item.value))
        }
      })
      if(role != "default"){
        fieldsGetRoleList.value.push(role)
      }
    })

    console.log(info)

    formInfo.value = info
  })

}


defineExpose({
  open
})
</script>

<style lang="less" scoped>
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
