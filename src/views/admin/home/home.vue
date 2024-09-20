<template>
  <div class="admin_home">
    <div class="home_left">
      <div class="welcome">
        欢迎回来！{{ }}
      </div>
      <div class="web_data">
        <a-statistic :value="125670" show-group-separator title="Downloads"/>
        <a-statistic :value="40509" extra="Comments" show-group-separator/>
        <a-statistic :value="125670" show-group-separator title="Downloads"/>
        <a-statistic :value="40509" extra="Comments" show-group-separator/>
      </div>
    </div>
    <div class="home_right"></div>
  </div>
</template>
<script lang="ts" setup>
import {useStore} from "@/stores";

const store = useStore()
</script>
<style>
.admin_home {
  display: flex;
  margin: 20px;

  .home_left {
    width: 70%;
    height: calc(100vh - 100px);
    font-size: 24px;
    margin-right: 20px;

    .welcome {
      margin: 20px;
    }

    .web_data {
      display: flex;
      align-items: center;
      justify-content: space-around;
      margin: 0 20px;
      height: 200px;
    }
  }

  .home_right {
    width: 30%;
    height: calc(100vh - 100px);
  }
}
</style>
