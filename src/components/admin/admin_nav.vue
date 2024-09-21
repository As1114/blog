<template>
  <div class="admin-nav">
    <div class="slogan">
      <img :height="60" alt="" src="/public/images/logo.png">
    </div>
    <div class="function">
      <div @click="goHome">
        <IconHome></IconHome>
      </div>
      <div @click="store.logout">注销</div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import {useRouter} from "vue-router";
import {useStore} from "@/stores";

const router = useRouter()
const store = useStore()

function goHome() {
  router.push({name: 'web_home'})
}
</script>
<style>
.admin-nav {
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: space-between;

  .slogan {
    margin-left: 10px;

    img {
      object-fit: cover;
    }
  }

  .function {
    display: flex;
    margin-right: 20px;
    align-items: center;
    font-size: 18px;

    div:hover {
      cursor: pointer;
      background-color: var(--color-neutral-3);
    }
  }
}
</style>
