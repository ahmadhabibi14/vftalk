let popupUpdateAvatar = document.getElementById("popupUpdateAvatar");
let avatarFileInput = document.getElementById("avatarFileInput");
let avatarFileName = document.getElementById("avatarFileName");
let avatarLoadingIcon = document.getElementById("avatarLoadingIcon");
let avatarFileIcon = document.getElementById("avatarFileIcon");
let userAvatarImg = document.getElementById("userAvatarImg");

function openPopupUpdateAvatar() {
  popupUpdateAvatar.classList.remove("hidden");
  popupUpdateAvatar.classList.add("flex");
}

var file;
avatarFileInput.addEventListener("change", function() {
  file = avatarFileInput.files[0];
  if (file) {
    avatarFileName.innerText = file.name;
  }
})

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
      popupUpdateAvatar.classList.replace("flex", "hidden");
      setTimeout(() => {
        location.reload();
      }, 3500);
    } else {
      popupUpdateAvatar.classList.replace("flex", "hidden");
      notifier.showError(outJson.error);
    }
  } );
  xhr.addEventListener( "error", function() {
    popupUpdateAvatar.classList.replace("flex", "hidden");
    notifier.showError( "Network error" );
  } );
  xhr.addEventListener( "abort", function() {
    popupUpdateAvatar.classList.replace("flex", "hidden");
    notifier.showWarning( "Upload aborted" );
  }, false );
  xhr.send(formData);
}

function cancelUpdateAvatar() {
  popupUpdateAvatar.classList.add("hidden");
  popupUpdateAvatar.classList.remove("flex");
}