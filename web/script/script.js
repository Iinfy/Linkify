const linkInput = document.getElementById("linkInput");
const shortBtn = document.getElementById("shortBtn");

shortBtn.addEventListener("click", async (e) => {
  e.preventDefault();
  const url = linkInput.value.trim();
  const base_url = window.location.origin;

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


    if (!linkInput.value.startsWith("https") && !linkInput.value.startsWith("http")) {
      alert("Link URL is invalid.");
      return
    }
    shortBtn.textContent = "shorted";
    shortBtn.style.backgroundColor = "green";
    addToRecentLink(base_url, data.url);
  } catch (error) {
    console.error(error);
  }
});

function addToRecentLink(base_url, shortUrl) {
  const fullUrl = `${base_url}/s/${shortUrl}`;
  const recentLinksList = document.getElementById("recentLinksList");

  if (!recentLinksList) {
    console.error("recentLinksList not found");
    return;
  }

  const item = document.createElement("div");
  item.className = "recent-item";
  item.style.opacity = "0";
  item.style.transform = "translateY(-10px)";
  item.style.transition = "all 0.3s ease";

  item.innerHTML = `
    <div class="text-wrapper">
      <a href="${fullUrl}" class="url" target="_blank" rel="noopener">${fullUrl}</a>
      <p class="Original-url">${linkInput.value}</p>
    </div>
    <button class="copy-btn">Copy</button>
  `;

  const copyBtn = item.querySelector('.copy-btn');
  copyBtn.addEventListener("click", (e) => {
    e.preventDefault();
    navigator.clipboard.writeText(fullUrl).then(() => {
      copyBtn.textContent = "Copied!";
      copyBtn.style.backgroundColor = "#7B9AB5";

      setTimeout(() => {
        copyBtn.textContent = "Copy";
        copyBtn.style.backgroundColor = "#383837";
      }, 2000);
    });
  });

  recentLinksList.insertBefore(item, recentLinksList.firstChild);

  while (recentLinksList.children.length > 3) {
    recentLinksList.removeChild(recentLinksList.lastChild);
  }

  requestAnimationFrame(() => {
    item.style.opacity = "1";
    item.style.transform = "translateY(0)";
  });

  setTimeout(() => {
    shortBtn.textContent = "Create";
    shortBtn.style.backgroundColor = "#383837";
  }, 3000);
}