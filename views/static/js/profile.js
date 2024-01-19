const popupUpdateAvatar = document.getElementById('popupUpdateAvatar');
const avatarFileInput = document.getElementById('avatarFileInput');
const avatarFileName = document.getElementById('avatarFileName');
const avatarLoadingIcon = document.getElementById('avatarLoadingIcon');
const avatarFileIcon = document.getElementById('avatarFileIcon');
const userAvatarImg = document.getElementById('userAvatarImg');

var file;
avatarFileInput.addEventListener('change', function() {
  file = avatarFileInput.files[0];
  if (file) avatarFileName.innerText = file.name;
})

function RespDOM_updateAvatar() {
  popupUpdateAvatar.classList.replace('flex', 'hidden');
  avatarLoadingIcon.classList.replace('block', 'hidden');
  avatarFileIcon.classList.replace('hidden', 'block');
  avatarFileName.innerText = 'Select image from your device';
  file = null;
  avatarFileInput.files = null;
}

const openPopupUpdateAvatar = () => popupUpdateAvatar.classList.replace('hidden', 'flex');
const cancelUpdateAvatar = () => {
  avatarFileInput.files = null; file = null; popupUpdateAvatar.classList.replace('flex', 'hidden');
  avatarFileName.innerText = 'Select image from your device';
}

function updateAvatar() {
  if (!file) return notifier.showError('Please select an image first.');
  avatarLoadingIcon.classList.replace('hidden', 'block');
  avatarFileIcon.classList.replace('block', 'hidden');
  var xhr = new XMLHttpRequest(), formData = new FormData();
  formData.append('avatar', file);
  xhr.open('POST', '/api/user-update-avatar', true);
  xhr.withCredentials = true;
  xhr.addEventListener( 'load', function() {
    const outJson = JSON.parse(xhr.responseText );
    if( xhr.status===200 ) {
      userAvatarImg.src = `/files${outJson.data.avatar}`;
      notifier.showSuccess(outJson.data.message);
      RespDOM_updateAvatar();
      setTimeout(() => location.reload(), 3500);
    } else {
      console.log(outJson);
      RespDOM_updateAvatar();
      notifier.showError(outJson.errors+': '+outJson.data);
    }
  } );
  xhr.addEventListener( 'error', function() {
    RespDOM_updateAvatar();
    notifier.showError( 'Network error' );
  } );
  xhr.addEventListener( 'abort', function() {
    RespDOM_updateAvatar();
    notifier.showWarning( 'Upload aborted' );
  }, false );
  xhr.send(formData);
}

const popupUpdateProfile = document.getElementById('popupUpdateProfile');
const fullNameInput = document.getElementById('fullNameInput');
const locationInput = document.getElementById('locationInput');
const websiteInput = document.getElementById('websiteInput');
const updateProfileLoadingIcon = document.getElementById('updateProfileLoadingIcon');
const updateProfileButtonText = document.getElementById('updateProfileButtonText');

function RespDOM_UpdateProfile() {
  updateProfileLoadingIcon.classList.replace('block', 'hidden');
  updateProfileButtonText.classList.replace('hidden', 'block');
}

const openPopupUpdateProfile = () => popupUpdateProfile.classList.replace('hidden', 'flex');
const cancelUpdateProfile = () => popupUpdateProfile.classList.replace('flex', 'hidden');

async function updateProfile() {
  if (!fullNameInput.value) return notifier.showError('Please enter your full name.');
  if (!locationInput.value) return notifier.showError('Please enter your location.');
  if (!websiteInput.value) return notifier.showError('Please enter your website.');

  updateProfileLoadingIcon.classList.replace('hidden', 'block');
  updateProfileButtonText.classList.replace('block', 'hidden');
  try {
    const resp = await fetch('/api/user-update-profile', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        full_name: fullNameInput.value,
        location: locationInput.value,
        website: websiteInput.value
      }),
    });
    const respjson = await resp.json();
    if (resp.ok) {
      RespDOM_UpdateProfile()
      notifier.showSuccess(respjson.data);
      setTimeout(()=> window.location.reload(), 1300);
    } else {
      notifier.showError(respjson.errors+': '+respjson.data);
      console.log(respjson);
      RespDOM_UpdateProfile()
      return;
    }
  } catch (e) {
    console.log(e);
    notifier.showError('Error: ', e);
    RespDOM_UpdateProfile();
    return;
  }
}