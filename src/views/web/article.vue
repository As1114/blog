<template>
  <div class="article">
    <web_template>
      <template #content>
        <div class="article_detail">
          <div class="head">
            <div class="title">
              {{ data.title }}
            </div>
            <div class="data">
              <span><icon-user/>{{ data.user_name }}</span>
              <span>
            <icon-clock-circle/>{{ dateFormat(data.created_at) }}
          </span>
              <span>
            <icon-eye/>{{ data.look_count }}次浏览
          </span>
              <span>
            <icon-message/>{{ data.comment_count }}条评论
          </span>
              <span>
            <icon-sort/>{{ data.category }}
          </span>
            </div>
          </div>
          <MdPreview v-model="data.content" :editor-id="mdId" :theme="store.themeString"></MdPreview>
        </div>
      </template>
      <template #aside>
      </template>
    </web_template>
  </div>
</template>

<script lang="ts" setup>
import 'md-editor-v3/lib/preview.css';
import {reactive, ref, watch} from "vue";
import {useStore} from "@/stores";
import {articleDetail, type articleDetailType} from "@/api/article";
import {useRoute, useRouter} from "vue-router";
import {Message} from "@arco-design/web-vue";
import {dateFormat} from "@/utils/date";
import Web_template from "@/components/web/web_template.vue";
import {MdPreview} from 'md-editor-v3';

const route = useRoute()
const router = useRouter()
const store = useStore()
const id = ref<string>(route.params.id as string)
const mdId = "mdpreview"

const data = reactive<articleDetailType>({
    abstract: "",
    category: "",
    collects_count: 0,
    comment_count: 0,
    content: "",
    cover_id: 0,
    cover_url: "",
    created_at: "",
    digg_count: 0,
    id: "",
    look_count: 0,
    title: "",
    updated_at: "",
    user_avatar: "",
    user_id: 0,
    user_name: ""
  }
)

async function getData() {
  let res = await articleDetail(id.value)
  if (res.code) {
    await router.push({
      name: "article_notfound"
    })
    Message.info(res.msg)
    return
  }
  Object.assign(data, res.data)
}

getData()

watch(() => route.params, () => {
  if (route.params.id) {
    id.value = route.params.id as string
    getData()
  }
}, {immediate: true, deep: true})
</script>

<style>
.article {
  .web_template_content {
    .article_detail {
      .head {
        max-height: 200px;

        .title {
          font-size: 36px;
          margin-bottom: 10px;
        }

        .data {
          span {
            font-size: 18px;
            margin-right: 10px;
          }
        }
      }

      .md-editor {
        .md-editor-preview-wrapper {
          padding: 0;

          .md-editor-preview {
            padding: 10px;
          }
        }
      }
    }
  }

}
</style>
