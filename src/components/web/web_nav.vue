<template>
  <login v-model:visible="lvisible"></login>
  <div class="web_nav">
    <div class="nav_slogan">
      <img :height="80" alt="" src="/public/images/logo.png">
    </div>
    <div class="nav_function">
      <web_search></web_search>
      <div class="nav_login">
        <div v-if="!store.isLogin" type="text" @click="updatevisible">登录</div>
        <div v-else type="text" @click="store.logout">注销</div>
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import Web_search from "@/components/web/web_search.vue";
import Login from "@/components/web/login/login.vue";
import {ref} from "vue";
import {useStore} from "@/stores";

const lvisible = ref(false)
const store = useStore()

function updatevisible() {
  lvisible.value = true
}
</script>
<style>
.web_nav {
  width: 100%;
  height: 80px;
  position: fixed;
  top: 0;
  align-items: center;
  display: flex;
  z-index: 100;
  justify-content: space-between;
  font-size: 18px;

  .nav_slogan {
    margin-left: 10px;

    img {
      object-fit: cover;
    }
  }

  .nav_function {
    display: flex;
    margin-right: 20px;
    align-items: center;

    div:hover {
      cursor: pointer;
      background-color: var(--color-neutral-3);
    }
  }
}
</style>
