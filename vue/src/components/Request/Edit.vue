<template>
  <a-modal
      v-model:visible="visible"
      title="edit"
      :maskClosable="false"
      :closable="false"
      destroyOnClose
      placement="right"
      width="600"
      style="top:20px"
  >
    <a-form
        :model="formInfo"
        :label-col="{ span: 3 }"
        :wrapper-col="{ span: 21 }"
        name="request"
    >

      <a-form-item label="tag" name="tag">
        <a-input
            v-model:value="formInfo['tag']"
            placeholder="请输入tag"
        />
      </a-form-item>

      <a-form-item label="method" name="method">
        <a-select
            v-model:value="formInfo['method']"
            :options="methodOptions"
            placeholder="请选择method"
        />
      </a-form-item>

      <a-form-item label="version" name="version">
        <a-input-number
            v-model:value="formInfo['version']"
            placeholder="请输入version"
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

      <a-form-item label="structure" name="structure">
        <a-textarea
            v-model:value="formInfo['structure']"
            placeholder="请输入structure"
            :rows="10"
        />
      </a-form-item>

      <a-form-item label="execQueue" name="execQueue">
        <a-input
            v-model:value="formInfo['exec_queue']"
            placeholder="exec_queue"
        />
      </a-form-item>

      <a-form-item label="debug" name="debug">
        <a-switch
            v-model:checked="formInfo['debug']"
            :checkedValue="1"
            :unCheckedValue="0"
            checked-children="是"
            un-checked-children="否"
        />
      </a-form-item>

    </a-form>
    <template #footer>
      <a-button style="margin-right: 8px" @click="visible = false">取消</a-button>
      <a-button type="primary" @click="submit">提交</a-button>
    </template>
  </a-modal>
</template>


<script lang="ts">
export default {
  name:"RequestEdit"
}
</script>

<script lang="ts" setup >
import {ref, watch} from "vue";
import api from '../../api'
import {message} from "ant-design-vue";

const visible = ref(false)

const emit = defineEmits(['update:visible', 'refresh'])


const formInfo = ref<any>({})


// 获取methods列表
const methodOptions = ref([
  {
    label:"GET",value:"GET"
  },
  {
    label:"HEAD",value:"HEAD"
  },
  {
    label:"POST",value:"POST"
  },
  {
    label:"PUT",value:"PUT"
  },
  {
    label:"DELETE",value:"DELETE"
  }

])


// 提交方法
const submit = () => {
  formInfo.value.structure = JSON.parse(JSON.stringify(formInfo.value.structure))
  api.request.update(formInfo.value).then(() => {
    message.success('编辑成功')
    visible.value = false
    emit('refresh')
  })
}

function open(info:any){
  formInfo.value = info
  visible.value = true
}

defineExpose({
  open
})
</script>

<style lang="less" scoped>

</style>
