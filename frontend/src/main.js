import './style.css';
import './app.css';

import "@shoelace-style/shoelace/dist/themes/dark.css";
import "@shoelace-style/shoelace/dist/components/select/select.js";
import "@shoelace-style/shoelace/dist/components/option/option.js";
import "@shoelace-style/shoelace/dist/components/input/input.js";
import '@shoelace-style/shoelace/dist/components/button/button.js';
import '@shoelace-style/shoelace/dist/components/divider/divider.js';
import '@shoelace-style/shoelace/dist/components/split-panel/split-panel.js';
import '@shoelace-style/shoelace/dist/components/dropdown/dropdown.js';
import '@shoelace-style/shoelace/dist/components/menu/menu.js';
import '@shoelace-style/shoelace/dist/components/menu-item/menu-item.js';
import '@shoelace-style/shoelace/dist/components/dialog/dialog.js';
import * as App from '../wailsjs/go/main/App'
import { initCategories } from './category.js';
import { initTransactions } from './transaction.js';

async function init() {
    try {
        initCategories();
        initTransactions();
        renderSelectBudget(); // TODO Переделать это полная хуйня
    } catch (err) {
        console.error(err);
    }
}

async function renderSelectBudget() {
    const budgets = await App.BudgetList();
    if (budgets.length === 0) {
        renderCreateBudget();
        return;
    }

    document.querySelector('#select-row').innerHTML = `
        <sl-select id="budget-select"></sl-select>
        <sl-button id="create-budget" variant="primary"">
            Создать
        </sl-button>
    `;

    const select = document.getElementById("budget-select");
    select.value = budgets[0].id.toString();

    budgets.forEach(budget => {
        const option = document.createElement("sl-option");

        option.value = budget.id;
        option.textContent = `${budget.name} (${budget.start} - ${budget.end})`;

        select.appendChild(option);
    });

    const button = document.getElementById("create-budget");

    button?.addEventListener("click", renderCreateBudget);

    select.addEventListener("sl-change", (event) => {
        console.log(event.target.value);
    });
}

async function renderCreateBudget() {
    document.querySelector('#select-row').innerHTML = `
        <sl-input
            id="budget-name"
            placeholder="Название"
            required>
        </sl-input>
        <sl-input 
            id="start-date"
            type="date"
            required>
        </sl-input>
        <sl-input
            id="end-date"
            type="date"
            required>
        </sl-input>
        <sl-button id="save-budget" variant="primary"">
            Сохранить
        </sl-button>
        <sl-button id="close-budget-create" variant="default"">
            Отмена
        </sl-button>
    `;
    
    const nameInput = document.getElementById("budget-name");
    const startDate = document.getElementById("start-date");
    const endDate = document.getElementById("end-date");
    const saveButton = document.getElementById("save-budget");
    const closeButton = document.getElementById("close-budget-create");

    saveButton?.addEventListener("click", async () => {
        if (!nameInput.value.trim()) {
            nameInput.setCustomValidity("Введите название");
            nameInput.reportValidity();
            return;
        }
        nameInput.setCustomValidity("");

        if (!startDate.value) {
            startDate.setCustomValidity("Выберите начало диапазона");
            startDate.reportValidity();
            return;
        }
        startDate.setCustomValidity("");

        if (!endDate.value) {
            endDate.setCustomValidity("Выберите конец диапазона");
            endDate.reportValidity();
            return;
        }
        endDate.setCustomValidity("");

        if (startDate.value > endDate.value) {
            endDate.setCustomValidity("Дата начала не может быть позже даты окончания");
            endDate.reportValidity();
            return;
        }
        endDate.setCustomValidity("");

        try {
            await App.CreateBudget({
                name: nameInput.value.trim(),
                start: startDate.value,
                end: endDate.value,
            });
            renderSelectBudget();
        } catch (err) {
            console.error(err);
        }
    });

    closeButton?.addEventListener("click", renderSelectBudget);
}

init();