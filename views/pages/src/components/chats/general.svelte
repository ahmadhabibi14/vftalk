<script lang="ts" type="module">
  import { onMount } from "svelte";

  export let username: string = '';

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
    const data = JSON.parse(e.data);
    console.log(data);
  }

  socket.onclose = (e) => {
    console.log('Closed:', e)
  }

  let message: string = '';

  function SendMessage() {
    const payload = {
      type: 'text',
      content: message
    }
    try {
      socket.send( JSON.stringify(payload) );
      console.log('Payload:', payload)
    } catch (e) {
      console.log(e);
    }
  }
</script>

<div class="min-h-full h-full relative">
  <div class="flex flex-col gap-2 min-h-full h-full p-6">
    <div class="flex w-full justify-start text-sm">
      <div class="w-fit max-w-[70%] rounded-e-md rounded-tl-md bg-zinc-100 p-3 flex flex-col gap-2">
        <span class="text-xs text-sky-600 font-semibold w-fit">@ahmadhabibi14</span>
        <p class="text-ellipsis break-all">
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
        </p>
      </div>
    </div>
  </div>
  <div class="-ml-[2px] w-[700px] p-6 fixed flex flex-row gap-3 items-stretch bottom-0 h-20 max-h-32 bg-white border-t-[2px] border-x-[2px] border-zinc-100">
    <textarea name="message" id="message" cols="30" rows="10" placeholder="" bind:value={message}></textarea>
    <button on:click|preventDefault={SendMessage} class="bg-lime-600 hover:bg-lime-500 cursor-pointer text-white py-2 px-5 rounded-md">Send</button>
  </div>
</div>