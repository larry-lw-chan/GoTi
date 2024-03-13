window.onload = function () {
  // Hamburger Menu Toggle
  document.getElementById("hamburger").addEventListener("click", function () {
    let header = document.querySelector("header");
    header.classList.toggle("change");
  });
};
