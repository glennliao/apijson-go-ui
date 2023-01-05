<template>

  <div style="background-color: #ececec; padding: 8px">
    <a-row :gutter="[16,16]">
      <a-col :span="8" v-for="item in list" :key="item.alias">
        <a-card  size="small" :bordered="false">
          <a-descriptions :title="item.alias">
            <template #extra>
              <a-button size="small" @click="edit(item)" type="primary">edit</a-button>
            </template>
            <a-descriptions-item label="name">{{item.name}}</a-descriptions-item>
            <a-descriptions-item label="row_key">{{item.row_key}}</a-descriptions-item>
            <a-descriptions-item label="debug">{{item.debug}}</a-descriptions-item>
            <a-descriptions-item label="detail">
              {{item.detail}}
            </a-descriptions-item>
            <a-descriptions-item label="created_at" :span="2">{{item.created_at}}</a-descriptions-item>

            <a-descriptions-item  v-for="method in methods" :label="method" :span="3">
             <div style="display: flex;flex-wrap: wrap">
               <a-tag v-for="roleItem in item[method]" :key="roleItem.value" size="small" color="green">{{roleItem.value}}</a-tag>
             </div>
            </a-descriptions-item>
          </a-descriptions>
        </a-card>
      </a-col>
    </a-row>
  </div>


  <Edit ref="EditModal" @refresh="loadList"/>
  <Add ref="AddModal" @refresh="loadList"/>

  <div class="fixed" style="bottom:20px;right:20px" >
    <a-button type="primary" shape="circle" @click="add">
      <template #icon><PlusOutlined /></template>
    </a-button>
  </div>


</template>


<script lang="ts" setup>
import {reactive, ref} from "vue";
import api from '../../api'
import Edit from "./Edit.vue";
import Add from "./Add.vue";
import {methods} from "../../consts";
import { PlusOutlined } from '@ant-design/icons-vue';

// 获取表格列表数据
const list = ref<any[]>([])


const loadList = () => {
  api.access.list({}).then((data: any) => {
    if (data?.List?.length) {
      list.value = data.List.map((item:any)=>{

        methods.forEach(method=>{
          item[method] = JSON.parse(item[method]).map((roleItem:string)=>{
            return {
              value:roleItem,
            }
          })
        })

        return {
          ...item,
        }
      })
    }
  })
}

loadList()


const EditModal = ref()
function edit(item:any){
  // @ts-ignore
  EditModal.value.open(item)
}

const AddModal = ref()
function add(){
  // @ts-ignore
  AddModal.value.open()
}


</script>


<style scoped>

</style>
