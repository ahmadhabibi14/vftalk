<script lang="ts" type="module">
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import RiBusinessSendPlane2Fill from 'svelte-icons-pack/ri/RiBusinessSendPlane2Fill';
  import { onMount } from 'svelte';
  import type { Chat } from 'types/chats';
  import { timeFormat } from '../../utils/formatter';

  export let username: string = '';
  let Chats: Chat[] = [];

  const url = new URL(window.location.href);
  let socketURL = `ws://${window.location.host}/api/room-general`;
  if ((url.protocol).includes('s')) socketURL = `wss://${window.location.host}/api/room-general`;
  let socket: WebSocket = new WebSocket(socketURL);

  onMount(() => {
    localStorage.setItem('username', username);
  })

  socket.onopen = () => {
    console.log('Connected to the room chat')
  }

  socket.onmessage = (e) => {
    const data: any = JSON.parse(e.data);
    const jsonObj: Chat = data as Chat;
    Chats = [...Chats, jsonObj];
  }

  socket.onclose = (e) => {
    console.log('Closed:', e)
  }

  let message: string = ``;

  function SendMessage() {
    const payload = {
      type: 'text',
      content: message
    }
    try {
      message = '';
      console.log(payload);
      socket.send( JSON.stringify(payload) );
    } catch (e) {
      console.log(e);
    }
  }
</script>

<div class="min-h-full h-full relative w-full">
  <div class="flex flex-col gap-2 min-h-full h-full w-full p-6">
    {#if Chats && Chats.length}
      {#each Chats as c}
        {#if c.sender === username}
          <div class="flex w-full justify-end text-sm">
            <div class="w-fit max-w-[70%] rounded-tr-md rounded-s-md bg-emerald-700 text-white p-2 flex flex-col">
              <p class="text-ellipsis break-all">{@html (c.content).replace(/\r\n|\n|\r/gm, '<br />')}</p>
              <span class="text-end text-[10px] font-light text-emerald-100">{timeFormat(c.datetime)}</span>
            </div>
          </div>
        {:else if c.sender === 'system' && c.type !== 'text'}
          <div class="flex w-full justify-center text-sm">
            <div class="w-fit max-w-full mx-5 bg-emerald-500/20 text-emerald-700 border border-emerald-700 rounded-full font-semibold py-1 px-3 flex flex-row">
              <p class="text-ellipsis break-all text-[10px]">{c.content}</p>
            </div>
          </div>
        {:else}
          <div class="flex w-full justify-start text-sm">
            <div class="w-fit max-w-[70%] rounded-e-md rounded-tl-md bg-zinc-100 p-2 flex flex-col">
              <span class="text-[10px] text-sky-600 font-semibold w-fit">@{c.sender}</span>
              <p class="text-ellipsis break-all">{@html (c.content).replace(/\r\n|\n|\r/gm, '<br />')}</p>
              <span class="text-start text-[10px] font-light text-zinc-700">{timeFormat(c.datetime)}</span>
            </div>
          </div>
        {/if}
      {/each}
    {/if}
  </div>
  <div class="-ml-[2px] w-[700px] p-3 fixed flex flex-row gap-3 items-end bottom-0 h-16 bg-white border-t-[2px] border-x-[2px] border-zinc-100">
    <textarea
      class="scrollbar-hide resize-none grow rounded-md border border-zinc-200 p-2 w-full h-full caret-emerald-700 focus:border-emerald-700 focus:outline focus:outline-emerald-700"
      name="message"
      id="message"
      rows="10"
      placeholder="Message..."
      bind:value={message}
    ></textarea>
    <button
      on:click|preventDefault={SendMessage}
      class="bg-emerald-700 hover:bg-emerald-600 cursor-pointer text-white h-full py-auto px-3 rounded-md">
      <Icon size="23" src={RiBusinessSendPlane2Fill} className="fill-white"/>
    </button>
  </div>
</div>

<style lang="postcss">
  .scrollbar-hide::-webkit-scrollbar {
    visibility: hidden;
    @apply bg-transparent w-2;
  }

  .scrollbar-hide::-webkit-scrollbar-thumb {
    visibility: hidden;
    @apply bg-zinc-200 w-2 rounded-full cursor-pointer;
  }

  .scrollbar-hide::-webkit-scrollbar:hover {
    visibility: initial;
  }
  .scrollbar-hide::-webkit-scrollbar-thumb:hover {
    visibility: initial;
    @apply bg-zinc-300 cursor-pointer;
  }

  .scrollbar-hide:hover::-webkit-scrollbar {
    visibility: initial;
  }
  .scrollbar-hide:hover::-webkit-scrollbar-thumb {
    visibility: initial;
  }
</style>