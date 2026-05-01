const linkInput = document.getElementById("linkInput");
const shortBtn = document.getElementById("shortBtn");

shortBtn.addEventListener("click", async (e) => {
  e.preventDefault();
  const url = linkInput.value.trim();
  const base_url = window.location.origin
  try {
    const response = await fetch(`/slink`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ url: url }),
    });

    const data = await response.json();

    if (!response.ok) {
      return;
    }

    shortBtn.textContent = "shorted";
    shortBtn.style.backgroundColor = "green";

    linkInput.value = `${base_url}/s/${data.url}`
  } catch (error) {
    console.error(error);
  }
});
