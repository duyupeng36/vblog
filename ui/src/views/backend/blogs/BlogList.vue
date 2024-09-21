<script setup name="BlogList">

import {BLOG_LIST} from "@/api/blog.js"
import {reactive} from "vue";
import {Message} from "@arco-design/web-vue";

const columns = reactive([
  {title: "ID", dataIndex: "id"},
  {title: "标题", dataIndex: "title"},
])

const data = reactive({
  items: [],
  total: 0
})

async function getBlogList(queryBlogsRequest) {
  try {
    let result = await BLOG_LIST(queryBlogsRequest)

    data.items = result.data.items;
    data.total = result.data.total;

  } catch (error) {
    Message.error(error.message)
  }
}


getBlogList({
  page_size:20,
  page_number:1,
})

</script>

<template>
  <div class="blogs-container">
    <a-table :columns="columns" :data="data.items" />
  </div>
</template>

<style scoped>

</style>