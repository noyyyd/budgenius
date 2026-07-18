import './style.css';
import './app.css';

import logo from './assets/images/logo-universal.png';
import {Greet,Budgets} from '../wailsjs/go/main/App';

async function init() {
    try {
        let budgets = await Budgets()
        renderBudgets(budgets);
    } catch (err) {
        console.error(err);
    }
}

function renderBudgets(budgets) {
    const container = document.getElementById("budgets");

    container.innerHTML = budgets
        .map(b => `<div>${b.id} — ${b.start} — ${b.end}</div>`)
        .join("");
}

document.querySelector('#app').innerHTML = `
    <img id="logo" class="logo">
      <div class="result" id="result">Please enter your name below 👇</div>
      <div class="input-box" id="input">
        <input class="input" id="name" type="text" autocomplete="off" />
        <button class="btn" onclick="greet()">Greet</button>
        <button class="btn" onclick="budgets()">Budgets</button>
      </div>
    </div>
`;
document.getElementById('logo').src = logo;

let nameElement = document.getElementById("name");
nameElement.focus();
let resultElement = document.getElementById("result");

// Setup the greet function
window.greet = async function () {
    // Get name
    let name = nameElement.value;

    // Check if the input is empty
    if (name === "") return;

    // Call App.Greet(name)
    try {
        // Call App.Greet(name)
        let result = await Greet(name);

        // Update result with data back from App.Greet()
        resultElement.innerText = result;

    } catch (err) {
        console.error(err);
    }
};

init();