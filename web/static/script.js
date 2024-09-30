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

// Conversion Functions
function convertLength() {
    const value = parseFloat(document.getElementById('lengthInput').value);
    const from = document.getElementById('lengthFrom').value.trim().toLowerCase();
    const to = document.getElementById('lengthTo').value.time().toLowerCase();

    const conversions = {
        ft: {
            cm: value * 30.48,
            m: value * 0.3048
        },
        m: {
            ft: value / 0.3048,
            cm: value * 100
        },
        cm: {
            ft: value / 30.48,
            m: value / 100
        }
    };

    const result = conversions[from] && conversions[from][to];
    if (result !== undefined) {
        document.getElementById('resultText').innerHTML = `Length Conversion: ${value} ${from} = ${result.toFixed(2)} ${to}`;
    } else {
        document.getElementById('resultText').innerHTML = 'Invalid conversion units.';
    }
}

