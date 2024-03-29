// Toggle active class on menu items
function toggleActive(elem) {
  const menu = document.querySelector(elem);

  menu.addEventListener("click", (e) => {
    if (e.target.tagName === "A") {
      menu.querySelector(".active").classList.remove("active");
      e.target.classList.add("active");
    }
  });
}
