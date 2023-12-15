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
  if (file) avatarFileName.innerText = file.name;
})

function updateAvatarResp() {
  popupUpdateAvatar.classList.replace("flex", "hidden");
  avatarLoadingIcon.classList.replace("block", "hidden");
  avatarFileIcon.classList.replace("hidden", "block");
  avatarFileName.innerText = "Select image from your device";
  file = null;
}

function updateAvatar() {
  if (!file) return notifier.showError("Please select an image first.");
  avatarLoadingIcon.classList.replace("hidden", "block");
  avatarFileIcon.classList.replace("block", "hidden");
  var xhr = new XMLHttpRequest();
  var formData = new FormData();
  formData.append("avatar", file);
  xhr.open("POST", "/api/user-update-avatar", true);
  xhr.withCredentials = true;
  xhr.addEventListener( "load", function() {
    const out = JSON.parse(xhr.responseText );
    const outJson = JSON.parse( out );
    if( xhr.status===200 ) {
      userAvatarImg.src = `/files${outJson.avatarUrl}`;
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

let popupUpdateProfile = document.getElementById("popupUpdateProfile");
let fullNameInput = document.getElementById("fullNameInput");
let locationInput = document.getElementById("locationInput");
let websiteInput = document.getElementById("websiteInput");

function updateProfileResp() {
  fullNameInput.value = "";
  locationInput.value = "";
  websiteInput.value = "";
}

function openPopupUpdateProfile() {
  popupUpdateProfile.classList.replace("hidden", "flex");
}

function cancelUpdateProfile() {
  popupUpdateProfile.classList.replace("flex", "hidden");
}

function updateProfile() {
  if (!fullNameInput.value) return notifier.showError("Please enter your full name.");
  if (!locationInput.value) return notifier.showError("Please enter your location.");
  if (!websiteInput.value) return notifier.showError("Please enter your website.");
  console.log('Update profile');
  notifier.showInfo("Updating profile...");
}