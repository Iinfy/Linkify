const linkInput = document.getElementById("linkInput")
const shortBtn = document.getElementById("shortBtn")

shortBtn.addEventListener("click", (e) => {
    e.preventDefault()
    const url = linkInput.value.trim()

    try {
        const response = await fetch("/slink", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({url: url})
        })
        
        const data = await response.json()
    } catch(error) {
        console.error(error)
    }

})
