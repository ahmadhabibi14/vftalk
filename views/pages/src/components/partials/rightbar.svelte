<script lang="ts" type="module">
  import { onMount } from "svelte";

  export let fullName: string = '';
  export let avatarUrl: string = '';
  export let username: string = '';

  let url: string = avatarUrl
  onMount(()=>{
    if(!url.includes('https://')) url = `/contents${avatarUrl}`;
  })
</script>

<div class="flex flex-col gap-5 w-full h-fit">
<div class="flex flex-row justify-start items-center gap-4">
  <div class="relative flex h-12 w-12 shrink-0 overflow-hidden rounded-full shadow-md">
    <img src={url} class="aspect-square h-full w-full" alt="Avatar">
  </div>
  <div class="flex flex-col gap-1 shrink truncate">
    <h2 class="truncate">{fullName}</h2>
    <p class="text-zinc-700 text-xs truncate">@{username}</p>
  </div>
</div>
</div>