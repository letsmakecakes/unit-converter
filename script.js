// Function to show the correct tab based on user selection
function showTab(tab) {
    // Hide all sections first
    document.getElementById('length-section').style.display = 'none';
    document.getElementById('weight-section').style.display = 'none';
    document.getElementById('temperature-section').style.display = 'none';

    // Remove the active class from all tabs
    document.getElementById('length-tab').classList.remove('active');
    document.getElementById('weight-tab').classList.remove('active');
    document.getElementById('temperature-tab').classList.remove('active');

    // Show the selected tab's section and add the active class
    if (tab === 'length') {
        document.getElementById('length-section').style.display = 'block';
        document.getElementById('length-tab').classList.add('active');
    } else if (tab === 'weight') {
        document.getElementById('weight-section').style.display = 'block';
        document.getElementById('weight-tab').classList.add('active');
    } else if (tab === 'temperature') {
        document.getElementById('temperature-section').style.display = 'block';
        document.getElementById('temperature-tab').classList.add('active');
    }
}

// Function to convert length
function convertLength() {
    const length = document.getElementById('lengthInput').value;
    const fromUnit = document.getElementById('lengthFrom').value;
    const toUnit = document.getElementById('lengthTo').value;

    // Call backend for length conversion (use the correct backend API)
    // Example result:
    const result = `${length} ${fromUnit} = ${length * 30.48} ${toUnit}`; // For demonstration purposes
    document.getElementById('resultValue').innerText = result;
    document.querySelector('.result').style.display = 'block';
}

// Function to convert weight
function convertWeight() {
    const weight = document.getElementById('weightInput').value;
    const fromUnit = document.getElementById('weightFrom').value;
    const toUnit = document.getElementById('weightTo').value;

    // Call backend for weight conversion (use the correct backend API)
    // Example result:
    const result = `${weight} ${fromUnit} = ${weight * 2.20462} ${toUnit}`; // For demonstration purposes
    document.getElementById('resultValue').innerText = result;
    document.querySelector('.result').style.display = 'block';
}

// Function to convert temperature
function convertTemperature() {
    const temp = document.getElementById('temperatureInput').value;
    const fromUnit = document.getElementById('temperatureFrom').value;
    const toUnit = document.getElementById('temperatureTo').value;

    // Call backend for temperature conversion (use the correct backend API)
    // Example result:
    let result = `${temp} ${fromUnit}`;
    if (fromUnit === 'C' && toUnit === 'F') {
        result += ` = ${(temp * 9/5) + 32} ${toUnit}`;
    } else if (fromUnit === 'F' && toUnit === 'C') {
        result += ` = ${(temp - 32) * 5/9} ${toUnit}`;
    }
    document.getElementById('resultValue').innerText = result;
    document.querySelector('.result').style.display = 'block';
}

// Function to reset the form
function resetForm() {
    document.getElementById('lengthInput').value = '';
    document.getElementById('weightInput').value = '';
    document.getElementById('temperatureInput').value = '';
    document.getElementById('resultValue').innerText = '';
    document.querySelector('.result').style.display = 'none';
}
