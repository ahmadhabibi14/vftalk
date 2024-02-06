<script lang="ts" type="module">
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import AiOutlineHome from 'svelte-icons-pack/ai/AiOutlineHome';
  import AiFillHome from 'svelte-icons-pack/ai/AiFillHome';
  import RiUserAccountCircleLine from 'svelte-icons-pack/ri/RiUserAccountCircleLine';
  import RiUserAccountCircleFill from 'svelte-icons-pack/ri/RiUserAccountCircleFill';
  import RiSystemSearchLine from "svelte-icons-pack/ri/RiSystemSearchLine";
  import RiSystemSearchFill from "svelte-icons-pack/ri/RiSystemSearchFill";
  import RiCommunicationChat3Line from 'svelte-icons-pack/ri/RiCommunicationChat3Line';
  import RiSystemSettings3Line from 'svelte-icons-pack/ri/RiSystemSettings3Line';
  import { onMount } from 'svelte';

  let path: string = window.location.pathname;
  onMount(() => path = window.location.pathname);
</script>

<div class="w-full h-full flex flex-col justify-between">
  <nav class="flex flex-col gap-3">
    <a class={`flex flex-row gap-3 py-2 pl-4 pr-5 rounded-full w-fit ${path === '/' ? 'bg-emerald-500/20 text-emerald-800 hover:bg-emerald-500/20' : 'hover:bg-zinc-100'}`} href="/">
      {#if path !== '/'}
        <Icon className="fill-current" size="24" src={AiOutlineHome}/>
      {/if}
      {#if path === '/'}
        <Icon className="fill-current" size="24" src={AiFillHome}/>
      {/if}
      <span>Home</span>
    </a>
    <a class={`flex flex-row gap-3 py-2 pl-4 pr-5 rounded-full w-fit ${path === '/profile' ? 'bg-emerald-500/20 text-emerald-800 hover:bg-emerald-500/20' : 'hover:bg-zinc-100'}`} href="/profile">
      {#if path !== '/profile'}
        <Icon className="fill-current" size="24" src={RiUserAccountCircleLine}/>
      {/if}
      {#if path === '/profile'}
        <Icon className="fill-current" size="24" src={RiUserAccountCircleFill}/>
      {/if}
      <span>Profile</span>
    </a>
    <a class={`flex flex-row gap-3 py-2 pl-4 pr-5 rounded-full w-fit ${path === '/explore' ? 'bg-emerald-500/20 text-emerald-800 hover:bg-emerald-500/20' : 'hover:bg-zinc-100'}`} href="/explore">
      {#if path !== '/explore'}
        <Icon className="fill-current" size="24" src={RiSystemSearchLine}/>
      {/if}
      {#if path === '/explore'}
        <Icon className="fill-current" size="24" src={RiSystemSearchFill}/>
      {/if}
      <span>Explore</span>
    </a>
    <a class={`flex flex-row gap-3 py-2 pl-4 pr-5 rounded-full w-fit ${path === '/chats' ? 'bg-emerald-500/20 text-emerald-800 hover:bg-emerald-500/20' : 'hover:bg-zinc-100'}`} href="/chats">
      <Icon className="fill-current" size="23" src={RiCommunicationChat3Line}/>
      <span>Chats</span>
    </a>
  </nav>
  <div>
    <a class={`flex flex-row gap-3 py-2 pl-4 pr-5 rounded-full w-fit ${path === '/setting' ? 'bg-emerald-500/20 text-emerald-800 hover:bg-emerald-500/20' : 'hover:bg-zinc-100'}`} href="/setting">
      <Icon className="fill-current" size="24" src={RiSystemSettings3Line}/>
      <span>Setting</span>
    </a>
  </div>
</div>