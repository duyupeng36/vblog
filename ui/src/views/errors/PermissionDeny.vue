<script setup name="PermissionDeny">

</script>

<template>
  <a-result
      status="403"
      subtitle="Access to this resource on the server is denied."
  >
    <template #extra>
      <a-space>
        <a-button type="primary">Back</a-button>
      </a-space>
    </template>
  </a-result>
</template>

<style scoped>

</style>