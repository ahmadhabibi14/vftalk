<script lang="ts" type="module">
  import axios from 'axios';
  import { NEWS_API_KEY } from 'constants/keys';
  import { formatDate} from 'utils/formatter';
  import { onMount } from 'svelte';
  import type { NewsItem } from 'types/news';
  
  let News: NewsItem[] = [];

  async function GetNews() {
    const excludeDomains: string = 'yahoo.com'
    const newsUrl: string = `https://newsapi.org/v2/everything?q=tech&language=en&sortBy=popularity&excludeDomains=${excludeDomains}&pageSize=100&apiKey=${NEWS_API_KEY}`
    await axios({
      method: 'GET',
      url: newsUrl,
    }).then((response) => {
      News = [...News, ...response.data.articles]
    }).catch((err) => {
      console.log(err);
    })
  }

  onMount( async () => await GetNews() );
</script>

<div class="flex flex-col gap-4">
{#if News && News.length}
  <div class="flex flex-row items-center gap-2">
    <span class="w-4 h-px bg-zinc-200"></span>
    <p class="font-semibold">News</p>
    <span class="grow h-px bg-zinc-200"></span>
  </div>
  <div class="flex flex-col gap-3">
    {#each News as n}
      <a href={n.url} target="_blank" class="group grid grid-cols-[auto_65%] p-4 hover:bg-zinc-100 active:bg-zinc-200 rounded-md gap-4 items-stretch">
        <div class=" border border-zinc-100 h-32 overflow-hidden rounded-md">
          <img src={n.urlToImage} alt="news" class="w-full h-full object-cover duration-75 group-hover:scale-110"/>
        </div>
        <div class="flex flex-col justify-between">
          <div class="flex flex-col gap-2">
            <h4 class="line-clamp-1 font-semibold">{n.title}</h4>
            <p class="text-xs text-zinc-600 line-clamp-2">{n.description}</p>
          </div>
          <div class="flex flex-row justify-between items-center">
            <div class="flex flex-col gap-1">
              <p class="text-blue-500 text-sm">{n.author}</p>
              <p class="text-xs">{formatDate(n.publishedAt)}</p>
            </div>
            <span class="text-xs py-1.5 px-3 rounded-full font-semibold text-purple-500 bg-purple-500/20 h-fit w-fit">{n.source.name}</span>
          </div>
        </div>
      </a>
    {/each}
  </div>
{/if}
</div>