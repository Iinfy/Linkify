const linkInput = document.getElementById("linkInput");
const shortBtn = document.getElementById("shortBtn");

shortBtn.addEventListener("click", async (e) => {
  e.preventDefault();
  const url = linkInput.value.trim();
  shortBtn.textContent = "shorted";
  shortBtn.style.backgroundColor = "green";

  try {
    const response = await fetch(`http://localhost:9090/slink`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ url: url }),
    });

    const data = await response.json();
  } catch (error) {
    console.error(error);
  }
});
