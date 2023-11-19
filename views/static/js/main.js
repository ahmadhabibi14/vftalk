let menuLists = document.querySelectorAll("#menu_list > a");
let currentPathname = window.location.pathname;

console.log("Current path name is " + currentPathname);

menuLists.forEach(function(link) {
  if (link.getAttribute("href") === currentPathname) {
    link.classList.add("font-bold");
    link.classList.add("bg-[#0a0a0a]");
    link.classList.remove("font-semibold");
    const svgs = link.querySelectorAll('svg');
    svgs.forEach(function(svg, index) {
      if (index === 0) {
        svg.classList.add("hidden");
      } else if (index === 1) {
        svg.classList.remove("hidden");
      }
    });
  }
});