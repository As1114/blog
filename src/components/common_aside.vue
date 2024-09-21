<template>
  <div :class="{collapsed:store.collapsed}" class="common_aside">
    <a-layout-sider
      :collapse="store.collapsed"
      :theme="store.themeString"
      collapsible
      @collapse="collapse"
    >
      <div v-if="props.showAvatar" class="avatar">
        <a-avatar v-if="store.collapsed" :size="40"><img :src="store.userStoreInfo.avatar" alt="" loading="lazy">
        </a-avatar>
        <a-avatar v-else :size="180"><img :src="store.userStoreInfo.avatar" alt="" loading="lazy"></a-avatar>
      </div>
      <div v-if="props.showStats && !store.collapsed" class="stats">
        <a-statistic :value="123" title="Fans"/>
        <a-statistic :value="459" extra="Articles"/>
        <a-statistic :value="509" extra="Likes"/>
      </div>
      <a-menu
        v-model:open-keys="openKeys"
        v-model:selected-keys="selectedKeys"
        @menu-item-click="clickMenu"
      >
        <template v-for="item in data" :key="item.name">
          <a-menu-item v-if="item.child?.length === 0" :key="item.name">
            {{ item.title }}
            <template #icon>
              <component :is="item.icon"></component>
            </template>
          </a-menu-item>
          <a-sub-menu v-if="item.child?.length!==0 " :key="item.name">
            <template #icon>
              <component :is="item.icon"></component>
            </template>
            <template #title>{{ item.title }}</template>
            <a-menu-item v-for="sub in item.child" :key="sub.name">
              {{ sub.title }}
              <template #icon>
                <component :is="sub.icon"></component>
              </template>
            </a-menu-item>
          </a-sub-menu>
        </template>
      </a-menu>
    </a-layout-sider>
  </div>
</template>
<script lang="ts" setup>
import {useStore} from "@/stores";
import {useRoute, useRouter} from "vue-router";
import {type Component, ref} from "vue";
import {IconBook, IconHome, IconSettings, IconStorage} from "@arco-design/web-vue/es/icon";

const store = useStore()
const route = useRoute()
const router = useRouter()


const selectedKeys = ref([route.name])
let openKeys = ref([route.matched[1].name])

export interface MenuType {
  title: string
  icon: Component
  name: string
  child?: MenuType[]
}

let menuList: MenuType[] = [
  {title: "首页", icon: IconHome, name: "web_home", child: []},
  {title: "信息", icon: IconHome, name: "user_info", child: []},
  {title: "文章", icon: IconBook, name: "record", child: []},
  {title: "收藏", icon: IconStorage, name: "collect", child: []},
  {title: "后台", icon: IconSettings, name: "admin_home", child: []},
]

if (store.isGeneral) {
  menuList = [
    {title: "首页", icon: IconHome, name: "web_home", child: []},
    {title: "信息", icon: IconHome, name: "user_info", child: []},
    {title: "文章", icon: IconBook, name: "record", child: []},
    {title: "收藏", icon: IconStorage, name: "collect", child: []},
  ]
}

interface Props {
  showAvatar: boolean
  showStats: boolean
  data?: MenuType[]
}

const props = defineProps<Props>()

const {
  data = menuList
} = props

function clickMenu(name: string) {
  router.push({
    name: name,
  })
}

function collapse(collapsed: boolean) {
  store.setCollapsed(collapsed)
}
</script>
<style>
.common_aside {
  width: 200px;

  .arco-layout-sider-has-trigger {
    padding-bottom: 0;

    .avatar {
      display: flex;
      align-items: center;
      justify-content: center;
      transition: all .8s;

      img {
        object-fit: cover;
      }
    }

    .avatar:hover {
      transform: rotate(360deg);
    }

    .stats {
      display: flex;
      justify-content: space-between;
      margin: 5px;

      .arco-statistic-value {
        font-size: 14px;
      }
    }
  }
}
</style>
