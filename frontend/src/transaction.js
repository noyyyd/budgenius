import {Transactions, Categories, CreateTransaction, DeleteTransaction} from '../wailsjs/go/main/App';

export async function initTransactions() {
    await render();
    initHandlers();
}

async function render() {
    const transactions = await Transactions();

    const tbody = document.querySelector('#transactions-table');
    tbody.innerHTML = '';
    
    for (const transaction of transactions) {
        const row = document.createElement('tr');
        row.dataset.id = transaction.ID;
        row.dataset.date = transaction.Date;
        row.dataset.amount = transaction.Amount;
        row.dataset.category = transaction.Category;

        const amount = transaction.IsIncome
            ? transaction.Amount
            : "-" + transaction.Amount;

        row.innerHTML = `
            <td>${transaction.Date}</td>
            <td>${transaction.Category}</td>
            <td>${transaction.Comment}</td>
            <td>${amount}</td>
            <td>
                <sl-dropdown>
                    <sl-button variant="text" slot="trigger" size="small">&#x22EE;</sl-button>
                        <sl-menu>
                            <sl-menu-item value="edit" disabled>Изменить</sl-menu-item>
                            <sl-menu-item value="remove">Удалить</sl-menu-item>
                        </sl-menu>
                </sl-dropdown>
            </td>
        `;
        tbody.append(row);
    }
}

function initHandlers() {
    const createDialog = document.getElementById('create-transaction-dialog');

    document.getElementById("create-transaction").addEventListener('click', openCreateDialog);
    createDialog.querySelector('[name="create"]').addEventListener('click', create);
    createDialog.querySelector('[name="cancel"]').addEventListener('click', closeCreateDialog);

    const removeDialog = document.getElementById('remove-transaction-dialog');

    // Обработчик на таблице отлавливает события выбора элемента меню и по value 
    // в sl-menu-item выбирает каким методом будет обработано событие.
    //
    // sl-select - специальное событие которое генерит shoelace при клике на sl-menu-item
    document.getElementById("transactions-table").addEventListener('sl-select', async event => {
        console.log(event)
        console.log(event.detail.item.value)
        switch (event.detail.item.value) {
            // TODO
            // case 'edit':
            //     await edit(event);
            //     break;
            case 'remove':
                await openRemoveDialog(event);
                break;
            default:
                console.log("Unknown event:", event)
                break;
        }
    });
    removeDialog.querySelector('[name="remove"]').addEventListener('click', remove);
    removeDialog.querySelector('[name="cancel"]').addEventListener('click', closeRemoveDialog);
}

async function openCreateDialog() {
    const dialog = document.getElementById('create-transaction-dialog');
    const categoryInput = dialog.querySelector('[name="category"]');

    const categories = await Categories();

    for (const category of categories) {
        const option = document.createElement('sl-option');

        option.value = category.ID;
        option.textContent = category.Name;

        categoryInput.appendChild(option);
    }

    dialog.show();
}

async function create() {
    const dialog = document.getElementById('create-transaction-dialog')
    const dateInput = dialog.querySelector('[name="date"]');
    const amountInput = dialog.querySelector('[name="amount"]');
    const commentInput = dialog.querySelector('[name="comment"]');
    const categoryInput = dialog.querySelector('[name="category"]');
    const error = dialog.querySelector('[name="error"]');

    if (!dateInput.value.trim()) {
        dateInput.setCustomValidity("Выберите дату совершения транзакции");
        dateInput.reportValidity();
        return;
    }
    dateInput.setCustomValidity("");

    const amount = Number(amountInput.value.trim().replace(',', '.'));

    if (!amountInput.value.trim()) {
        amountInput.setCustomValidity("Введите сумму");
    } else if (amount <= 0) {
        amountInput.setCustomValidity("Сумма должна быть больше нуля");
    } else if (!/^\d+(\.\d{1,2})?$/.test(amount)) {
        amountInput.setCustomValidity("Некорректная сумма");
    } else  {
        amountInput.setCustomValidity("");
    }

    if (!amountInput.checkValidity()) {
        amountInput.reportValidity();
        return;
    }
    amountInput.setCustomValidity("");

    if (!categoryInput.value) {
        categoryInput.setCustomValidity("Выберите категорию");
        categoryInput.reportValidity();
        return;
    }
    categoryInput.setCustomValidity("");

    try {
        await CreateTransaction({
            date: dateInput.value,
            amount: amount * 100,
            comment: commentInput.value.trim(),
            categoryID: Number(categoryInput.value),
        });
        dateInput.value = '';
        amountInput.value = '';
        commentInput.value = '';
        categoryInput.replaceChildren();;
        error.textContent = '';
        dialog.hide();
        render();
    } catch (err) {
        error.textContent = err;
        error.style.color = 'red';
    }
}

async function closeCreateDialog() {
    const dialog = document.getElementById('create-transaction-dialog');
    const dateInput = dialog.querySelector('[name="date"]');
    const amountInput = dialog.querySelector('[name="amount"]');
    const commentInput = dialog.querySelector('[name="comment"]');
    const categoryInput = dialog.querySelector('[name="category"]');
    const error = dialog.querySelector('[name="error"]');
    
    dateInput.value = '';
    amountInput.value = '';
    commentInput.value = '';
    categoryInput.replaceChildren();;
    error.textContent = '';
    
    dialog.hide();
}

async function openRemoveDialog(event) {
    const dialog = document.getElementById('remove-transaction-dialog');
    const text = dialog.querySelector('[name="text"]');

    const row = event.detail.item.closest('tr');
    text.textContent = `Вы уверены, что хотите удалить транзакцию на сумму ${row.dataset.amount} из категории "${row.dataset.category}" от ${row.dataset.date}?`;

    dialog.dataset.id = Number(row.dataset.id);
    dialog.show();
}

async function remove() {
    const dialog = document.getElementById('remove-transaction-dialog');
    const error = dialog.querySelector('[name="error"]');

    const id = Number(dialog.dataset.id);

    try {
        await DeleteTransaction(id);
        error.textContent = '';
        dialog.hide()
        render();
    } catch (err) {
        error.style.color = 'red';
        error.textContent = err;
    }
}

async function closeRemoveDialog() {
    const dialog = document.getElementById('remove-transaction-dialog');
    const text = dialog.querySelector('[name="text"]');
    const error = dialog.querySelector('[name="error"]');
    
    text.value = '';
    error.textContent = '';
    
    dialog.hide();
}