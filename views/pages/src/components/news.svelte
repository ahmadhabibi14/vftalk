<script lang="ts" type="module">
  import axios from 'axios';
  import Icon from 'svelte-icons-pack';
  import RiSystemLoader4Fill from 'svelte-icons-pack/ri/RiSystemLoader4Fill';
  import { NEWS_API_KEY } from 'constants/keys';
  import { formatDate, getOneMonthPastDate } from 'utils/formatter';
  import { onMount } from 'svelte';
  
  let News: any[] = [], OffsetNews: number = 1;
  let isLoadNews: boolean = false, noMoreNews: boolean = false;

  async function GetNews() {
    isLoadNews = true;
    const oneMonthPastDate: string = getOneMonthPastDate();
    const newsUrl: string = `https://newsapi.org/v2/everything?q=tech&language=en&from=${oneMonthPastDate}&sortBy=popularity&page=${OffsetNews}&pageSize=70&apiKey=${NEWS_API_KEY}`
    await axios({
      method: 'GET',
      url: newsUrl,
    }).then((response) => {
      isLoadNews = false;
      OffsetNews += 1;
      News = [...News, ...response.data.articles]
    }).catch((err) => {
      isLoadNews = false;
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
  <div class="flex justify-center text-sm">
    {#if noMoreNews}
      <div class="text-emerald-600 bg-white shadow py-2 px-5 rounded cursor-default">Tidak ada buku lagi</div>
    {:else}
      {#if isLoadNews}
        <div class="flex flex-row items-center gap-2 cursor-progress">
          <Icon size="14" src={RiSystemLoader4Fill} className="fill-emerald-600 -mt-1 animate-spin" />
          <span>Loading more</span>
        </div>
      {/if}
      {#if !isLoadNews}
        <button on:click={GetNews} class="text-emerald-600 bg-white shadow py-2 px-5 rounded hover:bg-zinc-50">Load more</button>
      {/if}
    {/if}
  </div>
{/if}
</div>