let email = document.getElementById("email");
let fullname = document.getElementById("fullname");
let username = document.getElementById("username");
let password = document.getElementById("password");
let registerBtn = document.getElementById("registerBtn");
let registerTxt = document.getElementById("registerTxt");
let registerLoadingIcon = document.getElementById("registerLoadingIcon");

registerBtn.addEventListener("click", async () => {
  if (username.value === "" || password.value === "" || email.value === "" || fullname.value === "") {
    alert("Please enter username and password");
    return;
  }
  registerBtn.disabled = true;
  registerTxt.style.display = "none";
  registerLoadingIcon.style.display = "block";

  try {
    const resp = await fetch("/api/register", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        email: email.value,
        password: password.value,
        username: username.value,
        fullname: fullname.value
      }),
    });

    if (resp.ok) {
      const creds = await resp.json();
			const successResp = await JSON.parse(creds);
      
      registerTxt.style.display = "block";
      registerLoadingIcon.style.display = "none";
      registerBtn.disabled = false;
      localStorage.setItem("username", successResp["username"]);
      window.location.href = "/";
    } else {
      const creds = await resp.json();
      const errResp = await JSON.parse(creds);
      alert(errResp["error"]);
      registerBtn.disabled = false;
      registerTxt.style.display = "block";
      registerLoadingIcon.style.display = "none";
      username.value = "";
      password.value = "";
      return;
    }
  } catch (e) {
    alert("Login failed");
    registerBtn.disabled = false;
    registerTxt.style.display = "block";
    registerLoadingIcon.style.display = "none";
    username.value = "";
    password.value = "";
  }
});