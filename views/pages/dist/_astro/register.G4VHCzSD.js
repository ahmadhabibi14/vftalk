import{S as $e,i as we,s as ke,b as V,a as U,c as S,d as b,e as k,f as H,g as h,h as y,j as ee,k as q,l as M,m as I,n as R,o as te,p as c,q as ye,r as B,t as $,u as C,v as Ce,w as E,x as Te}from"./index.vfLk9hgJ.js";import{I as je}from"./Icon._Bt0AYzA.js";import{R as Le}from"./RiSystemLoader4Fill.thdniLxq.js";import{G as ge,I as O}from"./growl._iKFL-fu.js";import{a as ze}from"./axios.L6U4YIEh.js";import"./each.-oqiv04n.js";/* empty css                       */function me(a){let n,s;return n=new je({props:{size:"23",src:Le,className:"fill-white animate-spin"}}),{c(){S(n.$$.fragment)},l(e){H(n.$$.fragment,e)},m(e,u){R(n,e,u),s=!0},i(e){s||($(n.$$.fragment,e),s=!0)},o(e){C(n.$$.fragment,e),s=!1},d(e){E(n,e)}}}function de(a){let n,s="Register";return{c(){n=k("span"),n.textContent=s},l(e){n=y(e,"SPAN",{"data-svelte-h":!0}),q(n)!=="svelte-17kzaxj"&&(n.textContent=s)},m(e,u){te(e,n,u)},d(e){e&&M(n)}}}function Ie(a){let n,s,e,u,w="Create an account !",v,i,f,G,D,d,N,P,g,l,W,x,J,K,_,A,Q,T,ne='<span class="h-px grow bg-zinc-400"></span> <span>or</span> <span class="h-px grow bg-zinc-400"></span>',X,j,ae='<a class="flex flex-row gap-3 justify-center items-center font-semibold py-2 rounded-md border-zinc-200 border hover:bg-zinc-100" href="/oauth/google"><img src="/icons/oauth/google.svg" class="w-5 h-auto" alt="Google"/> <span>Register with Google</span></a> <a class="flex flex-row gap-3 justify-center items-center font-semibold py-2 rounded-md border-zinc-200 border hover:bg-zinc-100" href="/oauth/facebook"><img src="/icons/oauth/facebook.svg" class="w-5 h-auto" alt="facebook"/> <span>Register with Facebook</span></a>',Y,L,se='Already have account? <a href="/login" class="text-emerald-700 hover:underline">Login</a>',F,Z,le,_e={};n=new ge({props:_e}),a[7](n);function be(t){a[8](t)}let oe={id:"email",label:"Email",type:"email",placeholder:"gojosatoru@example.com"};a[0]!==void 0&&(oe.value=a[0]),f=new O({props:oe}),V.push(()=>U(f,"value",be));function he(t){a[9](t)}let ie={id:"username",label:"Username",type:"text",placeholder:"gojosatoru98"};a[1]!==void 0&&(ie.value=a[1]),d=new O({props:ie}),V.push(()=>U(d,"value",he));function ve(t){a[10](t)}let re={id:"full_name",label:"Fullname",type:"text",placeholder:"Gojo Satoru"};a[2]!==void 0&&(re.value=a[2]),g=new O({props:re}),V.push(()=>U(g,"value",ve));function xe(t){a[11](t)}let ue={id:"password",label:"Password",type:"password",placeholder:"password123"};a[3]!==void 0&&(ue.value=a[3]),x=new O({props:ue}),V.push(()=>U(x,"value",xe));let r=a[4]&&me(),p=!a[4]&&de();return{c(){S(n.$$.fragment),s=b(),e=k("div"),u=k("h1"),u.textContent=w,v=b(),i=k("div"),S(f.$$.fragment),D=b(),S(d.$$.fragment),P=b(),S(g.$$.fragment),W=b(),S(x.$$.fragment),K=b(),_=k("button"),r&&r.c(),A=b(),p&&p.c(),Q=b(),T=k("div"),T.innerHTML=ne,X=b(),j=k("div"),j.innerHTML=ae,Y=b(),L=k("span"),L.innerHTML=se,this.h()},l(t){H(n.$$.fragment,t),s=h(t),e=y(t,"DIV",{class:!0});var o=ee(e);u=y(o,"H1",{class:!0,"data-svelte-h":!0}),q(u)!=="svelte-1f2nrsn"&&(u.textContent=w),v=h(o),i=y(o,"DIV",{class:!0});var m=ee(i);H(f.$$.fragment,m),D=h(m),H(d.$$.fragment,m),P=h(m),H(g.$$.fragment,m),W=h(m),H(x.$$.fragment,m),K=h(m),_=y(m,"BUTTON",{class:!0});var z=ee(_);r&&r.l(z),A=h(z),p&&p.l(z),z.forEach(M),m.forEach(M),Q=h(o),T=y(o,"DIV",{class:!0,"data-svelte-h":!0}),q(T)!=="svelte-54ypqb"&&(T.innerHTML=ne),X=h(o),j=y(o,"DIV",{class:!0,"data-svelte-h":!0}),q(j)!=="svelte-1mteh80"&&(j.innerHTML=ae),Y=h(o),L=y(o,"SPAN",{class:!0,"data-svelte-h":!0}),q(L)!=="svelte-uff996"&&(L.innerHTML=se),o.forEach(M),this.h()},h(){I(u,"class","font-bold text-2xl"),I(_,"class","bg-emerald-700 hover:bg-emerald-600 font-semibold justify-center text-white flex items-center rounded-md py-2 w-full"),I(i,"class","flex flex-col gap-4"),I(T,"class","flex flex-row gap-2 items-center"),I(j,"class","flex flex-col gap-3"),I(L,"class","text-sm text-center font-semibold"),I(e,"class","w-[500px] h-fit bg-white shadow-md p-5 rounded-md flex flex-col gap-4")},m(t,o){R(n,t,o),te(t,s,o),te(t,e,o),c(e,u),c(e,v),c(e,i),R(f,i,null),c(i,D),R(d,i,null),c(i,P),R(g,i,null),c(i,W),R(x,i,null),c(i,K),c(i,_),r&&r.m(_,null),c(_,A),p&&p.m(_,null),c(e,Q),c(e,T),c(e,X),c(e,j),c(e,Y),c(e,L),F=!0,Z||(le=ye(_,"click",a[6]),Z=!0)},p(t,[o]){const m={};n.$set(m);const z={};!G&&o&1&&(G=!0,z.value=t[0],B(()=>G=!1)),f.$set(z);const fe={};!N&&o&2&&(N=!0,fe.value=t[1],B(()=>N=!1)),d.$set(fe);const pe={};!l&&o&4&&(l=!0,pe.value=t[2],B(()=>l=!1)),g.$set(pe);const ce={};!J&&o&8&&(J=!0,ce.value=t[3],B(()=>J=!1)),x.$set(ce),t[4]?r?o&16&&$(r,1):(r=me(),r.c(),$(r,1),r.m(_,A)):r&&(Te(),C(r,1,1,()=>{r=null}),Ce()),t[4]?p&&(p.d(1),p=null):p||(p=de(),p.c(),p.m(_,null))},i(t){F||($(n.$$.fragment,t),$(f.$$.fragment,t),$(d.$$.fragment,t),$(g.$$.fragment,t),$(x.$$.fragment,t),$(r),F=!0)},o(t){C(n.$$.fragment,t),C(f.$$.fragment,t),C(d.$$.fragment,t),C(g.$$.fragment,t),C(x.$$.fragment,t),C(r),F=!1},d(t){t&&(M(s),M(e)),a[7](null),E(n,t),E(f),E(d),E(g),E(x),r&&r.d(),p&&p.d(),Z=!1,le()}}}function Se(a,n,s){let e="",u="",w="",v="",i=!1,f=ge;async function G(){if(u==""||v==""||e==""||w=="")return f.showWarning("Please fill all fields");s(4,i=!0),await ze({method:"post",url:"/api/register",data:{email:e,username:u,full_name:w,password:v},headers:{"Content-Type":"application/json"}}).then(l=>{s(4,i=!1),console.log(l.data),f.showSuccess(l.data.data.message),setTimeout(()=>window.location.href="/",1200)}).catch(l=>{s(4,i=!1),console.log(l.response),f.showError(l.response.data.errors)})}function D(l){V[l?"unshift":"push"](()=>{f=l,s(5,f)})}function d(l){e=l,s(0,e)}function N(l){u=l,s(1,u)}function P(l){w=l,s(2,w)}function g(l){v=l,s(3,v)}return[e,u,w,v,i,f,G,D,d,N,P,g]}class Pe extends $e{constructor(n){super(),we(this,n,Se,Ie,ke,{})}}export{Pe as default};
