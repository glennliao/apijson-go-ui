<template>
  <a-card
      :title="role"
      size="small"
      style="margin-bottom: 10px;"
  >
    <div class="content">
      <div class="in">
        <div class="title flex">
          <span>in</span>
        </div>

        <div class="px-2 pt-2">

            <div style="display: flex; justify-content: space-between;padding: 4px" v-for="(item) in roleData.inList"
                 :key="item">
              <div style="width: 100px;">
                <a-checkbox @change="inChange(item.value)" :checked="!item.checked">{{ item.label }}</a-checkbox>
              </div>
              <div class="flex-1">
                <a-select
                    v-model:value="fieldsGet[role].in[item.value]"
                    :options="operatorOptions"
                    mode="multiple"
                    :maxTagCount="4"
                    style="width: 200px;"
                    :disabled="item.checked"
                    placeholder="请选择操作符"
                    @change="operatorChange(item, $event)"
                />
              </div>

          </div>
        </div>
      </div>

      <div class="out">
        <div class="title">
          <span>out</span>
        </div>
        <div class="px-2 pt-2" style="padding: 4px">
          <a-form-item v-for="item in roleData.outList" style="margin: 5px 0; width: 100%">
            <a-checkbox
                @change="outChange(item.value)"
                :checked="!item.checked"
            >
              {{ item.label }}
            </a-checkbox>
          </a-form-item>
        </div>
      </div>
    </div>
  </a-card>
</template>

<script setup lang="ts">

import {computed, shallowReactive,ref} from "vue";

const props = defineProps({
  inOptions: {
    type: Array,
    default() {
      return []
    }
  },
  outOptions: {
    type: Array,
    default() {
      return []
    }
  },
  fieldsGet: {
    type: Object,
    default() {
      return {}
    }
  },
  handelType: {
    type: String,
    default: 'add'
  },
  role: {},
})
const emit = defineEmits(['change'])

const operatorOptions = ref([
  {
    label:"等于",value:"="
  },
  {
    label:"任意",value:"*"
  },
  {
    value: "$%",
    label: "右Like",
  }
])



const inOptions = computed(() => {
  let arr: any[] = []
  props.inOptions?.forEach((item: any) => {
    arr.push({
      ...item
    })
  })

  return arr
})
const outOptions = computed(() => {
  let arr: any[] = []
  props.outOptions?.forEach((item: any) => {
    arr.push({
      ...item
    })
  })

  return arr
})
const roleData = shallowReactive({
  inList: inOptions.value,
  outList: outOptions.value,
})

// in @change 事件触发 反选当前点击的多选框并改变对应的select禁用状态
const inChange = (value: string) => {
  roleData.inList = roleData.inList.map((item: any) => {
    if (item.value === value) {
      item.checked = !item.checked
    }
    return item
  })
}

//  out @change 事件触发 反选当前点击的多选框  处理所选参数并向外抛出事件
const outChange = (value: string) => {
  roleData.outList = roleData.outList.map((item: any) => {
    if (item.value === value) {
      item.checked = !item.checked
    }
    return item
  })

  let out: any = {}
  if (roleData.outList?.length) {
    roleData.outList.filter((item: any) => !item.checked).forEach((item: any) => {
      out[item.value] = ""
    })
  }
  emit('change', {out, key: props.role, type: 'out'})
}

// operator @change 事件触发 判断所选最后一项是否为 "*" 然后抛出事件进行相应处理
const operatorChange = (item: any, valueArr: any) => {
  const len = valueArr?.length - 1

  const flag = valueArr?.length && valueArr[len] === '*'

  emit('change', {flag, key: props.role, item, type: 'in', valueArr})
}
</script>


<style lang="less" scoped>
.content {
  border: 1px solid #f0f0f0;
  display: flex;

  .in {
    width: 50%;
    border-right: 1px solid #f0f0f0;

    :deep(.ant-form-item-control) {
      max-width: 100%;
    }
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
