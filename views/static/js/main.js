// Dropdown Menu
let dropdownMenu = document.getElementById("dropdownMenu");
let menuBtn = document.getElementById("menuBtn");
let logoutBtn = document.getElementById("logoutBtn");

dropdownMenu.style.display = "none";
menuBtn.addEventListener("click", () => {
  dropdownMenu.style.display = dropdownMenu.style.display === "none" ? "flex" : "none";
});

logoutBtn.addEventListener("click", () => {
  document.cookie = `auth=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;`;
  localStorage.removeItem('username');
  setTimeout(() => {
    window.location.href = "/";
  }, 1500)
})