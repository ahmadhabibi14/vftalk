<script lang="ts" type="module">
  import Icon from 'svelte-icons-pack';
  import RiSystemLoader4Fill from 'svelte-icons-pack/ri/RiSystemLoader4Fill';
  import Inputbox from '../inputbox.svelte';
  import Growl from '../growl.svelte';
  import axios from 'axios';

  let email: string = '';
  let username: string = '';
  let full_name: string = '';
  let password: string = '';
  let isSubmitted: boolean = false;

  let growl: any = Growl;

  async function Login() {
    if (username == '' || password == '' 
      || email == '' || full_name == '') return growl.showWarning('Please fill all fields');
    isSubmitted = true;
    await axios({
			method: 'post',
			url: '/api/register',
			data: { email, username, full_name, password},
			headers: { 'Content-Type': 'application/json' },
		}).then((res) => {
      isSubmitted = false;
      console.log(res.data);
      growl.showSuccess(res.data.data.message);
      setTimeout(() => window.location.href = '/', 1200);
    }).catch((err) => {
      isSubmitted = false;
      console.log(err.response);
      growl.showError(err.response.data.errors)
    })
  }
</script>

<Growl bind:this={growl}/>

<div class="w-[500px] h-fit bg-white shadow-md p-5 rounded-md flex flex-col gap-4">
  <h1 class="font-bold text-2xl">Create an account !</h1>
  <div class="flex flex-col gap-4">
    <Inputbox
      id="email"
      label="Email"
      type="email"
      placeholder="gojosatoru@example.com"
      bind:value={email}
    />
    <Inputbox
      id="username"
      label="Username"
      type="text"
      placeholder="gojosatoru98"
      bind:value={username}
    />
    <Inputbox
      id="full_name"
      label="Fullname"
      type="text"
      placeholder="Gojo Satoru"
      bind:value={full_name}
    />
    <Inputbox
      id="password"
      label="Password"
      type="password"
      placeholder="password123"
      bind:value={password}
    />
    <button on:click={Login} class="bg-emerald-700 hover:bg-emerald-600 font-semibold justify-center text-white flex items-center rounded-md py-2 w-full">
      {#if isSubmitted}
        <Icon size="23" src={RiSystemLoader4Fill} className="fill-white animate-spin" />
      {/if}
      {#if !isSubmitted}
        <span>Register</span>
      {/if}
    </button>
  </div>

  <div class="flex flex-row gap-2 items-center">
    <span class="h-px grow bg-zinc-400"></span>
    <span>or</span>
    <span class="h-px grow bg-zinc-400"></span>
  </div>

  <div class="flex flex-col gap-3">
    <a class="flex flex-row gap-3 justify-center items-center font-semibold py-2 rounded-md border-zinc-200 border hover:bg-zinc-100" href="/oauth/google" >
      <img src="/icons/oauth/google.svg" class="w-5 h-auto" alt="Google" />
      <span>Continue with Google</span>
    </a>
    <a class="flex flex-row gap-3 justify-center items-center font-semibold py-2 rounded-md border-zinc-200 border hover:bg-zinc-100"  href="/oauth/facebook">
      <img src="/icons/oauth/facebook.svg" class="w-5 h-auto" alt="facebook" />
      <span>Continue with Facebook</span>
    </a>
  </div>

  <span class="text-sm text-center font-semibold">Already have account? <a href="/login" class="text-emerald-700 hover:underline">Login</a></span>
</div>