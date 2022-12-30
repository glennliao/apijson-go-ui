<template>
  <a-table
      :columns="columns"
      :data-source="dataSource"
      size="middle"
      :pagination="pagination"
      @change="handleTableChange"
      :expandedRowKeys="curExpandedRowKeys"
      row-key="id"
      :expandIcon="expandIcon"
      :expandIconAsCell="false"
  >
    <template #bodyCell="{ column, text, record }">

      <template v-if="column.dataIndex === 'structure'">
        <div>
          <a class="text-1890ff"
             @click="handleExpand(record.id)">{{ curExpandedRowKeys[0] === record.id ? '关闭' : '查看' }}</a>
        </div>
      </template>

      <template v-else-if="column.dataIndex === 'operation'">
        <div class="editable-row-operations">
          <a class="text-1890ff" @click="edit(record)">编辑</a>
        </div>
      </template>
    </template>

    <template #expandedRowRender="{ record }">
      <div style="height: 300px; overflow-y: scroll;">
        <pre>{{ JSON.stringify(JSON.parse(record.structure), null, 2) }}</pre>
      </div>
    </template>
  </a-table>

  <Edit ref="EditModal"  @refresh="lookupRequestList"/>
</template>

<script lang="ts" setup>
import {reactive, ref} from "vue";
import Edit from './Edit.vue'
import api from '../../api'

const columns = [
  {
    title: 'tag',
    dataIndex: 'tag',
    width: 150
  },
  {
    title: 'method',
    dataIndex: 'method',
    width: 100
  },
  {
    title: 'version',
    dataIndex: 'version',
    width: 100
  },
  {
    title: 'debug',
    dataIndex: 'debug',
    width: 80
  },
  {
    title: 'detail',
    dataIndex: 'detail',
    ellipsis: true,
  },
  {
    title: 'structure',
    dataIndex: 'structure',
    width: 100
  },
  {
    title:"execQueue",
    dataIndex: "exec_queue",
  },
  {
    title: '操作',
    dataIndex: 'operation',
    width: 100
  },
];

// 表格分页配置
const pagination = reactive({
  current: 1,
  total: 0,
  pageSize: 10,
  pageSizeOptions: ["10", "20", "50"],
  responsive: true,
  showSizeChanger: true,
})

// 获取表格列表数据
const dataSource = ref<any[]>([])
const lookupRequestList = () => {
  api.request.list({_pn: pagination.current, _ps: pagination.pageSize}).then((data: any) => {
    pagination.total = data.total
    if (data?.List?.length) {
      dataSource.value = data.List
    }
  })
}
lookupRequestList()

// table @change 分页、排序、筛选变化时触发
const handleTableChange = ({current, pageSize}: any) => {
  pagination.current = current
  pagination.pageSize = pageSize
  lookupRequestList()
}

// edit 点击列表项编辑时触发
const visible = ref(false)
const curInfo = reactive<any>({
  handelType: 'add',
})
const EditModal = ref()
const edit = (info: any) => {
  EditModal.value.open(info)
}



// 隐藏表格展开按钮
const expandIcon = () => {
  return null;
}

// 自定义展开
const curExpandedRowKeys = ref<any[]>([])
const handleExpand = (id: string) => {
  if (curExpandedRowKeys.value.length > 0) {
    let index = curExpandedRowKeys.value.findIndex(item => item === id);
    if (index !== -1) {
      curExpandedRowKeys.value.splice(index, 1);
    } else {
      curExpandedRowKeys.value.splice(0, curExpandedRowKeys.value.length);
      curExpandedRowKeys.value.push(id);
    }
  } else {
    curExpandedRowKeys.value.push(id);
  }
}
</script>

<style lang="less" scoped>
.editable-row-operations a {
  margin-right: 8px;
}

</style>
