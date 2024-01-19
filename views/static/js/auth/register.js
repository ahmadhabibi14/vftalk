const email = document.getElementById('email');
const fullname = document.getElementById('fullname');
const username = document.getElementById('username');
const password = document.getElementById('password');
const registerBtn = document.getElementById('registerBtn');
const registerTxt = document.getElementById('registerTxt');
const registerLoadingIcon = document.getElementById('registerLoadingIcon');

function RespDOM_Register() {
  registerTxt.style.display = 'block';
  registerLoadingIcon.style.display = 'none';
  registerBtn.disabled = false;
}

registerBtn.addEventListener('click', async () => {
  if (username.value === '' || password.value === '' || email.value === '' || fullname.value === '') {
    return notifier.showWarning('Please fill with correct inputs');
  }
  registerBtn.disabled = true;
  registerTxt.style.display = 'none';
  registerLoadingIcon.style.display = 'block';
  try {
    const resp = await fetch('/api/register', {
      method: 'POST', headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        email: email.value,
        password: password.value,
        username: username.value,
        full_name: fullname.value
      }),
    });
    const respjson = await resp.json();
    if (resp.ok) {
      localStorage.setItem('username', respjson.data.username);
      notifier.showSuccess(respjson.data.message);
      RespDOM_Register();
      setTimeout(() => window.location.href = '/', 1200);
    } else {
      notifier.showError(respjson.errors);
      console.log(respjson);
      RespDOM_Register();
      return;
    }
  } catch (e) {
    console.log(e);
    notifier.showError(e);
    RespDOM_Register();
    return;
  }
});