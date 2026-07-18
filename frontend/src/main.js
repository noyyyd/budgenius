import './style.css';
import './app.css';

import "@shoelace-style/shoelace/dist/themes/dark.css";
import "@shoelace-style/shoelace/dist/components/select/select.js";
import "@shoelace-style/shoelace/dist/components/option/option.js";
import "@shoelace-style/shoelace/dist/components/icon/icon.js";
import '@shoelace-style/shoelace/dist/components/button/button.js';
import logo from './assets/images/logo-universal.png';
import {Greet,Budgets} from '../wailsjs/go/main/App';

async function init() {
    try {
        renderBudgetsSelect();
    } catch (err) {
        console.error(err);
    }
}

async function renderBudgetsSelect() {
    const select = document.getElementById("budget-select");

    select.addEventListener("sl-change", (event) => {
        console.log(event.target.value);
    });

    const budgets = await Budgets();

    budgets.forEach(budget => {
        const option = document.createElement("sl-option");

        option.value = budget.id;
        option.textContent = `${budget.name} (${budget.start} - ${budget.end})`;

        select.appendChild(option);
    });
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