let eye = document.getElementById('eye');
let eyeSlash = document.getElementById('eye-slash');

function togglePassword() {
  if (password.type === 'password') {
    password.type = 'text';
    eyeSlash.classList.replace('hidden', 'block');
    eye.classList.replace('block', 'hidden');
  } else {
    password.type = 'password';
    eyeSlash.classList.replace('block', 'hidden');
    eye.classList.replace('hidden', 'block');
  }
}