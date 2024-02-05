<script lang="ts" type="module">
  import { onMount } from "svelte";

  export let userlists = ``;

  onMount(()=>{
    console.log(userlists);
  })
</script>

<div>
  {userlists}
</div>