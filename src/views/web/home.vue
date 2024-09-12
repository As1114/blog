<template>
  <div class="web_home">
    <web_template>
      <template #content>
        <a-carousel
          :autoPlay="true"
          :style="{width: '100%',height: '240px',}"
          animation-name="card"
          indicator-position="outer"
          show-arrow="hover"
        >
          <a-carousel-item :style="{ width: '60%' }">
            <a-image :preview="false" src="/public/images/classroom.jpg"></a-image>
          </a-carousel-item>
          <a-carousel-item :style="{ width: '60%' }">
            <a-image :preview="false" src="/public/images/classroom.jpg"></a-image>
          </a-carousel-item>
          <a-carousel-item :style="{ width: '60%' }">
            <a-image :preview="false" src="/public/images/classroom.jpg"></a-image>
          </a-carousel-item>
        </a-carousel>
        <web_article_list></web_article_list>
      </template>
      <template #aside>
      </template>
    </web_template>
  </div>
</template>
<script lang="ts" setup>
import Web_template from "@/components/web/web_template.vue";
import Web_article_list from "@/components/web/web_article_list.vue";
</script>
<style>
.web_home {
  width: 100%;

  .web_template {
    width: 100%;

    .web_template_content {
      .arco-carousel {
        img {
          width: 100%;
          height: 100%;
          object-fit: cover;
        }
      }
    }
  }
}
</style>
