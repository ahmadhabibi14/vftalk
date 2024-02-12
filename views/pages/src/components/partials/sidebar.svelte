<script lang="ts" type="module">
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import AiFillHome from 'svelte-icons-pack/ai/AiFillHome';
  import RiUserAccountCircleFill from 'svelte-icons-pack/ri/RiUserAccountCircleFill';
  import RiSystemSearch2Fill from 'svelte-icons-pack/ri/RiSystemSearch2Fill';
  import RiCommunicationChatSmile2Fill from 'svelte-icons-pack/ri/RiCommunicationChatSmile2Fill';
  import RiSystemSettings3Fill from 'svelte-icons-pack/ri/RiSystemSettings3Fill';
  import RiSystemAppsFill from 'svelte-icons-pack/ri/RiSystemAppsFill';
  import { onMount } from 'svelte';

  let path: string = window.location.pathname;
  onMount(() => path = window.location.pathname);

  const year: number = new Date().getFullYear();
</script>

<nav class="flex flex-col gap-3 font-semibold">
  <a class={`flex flex-row items-center gap-3 py-2 pl-4 pr-5 rounded-md w-full ${path === '/' ? 'bg-zinc-100 hover:bg-zinc-100' : 'hover:bg-zinc-100 active:bg-zinc-200'}`} href="/">
    <Icon className="fill-blue-500" size="26" src={AiFillHome}/>
    <span>Home</span>
  </a>
  <a class={`flex flex-row items-center gap-3 py-2 pl-4 pr-5 rounded-md w-full ${path === '/profile' ? 'bg-zinc-100 hover:bg-zinc-100' : 'hover:bg-zinc-100 active:bg-zinc-200'}`} href="/profile">
    <Icon className="fill-emerald-500" size="26" src={RiUserAccountCircleFill}/>
    <span>Profile</span>
  </a>
  <a class={`flex flex-row items-center gap-3 py-2 pl-4 pr-5 rounded-md w-full ${path === '/explore' ? 'bg-zinc-100 hover:bg-zinc-100' : 'hover:bg-zinc-100 active:bg-zinc-200'}`} href="/explore">
    <Icon className="fill-purple-500" size="26" src={RiSystemSearch2Fill}/>
    <span>Explore</span>
  </a>
  <a class={`flex flex-row items-center gap-3 py-2 pl-4 pr-5 rounded-md w-full ${path === '/chats' ? 'bg-zinc-100 hover:bg-zinc-100' : 'hover:bg-zinc-100 active:bg-zinc-200'}`} href="/chats">
    <Icon className="fill-amber-500" size="26" src={RiCommunicationChatSmile2Fill}/>
    <span>Chats</span>
  </a>
  <a class={`flex flex-row items-center gap-3 py-2 pl-4 pr-5 rounded-md w-full ${path === '/rooms' ? 'bg-zinc-100 hover:bg-zinc-100' : 'hover:bg-zinc-100 active:bg-zinc-200'}`} href="/rooms">
    <Icon className="fill-sky-500" size="26" src={RiSystemAppsFill}/>
    <span>Rooms</span>
  </a>
  <a class={`flex flex-row items-center gap-3 py-2 pl-4 pr-5 rounded-md w-full ${path === '/setting' ? 'bg-zinc-100 hover:bg-zinc-100' : 'hover:bg-zinc-100 active:bg-zinc-200'}`} href="/setting">
    <Icon className="fill-zinc-700" size="26" src={RiSystemSettings3Fill}/>
    <span>Setting</span>
  </a>
</nav>