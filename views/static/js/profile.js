let popupUpdateAvatar = document.getElementById("popupUpdateAvatar");
let avatarFileInput = document.getElementById("avatarFileInput");
let avatarFileName = document.getElementById("avatarFileName");
let avatarLoadingIcon = document.getElementById("avatarLoadingIcon");
let avatarFileIcon = document.getElementById("avatarFileIcon");
let userAvatarImg = document.getElementById("userAvatarImg");

function openPopupUpdateAvatar() {
  popupUpdateAvatar.classList.replace("hidden", "flex");
}

var file;
avatarFileInput.addEventListener("change", function() {
  file = avatarFileInput.files[0];
  if (file) {
    avatarFileName.innerText = file.name;
  }
})

function updateAvatarResp() {
  popupUpdateAvatar.classList.replace("flex", "hidden");
  avatarLoadingIcon.classList.replace("block", "hidden");
  avatarFileIcon.classList.replace("hidden", "block");
  avatarFileName.innerText = "Select image from your device";
  file = null;
}

function updateAvatar() {
  avatarLoadingIcon.classList.replace("hidden", "block");
  avatarFileIcon.classList.replace("block", "hidden");
  if (!file) {
    alert("Please select an image first.");
    return;
  }
  var xhr = new XMLHttpRequest();
  var formData = new FormData();
  formData.append("avatar", file);
  xhr.open("POST", "/api/user-update-avatar", true);
  xhr.withCredentials = true;
  xhr.addEventListener( "load", function() {
    const out = JSON.parse(xhr.responseText );
    const outJson = JSON.parse( out );
    if( xhr.status===200 ) {
      userAvatarImg.src = `/files/${outJson.avatarUrl}`;
      notifier.showSuccess("Image uploaded successfully.");
      updateAvatarResp();
      setTimeout(() => {
        location.reload();
      }, 3500);
    } else {
      updateAvatarResp();
      notifier.showError(outJson.error);
    }
  } );
  xhr.addEventListener( "error", function() {
    updateAvatarResp();
    notifier.showError( "Network error" );
  } );
  xhr.addEventListener( "abort", function() {
    updateAvatarResp();
    notifier.showWarning( "Upload aborted" );
  }, false );
  xhr.send(formData);
}

function cancelUpdateAvatar() {
  popupUpdateAvatar.classList.replace("flex", "hidden");
}