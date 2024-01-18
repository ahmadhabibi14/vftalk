let popupUpdateAvatar = document.getElementById('popupUpdateAvatar');
let avatarFileInput = document.getElementById('avatarFileInput');
let avatarFileName = document.getElementById('avatarFileName');
let avatarLoadingIcon = document.getElementById('avatarLoadingIcon');
let avatarFileIcon = document.getElementById('avatarFileIcon');
let userAvatarImg = document.getElementById('userAvatarImg');

function openPopupUpdateAvatar() {
  popupUpdateAvatar.classList.replace('hidden', 'flex');
}

var file;
avatarFileInput.addEventListener('change', function() {
  file = avatarFileInput.files[0];
  if (file) avatarFileName.innerText = file.name;
})

function updateAvatarResp() {
  popupUpdateAvatar.classList.replace('flex', 'hidden');
  avatarLoadingIcon.classList.replace('block', 'hidden');
  avatarFileIcon.classList.replace('hidden', 'block');
  avatarFileName.innerText = 'Select image from your device';
  file = null;
}

function updateAvatar() {
  if (!file) return notifier.showError('Please select an image first.');
  avatarLoadingIcon.classList.replace('hidden', 'block');
  avatarFileIcon.classList.replace('block', 'hidden');
  var xhr = new XMLHttpRequest();
  var formData = new FormData();
  formData.append('avatar', file);
  xhr.open('POST', '/api/user-update-avatar', true);
  xhr.withCredentials = true;
  xhr.addEventListener( 'load', function() {
    const out = JSON.parse(xhr.responseText );
    const outJson = JSON.parse( out );
    if( xhr.status===200 ) {
      userAvatarImg.src = `/files${outJson.avatarUrl}`;
      notifier.showSuccess('Image uploaded successfully.');
      updateAvatarResp();
      setTimeout(() => {
        location.reload();
      }, 3500);
    } else {
      updateAvatarResp();
      notifier.showError(outJson.error);
    }
  } );
  xhr.addEventListener( 'error', function() {
    updateAvatarResp();
    notifier.showError( 'Network error' );
  } );
  xhr.addEventListener( 'abort', function() {
    updateAvatarResp();
    notifier.showWarning( 'Upload aborted' );
  }, false );
  xhr.send(formData);
}

function cancelUpdateAvatar() {
  popupUpdateAvatar.classList.replace('flex', 'hidden');
}

let popupUpdateProfile = document.getElementById('popupUpdateProfile');
let fullNameInput = document.getElementById('fullNameInput');
let locationInput = document.getElementById('locationInput');
let websiteInput = document.getElementById('websiteInput');
let updateProfileLoadingIcon = document.getElementById('updateProfileLoadingIcon');
let updateProfileButtonText = document.getElementById('updateProfileButtonText');

function updateProfileResp() {
  updateProfileLoadingIcon.classList.replace('block', 'hidden');
  updateProfileButtonText.classList.replace('hidden', 'block');
}

function openPopupUpdateProfile() {
  popupUpdateProfile.classList.replace('hidden', 'flex');
}

function cancelUpdateProfile() {
  updateProfileResp()
  popupUpdateProfile.classList.replace('flex', 'hidden');
}

async function updateProfile() {
  if (!fullNameInput.value) return notifier.showError('Please enter your full name.');
  if (!locationInput.value) return notifier.showError('Please enter your location.');
  if (!websiteInput.value) return notifier.showError('Please enter your website.');

  updateProfileLoadingIcon.classList.replace('hidden', 'block');
  updateProfileButtonText.classList.replace('block', 'hidden');
  try {
    const resp = await fetch('/api/user-update-profile', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/jsoa',
      },
      body: JSON.stringify({
        full_name: fullNameInput.value,
        location: locationInput.value,
        website: websiteInput.value
      }),
    });

    if (resp.ok) {
      const creds = await resp.json();
      console.log(creds);
      updateProfileResp()
      notifier.showSuccess(creds.data);
      setTimeout(()=>{
        window.location.reload();
      }, 1300)
    } else {
      const creds = await resp.json();
      notifier.showError(creds.errors+': '+creds.data);
      console.log(creds);
      updateProfileResp()
      return;
    }
  } catch (e) {
    console.log(e);
    notifier.showError('Error: ', e);
    updateProfileResp()
    return
  }
}