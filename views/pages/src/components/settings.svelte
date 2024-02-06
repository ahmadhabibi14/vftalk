<script lang="ts" type="module">
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import RiSystemLogoutBoxRLine from 'svelte-icons-pack/ri/RiSystemLogoutBoxRLine';
  import RiSystemLoader4Fill from 'svelte-icons-pack/ri/RiSystemLoader4Fill';

  let showLogoutPopup: boolean = false;
  function logout(): void {
    showLogoutPopup = true;
    let expiry: Date = new Date(), cookieName = 'auth';
    expiry.setTime(expiry.getTime() - 3600);
    document.cookie = cookieName + "=; expires=" + expiry.toUTCString() + "; path=/"
    setTimeout(()=> window.location.href = '/', 1200);
  }
</script>

{#if showLogoutPopup}
  <div class="top-0 right-0 bottom-0 left-0 h-full w-full bg-zinc-950/30 flex justify-center fixed z-[99999]">
    <div class="bg-white w-[300px] flex flex-col justify-center items-center gap-2 p-5 rounded shadow-lg h-fit mt-24">
      <Icon size="40" src={RiSystemLoader4Fill} className="fill-emerald-700 -mt-1 animate-spin" />
      <span>Logging out...</span>
    </div>
  </div>
{/if}

<div>
  <button on:click={logout} class="py-1.5 pl-3 pr-4 rounded-md bg-red-500 hover:bg-red-400 text-white flex flex-row items-center gap-2">
    <Icon className="fill-current" size="15" src={RiSystemLogoutBoxRLine}/>
    <span>Logout</span>
  </button>
</div>