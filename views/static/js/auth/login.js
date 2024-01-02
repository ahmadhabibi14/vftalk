let username = document.getElementById("username");
let password = document.getElementById("password");
let loginBtn = document.getElementById("loginBtn");
let loginTxt = document.getElementById("loginTxt");
let loginLoadingIcon = document.getElementById("loginLoadingIcon");

loginBtn.addEventListener("click", async () => {
  if (username.value === "" || password.value === "") {
    notifier.showError("Please enter username and password");
    return;
  }
  loginBtn.disabled = true;
  loginTxt.style.display = "none";
  loginLoadingIcon.style.display = "block";

  console.log('Username:', username.value);
  console.log('Password:', password.value);
  try {
    const resp = await fetch("/api/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        username: username.value,
        password: password.value,
      }),
    });

    if (resp.ok) {
      const creds = await resp.json();

      console.log(creds);
      
      loginTxt.style.display = "block";
      loginLoadingIcon.style.display = "none";
      loginBtn.disabled = false;
      localStorage.setItem("username", successResp["data"]["username"]);

      notifier.showSuccess("Login successful");
      setTimeout(() => {
        window.location.href = "/";
      }, 1200);
    } else {
      const creds = await resp.json();
      notifier.showError('Error');
      console.log(creds);
      loginBtn.disabled = false;
      loginTxt.style.display = "block";
      loginLoadingIcon.style.display = "none";
      // username.value = "";
      // password.value = "";
      return;
    }
  } catch (e) {
    console.log(e);
    notifier.showError("Login failed");
    loginBtn.disabled = false;
    loginTxt.style.display = "block";
    loginLoadingIcon.style.display = "none";
    // username.value = "";
    // password.value = "";
  }
});