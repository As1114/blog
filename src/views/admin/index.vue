<template>
  <div class="admin_index">
    <div class="admin_index_header">
      <admin_nav></admin_nav>
    </div>
    <div class="admin_index_content">
      <common_aside :class="{collapsed:store.collapsed}" :data="menuList" :show-avatar="false"
                    :show-stats="false"></common_aside>
      <div class="content">
        <router-view v-slot="{Component}">
          <transition mode="out-in" name="fade">
            <component :is="Component"></component>
          </transition>
        </router-view>
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import Admin_nav from "@/components/admin/admin_nav.vue";
import Common_aside, {type MenuType} from "@/components/common_aside.vue";
import {useStore} from "@/stores";
import {type Component} from "vue";
import {
  IconBook,
  IconCalendarClock,
  IconEar,
  IconHome,
  IconMessageBanned,
  IconOrderedList,
  IconPenFill,
  IconSettings,
  IconSortAscending,
  IconStorage,
  IconUser,
  IconUserGroup
} from "@arco-design/web-vue/es/icon";

const store = useStore()

const menuList: MenuType[] = [
  {title: "首页", icon: IconHome, name: "admin_home", child: []},
  {
    title: "文章管理", icon: IconBook, name: "article_manage", child: [
      {title: "文章列表", icon: IconStorage, name: "article_list"},
      {title: "文章反馈", icon: IconEar, name: "article_feedback"},
      {title: "分类列表", icon: IconSortAscending, name: "category_list"},
      {title: "评论列表", icon: IconPenFill, name: "comment_list"},
    ]
  },
  {
    title: "用户管理", icon: IconUserGroup, name: "user_manage", child: [
      {title: "用户列表", icon: IconUser, name: "user_list"},
      {title: "黑名单", icon: IconMessageBanned, name: "back_list"},
      {title: "消息列表", icon: IconOrderedList, name: "message_list"},
    ]
  },
  {
    title: "系统管理", icon: IconSettings, name: "system_manage", child: [
      {title: "日志列表", icon: IconCalendarClock, name: "log_list"},
      {title: "推荐列表", icon: IconCalendarClock, name: "advertisement_list"},
    ]
  },
]
</script>
<style lang="scss">
.admin_index {
  display: flex;
  height: 100vh;
  flex-direction: column;

  .admin_index_content {
    height: calc(100vh - 60px);
    display: flex;
    transition: all .3s;

    .common_aside {
      width: 200px;

      .arco-layout-sider {
        height: calc(100vh - 108px);
      }

    }

    .common_aside.collapsed {
      width: 48px;

      & ~ .content {
        width: calc(100% - 48px);
      }
    }

    .content {
      width: calc(100% - 200px);
      overflow-y: auto;
      overflow-x: hidden;
    }
  }
}

.fade-leave-to {
  opacity: 0;
  transform: translateX(30px);
}

.fade-enter-active {
  transform: translateX(-30px);
  opacity: 0;
}

.fade-enter-to {
  transform: translateX(0px);
  opacity: 1;
}

.fade-leave-active, .fade-enter-active {
  transition: all 0.3s ease-out;
}
</style>
