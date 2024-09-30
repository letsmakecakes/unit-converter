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

function convertWeight() {
    const value = parseFloat(document.getElementById('weightInput').value);
    const from = document.getElementById('weightFrom').value.trim().toLowerCase();
    const to = document.getElementById('weightTo').value.trim().toLowerCase();

    const conversions = {
        kg: {
            lbs: value * 2.20462,
            g: value * 1000
        },
        lbs: {
            kg: value / 2.20462,
            g: value * 453.592
        },
        g: {
            kg: value / 1000,
            lbs: value / 453.592
        }
    };

    const result = conversions[from] && conversions[from][to];
    if(result !== undefined) {
        document.getElementById('resultText').innerText = `Weight Conversion: ${value} ${from} = ${result.toFixed(2)} ${to}`;
    } else {
        document.getElementById('resultText').innerText = 'Invalud conversion  units.';
    }
}

function convertTemperature() {
    const value = parseFloat(document.getElementById('temperatureInput').value);
    const from = document.getElementById('temperatureFrom').value.trim().toUpperCase();
    const to = document.getElementById('temperatureTo').value.trim().toUpperCase();

    let result;

    if (from === 'C' && to === 'F') {
        result = (value * 9/5) + 32;
    } else if (from === 'F' && to === 'C') {
        result = (value - 32) * 5/9; 
    } else if (from === to) {
        result = value;
    } else {
        document.getElementById('resultText').innerText = 'Invalid temperature units.';
        return;
    }

    document.getElementById('resultText').innerText = `Temperature Conversion: ${value} ${from} = ${result.toFixed(2)} ${to}`;
}

