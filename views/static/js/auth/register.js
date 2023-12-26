let email = document.getElementById("email");
let fullname = document.getElementById("fullname");
let username = document.getElementById("username");
let password = document.getElementById("password");
let registerBtn = document.getElementById("registerBtn");
let registerTxt = document.getElementById("registerTxt");
let registerLoadingIcon = document.getElementById("registerLoadingIcon");

registerBtn.addEventListener("click", async () => {
  if (username.value === "" || password.value === "" || email.value === "" || fullname.value === "") {
    notifier.showError("Please fill your credentials");
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
        full_name: fullname.value
      }),
    });

    console.log('Resp: ', resp);

    if (resp.ok) {
      const creds = await resp.json();
      console.log('Creds:', creds)
			const successResp = await JSON.parse(creds);
      console.log('Success Resp:', successResp)
      
      registerTxt.style.display = "block";
      registerLoadingIcon.style.display = "none";
      registerBtn.disabled = false;
      localStorage.setItem("username", successResp["username"]);
      notifier.showSuccess("Register successfully!")
      setTimeout(() => {
        window.location.href = "/";
      }, 1500);
    } else {
      const creds = await resp.json();
      console.log(creds);
      const errResp = await JSON.parse(creds);
      notifier.showError(errResp["errors"]);
      registerBtn.disabled = false;
      registerTxt.style.display = "block";
      registerLoadingIcon.style.display = "none";
      username.value = "";
      password.value = "";
      return;
    }
  } catch (e) {
    notifier.showError("Register failed");
    registerBtn.disabled = false;
    registerTxt.style.display = "block";
    registerLoadingIcon.style.display = "none";
  }
});