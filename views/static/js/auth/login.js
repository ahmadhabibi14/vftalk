let username = document.getElementById("username");
let password = document.getElementById("password");
let loginBtn = document.getElementById("loginBtn");
let loginTxt = document.getElementById("loginTxt");
let loginLoadingIcon = document.getElementById("loginLoadingIcon");

loginBtn.addEventListener("click", async () => {
  if (username.value === "" || password.value === "") {
    alert("Please enter username and password");
    return;
  }
  loginBtn.disabled = true;
  loginTxt.style.display = "none";
  loginLoadingIcon.style.display = "block";

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
			const successResp = await JSON.parse(creds);
      
      loginTxt.style.display = "block";
      loginLoadingIcon.style.display = "none";
      loginBtn.disabled = false;
      localStorage.setItem("username", successResp["username"]);
      window.location.href = "/";
    } else {
      const creds = await resp.json();
      const errResp = await JSON.parse(creds);
      alert(errResp["error"]);
      loginBtn.disabled = false;
      loginTxt.style.display = "block";
      loginLoadingIcon.style.display = "none";
      username.value = "";
      password.value = "";
      return;
    }
  } catch (e) {
    alert("Login failed");
    loginBtn.disabled = false;
    loginTxt.style.display = "block";
    loginLoadingIcon.style.display = "none";
    username.value = "";
    password.value = "";
  }
});
