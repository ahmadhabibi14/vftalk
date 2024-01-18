let menuLists = document.querySelectorAll('#menu_list > a');
let logoutBtn = document.getElementById('logoutBtn');
let logoutLoadingPopup = document.getElementById('logoutLoadingPopup');
let currentPathname = window.location.pathname;

menuLists.forEach(function(link) {
  if (link.getAttribute('href') === currentPathname || (link.getAttribute('href') === '/' && currentPathname === '/direct')) {
    link.classList.add('font-bold');
    link.classList.add('bg-[#131313]');
    link.classList.remove('font-semibold');
    const svgs = link.querySelectorAll('svg');
    svgs.forEach(function(svg, index) {
      if (index === 0) {
        svg.classList.add('hidden');
      } else if (index === 1) {
        svg.classList.remove('hidden');
      }
    });
  }
});

logoutBtn.addEventListener('click', () => {
  logoutLoadingPopup.classList.remove('hidden');
  logoutLoadingPopup.classList.add('flex');
  document.cookie = `auth=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;`;
  localStorage.removeItem('username');
  setTimeout(() => {
    window.location.href = '/';
  }, 1500)
})
