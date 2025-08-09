document.getElementById("qrForm").addEventListener("submit", function(event) {
	event.preventDefault();  // previne o reload da página
	gerarQR();
});

function gerarQR() {
	const input = document.getElementById("urlInput").value.trim();
	if (!input) {
		alert("Digite um endereço!");
		return;
	}
	const url = "https://" + input;
	const img = document.getElementById("qrCodeImg");
	img.src = "/qrcode?url=" + encodeURIComponent(url);
	img.style.display = "block";
}
