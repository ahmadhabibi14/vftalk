let popupUpdateAvatar = document.getElementById("popupUpdateAvatar");

function openPopupUpdateAvatar() {
  popupUpdateAvatar.classList.remove("hidden");
  popupUpdateAvatar.classList.add("flex");
}

function updateAvatar() {
  popupUpdateAvatar.classList.add("hidden");
  popupUpdateAvatar.classList.remove("flex");
  var avatar = document.getElementById("avatar");
  var file = avatar.files[0];
  if (!file) {
    alert("Please select an image first.");
    return;
  }
  var xhr = new XMLHttpRequest();
  var formData = new FormData();
  formData.append("avatar", file);
  xhr.open("POST", "/api/user-update-avatar", true);
  xhr.withCredentials = true;
  xhr.addEventListener( 'load', function( event ) {
    if( xhr.status===200 ) {
      const out = JSON.parse( event.target.responseText );
      notifier.showSuccess("Image uploaded successfully.");
      console.log(out);
    } else if( xhr.status===413 ) {
      notifier.showError( 'Image too large' );
    } else {
      notifier.showError( `Error: ${xhr.status}  ${xhr.statusText}` );
      console.log( event.target.responseText );
    }
  } );
  xhr.addEventListener( 'error', function( event ) {
    alert( 'Network error' );
  } );
  xhr.addEventListener( 'abort', function( event ) {
    alert( 'Upload aborted' );
  }, false );
  xhr.send(formData);
}

function cancelUpdateAvatar() {
  popupUpdateAvatar.classList.add("hidden");
  popupUpdateAvatar.classList.remove("flex");
}