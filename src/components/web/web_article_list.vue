<template>
  <div class="web_article_list">
    <template v-if="data.count">
      <div class="article_list">
        <div v-for="article in data.list" :key="article.id" class="item" @click="enter(article.id)">
          <a-image :description="article.title" :src="article.cover_url" alt="" fit="cover"></a-image>
          <div class="content">
            <div class="top">
              <span> <icon-user/>{{ article.user_name }}</span>
              <span><icon-clock-circle/>{{ dateFormat(article.created_at) }}</span>
              <span><icon-eye/>{{ article.look_count }}次浏览</span>
              <span><icon-message/>{{ article.comment_count }}条评论</span>
              <span><icon-sort/>{{ article.category }}</span>
            </div>
            <div class="bottom">
              <a-typography-paragraph :ellipsis="{ rows: 2, css: true }">{{
                  article.abstract
                }}
              </a-typography-paragraph>
            </div>
          </div>
        </div>
      </div>
      <div class="page">
        <a-pagination v-model:current="params.page" :total="data.count" show-total @change="pageChange"></a-pagination>
      </div>
    </template>
    <template v-else>
      <a-empty>
        <template #image>
          <a-image :preview="false" src="/public/images/页面为空.png"></a-image>
        </template>
        No data, please reload!
      </a-empty>
    </template>
  </div>
</template>
<script lang="ts" setup>
import {articleList, type articleListType} from "@/api/article";
import type {listDataType, paramsType} from "@/api";
import {reactive} from "vue";
import {useRouter} from "vue-router";
import {dateFormat} from "@/utils/date";

const router = useRouter()

const params = reactive<paramsType>({
  page: 1,
  page_size: 10,
})
const data = reactive<listDataType<articleListType>>({
  list: [],
  count: 0
})

async function getData(p?: paramsType) {
  if (p) {
    Object.assign(params, p)
  }
  let res = await articleList(params)
  data.list = res.data.list
  data.count = res.data.count
}

function pageChange() {
  getData()
}

getData()


function enter(id: string) {
  router.push(`/article/${id}`)
}
</script>
<style>
.web_article_list {
  background-color: var(--color-bg-1);
  margin-top: 10px;

  .article_list {
    .item {
      margin-bottom: 10px;
      border: 1px solid var(--color-neutral-2);
      background-color: var(--color-bg-3);

      .arco-image {
        width: 100%;

        img {
          width: 100%;
          height: 180px;
        }

        .arco-image-footer {
          position: absolute;
          top: 0;

          .arco-image-footer-caption {
            display: flex;
            justify-content: center;

            .arco-image-footer-caption-description {
              font-size: 18px;
              font-weight: 800;
            }
          }
        }
      }

      .content {
        font-size: 16px;

        .top {
          margin: 5px 10px;

          span {
            margin-right: 10px;
          }
        }

        .bottom {
          margin: 5px 10px;
        }
      }


    }

    .item:nth-child(3n+2) {
      display: flex;
      align-items: center;
      justify-content: center;
      width: 100%;

      .arco-image {
        width: 50%;

        img {
          width: 100%;
        }


      }

      .content {
        width: 50%;

        .top {
          display: grid;
          grid-template-columns: repeat(2, 1fr);
        }
      }
    }

    .item:nth-child(3n+3) {
      display: flex;
      align-items: center;
      justify-content: center;
      flex-direction: row-reverse;
      width: 100%;

      .arco-image {
        width: 50%;

        img {
          width: 100%;
        }


      }

      .content {
        width: 50%;

        .top {
          display: grid;
          grid-template-columns: repeat(2, 1fr);
        }

      }

    }

  }

  .page {
    display: flex;
    justify-content: center;
  }
}
</style>
