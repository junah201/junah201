const h1 = document.getElementById("title");

function handleTitleClick() {
	h1.classList.toggle("clicked");
}

h1.onclick = handleTitleClick;
