<template>
  <a-tabs
      v-model:activeKey="activeKey"
      :tab-position="mode"
      :style="{ height: '100%' }"
  >
    <a-tab-pane key="Access" tab="Access"><Access/></a-tab-pane>
    <a-tab-pane key="Request" tab="Request"><Request/></a-tab-pane>
    <a-tab-pane key="Function" tab="Function"><Function/></a-tab-pane>
  </a-tabs>
</template>

<script lang="ts" setup>
import {  ref } from 'vue';
import type { TabsProps } from 'ant-design-vue';
import Access from '../components/Access/index.vue'
import Request from '../components/Request/index.vue'
import Function from '../components/Function/index.vue'
const mode = ref<TabsProps['tabPosition']>('left');
const activeKey = ref("Access");
</script>

<style scoped>

</style>