<script>
  import { onMount } from 'svelte';
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import AiOutlineEye from 'svelte-icons-pack/ai/AiOutlineEye';
  import AiOutlineEyeInvisible from 'svelte-icons-pack/ai/AiOutlineEyeInvisible';

  export let id;
  export let value;
  export let label;
  export let type = 'text'; // email, phone, password, text, number, radio, checkbox
  export let placeholder = '';
  export let required = false;
  export let values = [];
  export let selectshow = '';
  /**
    * @typedef {Object} radio
    * @property {string} label
    * @property {string} value
    */
  /**
    * @type {Array<radio>}
    */
  export let radios = [];
  let isShowPassword = false;
  let inputElm;
  
  onMount(() => {
    if (inputElm) inputElm.type = type
  })

  function toggleShowPassword() {
    isShowPassword = !isShowPassword;
    if (isShowPassword) inputElm.type = 'text';
    else inputElm.type = 'password';
  }

  function onRadio(event) { value = event.currentTarget.value }
</script>

{#if type === 'radio'}
  <div class="flex flex-col gap-2">
    <span>{label}</span>
    <div class="flex flex-row gap-3 items-center">
      {#each radios as r}
        <div class="flex flex-row gap-1 items-center cursor-pointer">
          <input
            class="cursor-pointer"
            id={r.label}
            checked={value === r.value}
            on:change={onRadio}
            type="radio"
            name={label}
            value={r.value}
          />
          <label for={r.label}>{r.label}</label>
        </div>
      {/each}
    </div>
  </div>
{:else if type === 'select'}
  <div class="flex flex-col gap-2">
    <label for={id} class="text-sm text-zinc-600 ml-2">
      <span>{label}</span>
      {#if required}
        <span class="text-red-500 !text-lg"> *</span>
      {/if}
    </label>
    <select
      name={id}
      id={id}
      bind:value={value}
      class="rounded-md w-full border border-zinc-300 py-2 caret-emerald-700 focus:border-emerald-700 focus:outline focus:outline-emerald-700 px-4"
    >
    {#if values && values.length}
      {#each values as v}
        <option value={v}>{v[selectshow] || v[id] || v}</option>
      {/each}
    {/if}
    </select>
  </div>
{:else}
<div class={`flex flex-col gap-2 ${type === 'password' ? 'relative' : ''}`} >
  <label for={id} class="text-sm text-zinc-600 ml-2">
    <span>{label}</span>
    {#if required}
      <span class="text-red-500"> *</span>
    {/if}
  </label>
  <input
    pattern={type === 'tel' ? "[0-9]{3}-[0-9]{3}-[0-9]{4}" : ""}
    class={`rounded-md w-full border border-zinc-300 py-2 caret-emerald-700 focus:border-emerald-700 focus:outline focus:outline-emerald-700 ${type === 'password' ? 'pl-4 pr-10' : 'px-4'}`}
    bind:value={value} {id} bind:this={inputElm} {placeholder}/>
  {#if type === 'password'}
    <button
      class="absolute right-2 top-9"
      title="Show/Hide Password"
      on:click={toggleShowPassword}
    >
      {#if !isShowPassword}
        <Icon className="fill-zinc-700 hover:fill-emerald-700" size="24" src={AiOutlineEye}/>
      {/if}
      {#if isShowPassword}
        <Icon className="fill-zinc-700 hover:fill-emerald-700" size="24" src={AiOutlineEyeInvisible}/>
      {/if}
    </button>
  {/if}
</div>
{/if}