<script lang="ts" type="module">
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import RiBusinessSendPlane2Fill from 'svelte-icons-pack/ri/RiBusinessSendPlane2Fill';
  import RiSystemTimeLine from 'svelte-icons-pack/ri/RiSystemTimeLine';
  import RiSystemCheckDoubleLine from 'svelte-icons-pack/ri/RiSystemCheckDoubleLine';
  import RiSystemErrorWarningFill from 'svelte-icons-pack/ri/RiSystemErrorWarningFill';
  import RiSystemLoader2Line from 'svelte-icons-pack/ri/RiSystemLoader2Line';
  import { onMount } from 'svelte';
  import type { Chat, ChatIn, ChatSendState } from 'types/chats';
  import { ChatState } from 'constants/chats';
  import { timeFormat } from 'utils/date_formatter';

  export let username: string = '';
  let Chats: Chat[] = [];
  let ChatsCount: number = 0;
  let ChatSendState: ChatSendState = {
    state: ChatState.Load,
    index: ChatsCount,
  };

  const url = new URL(window.location.href);
  let socketURL = `ws://${window.location.host}/api/room-general`;
  if ((url.protocol).includes('s')) socketURL = `wss://${window.location.host}/api/room-general`;
  let socket: WebSocket = new WebSocket(socketURL);

  onMount( () => localStorage.setItem('username', username) );

  socket.onopen = () => {
    console.log('Connected to the room chat')
  }

  socket.onmessage = (e) => {
    const data: any = JSON.parse(e.data);
    const jsonObj: Chat = data as Chat;
    if (jsonObj.sender !== username) {
      if (jsonObj.type === 'text') jsonObj.content = (jsonObj.content).replace(/\r\n|\n|\r/gm, '<br />');
      Chats = [...Chats, jsonObj];
      setTimeout(() => {
        const ELM: HTMLElement = document.getElementById(`CH_${ChatsCount}`);
        ELM.scrollIntoView({ behavior: 'smooth' });
        ChatsCount += 1;
      }, 300);
    }
  }

  socket.onclose = (e) => {
    // TODO: use growl, and popup for reconnecting
    console.log('Closed:', e);
  }

  socket.onerror = (e) => {
    console.log('Error:', e);
  }

  let message: string = ``;
  let isSending: boolean = false;

  function SendMessage(): void {
    if (!message) return;
    isSending = true;
    const payload: ChatIn = {
      type: 'text',
      content: message
    }
    const chatOut: Chat = {
      sender: username,
      type: 'text',
      content: message.replace(/\r\n|\n|\r/gm, '<br />'),
      datetime: new Date(),
    }
    Chats = [...Chats, chatOut];
    ChatSendState.state = ChatState.Load;
    setTimeout(() => {
      const ELM: HTMLElement = document.getElementById(`CH_${ChatsCount}`);
      ELM.scrollIntoView({ behavior: 'smooth' });
      ChatsCount += 1;
    }, 300);
    try {
      message = '';
      socket.send( JSON.stringify(payload) );
      ChatSendState = {
        state: ChatState.Sent,
        index: ChatsCount-1
      }
      isSending = false;
    } catch (e) {
      console.log('ERROR sending message',e);
      ChatSendState = {
        state: ChatState.Error,
        index: ChatsCount-1
      }
      isSending = false;
    }
  }
</script>

<div class="min-h-full h-full relative w-full">
  <div class="flex flex-col gap-2 min-h-full h-full w-full px-5 pt-5 pb-0 mb-20">
    {#if Chats && Chats.length}
      {#each Chats as c, idx}
        {#if c.sender === username}
          <div class="flex w-full justify-end text-sm" id={`CH_${idx}`}>
            <div class="w-fit max-w-[70%] rounded-tr-md rounded-s-md bg-emerald-700 text-white p-2 gap-1 flex flex-col">
              <p class="text-ellipsis break-all">{@html c.content}</p>
              <div class="flex flex-row gap-1 justify-end items-center text-xs font-light text-emerald-100">
                {#if ChatSendState.state === ChatState.Load && ChatSendState.index === idx}
                  <Icon size="12" src={RiSystemTimeLine} className="fill-emerald-200"/>
                {:else if ChatSendState.state === ChatState.Error && ChatSendState.index === idx}
                  <Icon size="12" src={RiSystemErrorWarningFill} className="fill-emerald-200"/>
                {:else}
                  <Icon size="12" src={RiSystemCheckDoubleLine} className="fill-emerald-200"/>
                {/if}
                <span>{timeFormat(c.datetime)}</span>
              </div>
            </div>
          </div>
        {:else if c.sender === 'system' && c.type !== 'text'}
          <div class="flex w-full justify-center text-sm" id={`CH_${idx}`}>
            <div class="w-fit max-w-full mx-5 bg-emerald-500/20 text-emerald-700 border border-emerald-700 rounded-full font-semibold py-1 px-3 flex flex-row">
              <p class="text-ellipsis break-all text-[10px]">{c.content}</p>
            </div>
          </div>
        {:else}
          <div class="flex w-full justify-start text-sm" id={`CH_${idx}`}>
            <div class="w-fit max-w-[70%] rounded-e-md rounded-tl-md bg-zinc-100 p-2 gap-1 flex flex-col">
              <span class="text-[10px] text-sky-600 font-semibold w-fit">@{c.sender}</span>
              <p class="text-ellipsis break-all">{@html c.content}</p>              
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
      disabled={isSending}
      on:click|preventDefault={SendMessage}
      class="bg-emerald-700 disabled:bg-emerald-600 hover:bg-emerald-600 cursor-pointer text-white h-full py-auto px-3 rounded-md">
      {#if !isSending}
        <Icon size="23" src={RiBusinessSendPlane2Fill} className="fill-white"/>
      {/if}
      {#if isSending}
        <Icon size="23" src={RiSystemLoader2Line} className="fill-white animate-spin"/>
      {/if}
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