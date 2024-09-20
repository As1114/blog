<template>
  <div class="web_article_list">
    <a-empty>
      <template #image>
        <a-image :preview="false" src="/public/images/页面为空.png"></a-image>
      </template>
      No data, please reload!
    </a-empty>
  </div>
</template>
<script lang="ts" setup>
</script>
<style>
.web_article_list {
  background-color: var(--color-bg-1);
  
}
</style>
