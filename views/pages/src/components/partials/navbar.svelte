<script lang="ts" type="module">
  const path:string = window.location.pathname;

  type navbar = {
    link: string;
    title: string;
    blank: boolean;
  }
  
  const navbars:navbar[] = [
    {link: '/about', title: 'About', blank: false},
    {link: '/contact', title: 'Contact', blank: false},
    {link: 'https://github.com/ahmadhabibi14/chat-app', title: 'Source', blank: true},
    {link: '/login', title: 'Login', blank: false},
  ]
</script>

{#each navbars as nav}
  <a href={nav.link} class="hover:text-emerald-600 relative h-full w-fit flex items-center" target={nav.blank ? '_blank' : ''}>
    <span>{nav.title}</span>
    {#if path === nav.link}
      <span class="absolute h-[3px] w-full flex justify-center bottom-5">
        <span class="bg-emerald-700 rounded-full h-[3px] w-[50px]"></span>
      </span>
    {/if}
  </a>
{/each}