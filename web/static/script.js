// Function to load HTML content dynamically
function showTab(tab) {
    const converterContent = document.getElementById('converter-content');
    fetch(`${tab}.html`)
        .then(response => {
            if (!response.ok) {
                throw new Error(`Could not fetch ${tab}.html`);
            }
            return response.text();
        })
        .then(data => {
            converterContent.innerHTML = data;
            // Update active tab
            document.querySelectorAll('.tabs a').forEach(a => a.classList.remove('active'));
            document.getElementById(`${tab}-tab`).classList.add('active');
        })
        .catch(error => {
            console.error('Error loading tab:', error);
            converterContent.innerHTML = `<p>Error loading content. Please try again later.</p>`
        });
}