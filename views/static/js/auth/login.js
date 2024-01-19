const username = document.getElementById('username');
const password = document.getElementById('password');
const loginBtn = document.getElementById('loginBtn');
const loginTxt = document.getElementById('loginTxt');
const loginLoadingIcon = document.getElementById('loginLoadingIcon');

function RespDOM_Login() {
  loginTxt.style.display = 'block';
  loginLoadingIcon.style.display = 'none';
  loginBtn.disabled = false;
}

loginBtn.addEventListener('click', async () => {
  if (username.value === '' || password.value === '') {
    return notifier.showWarning('Please fill the correct input');
  }
  loginBtn.disabled = true;
  loginTxt.style.display = 'none';
  loginLoadingIcon.style.display = 'block';
  try {
    const resp = await fetch('/api/login', {
      method: 'POST', headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ username: username.value, password: password.value })
    });
    const respjson = await resp.json();
    if (resp.ok) {
      localStorage.setItem('username', respjson.data.username);
      notifier.showSuccess(respjson.data.message);
      RespDOM_Login();
      setTimeout(() => window.location.href = '/', 1200);
    } else {
      notifier.showError(respjson.errors);
      console.log(respjson);
      RespDOM_Login();
      return;
    }
  } catch (e) {
    console.log(e);
    notifier.showError(e);
    RespDOM_Login();
    return;
  }
});