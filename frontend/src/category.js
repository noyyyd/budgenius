import {Categories, CreateCategory, DeleteCategory} from '../wailsjs/go/main/App';

export async function initCategories() {
    await render();
    initHandlers();
}

const categoryType = new Map([
    ['income', 'Доходы'],
    ['expenses', 'Расходы'],
]);

async function render() {
    const categories = await Categories();

    const tbody = document.querySelector('#categories-table');
    tbody.innerHTML = '';
    
    for (const category of categories) {
        const row = document.createElement('tr');
        row.dataset.id = category.ID;
        row.dataset.name = category.Name;

        const type = categoryType.get(category.Type)

        row.innerHTML = `
            <td>${category.Name}</td>
            <td>${type}</td>
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
    const createDialog = document.getElementById('create-category-dialog');

    document.getElementById("create-category").addEventListener('click', openCreateDialog);
    createDialog.querySelector('[name="create"]').addEventListener('click', create);
    createDialog.querySelector('[name="cancel"]').addEventListener('click', closeCreateDialog);

    const removeDialog = document.getElementById('remove-category-dialog');

    // Обработчик на таблице отлавливает события выбора элемента меню и по value 
    // в sl-menu-item выбирает каким методом будет обработано событие.
    //
    // sl-select - специальное событие которое генерит shoelace при клике на sl-menu-item
    document.getElementById("categories-table").addEventListener('sl-select', async event => {
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
    document.getElementById('create-category-dialog').show();
}

async function create() {
    const dialog = document.getElementById('create-category-dialog');
    const nameInput = dialog.querySelector('[name="name"]');
    const typeInput = dialog.querySelector('[name="type"]');
    const error = dialog.querySelector('[name="error"]');

    if (!nameInput.value.trim()) {
        nameInput.setCustomValidity("Введите название");
        nameInput.reportValidity();
        return;
    }
    nameInput.setCustomValidity("");

    if (!typeInput.value.trim()) {
        typeInput.setCustomValidity("Выберите тип");
        typeInput.reportValidity();
        return;
    }
    typeInput.setCustomValidity("");

    try {
        await CreateCategory({
            name: nameInput.value.trim(),
            type: typeInput.value,
        });
        error.textContent = '';
        nameInput.value = '';
        typeInput.value = '';
        dialog.hide()
        render();
    } catch (err) {
        error.textContent = err;
        error.style.color = 'red';
    }
}

async function closeCreateDialog() {
    const dialog = document.getElementById('create-category-dialog');
    const nameInput = dialog.querySelector('[name="name"]');
    const typeInput = dialog.querySelector('[name="type"]');
    const error = dialog.querySelector('[name="error"]');
    
    nameInput.value = '';
    typeInput.value = '';
    error.textContent = '';
    
    dialog.hide();
}

async function openRemoveDialog(event) {
    const dialog = document.getElementById('remove-category-dialog');
    const text = dialog.querySelector('[name="text"]');

    const row = event.detail.item.closest('tr');
    text.textContent = `Вы уверены что хотите удалить категорию "${row.dataset.name}"?`;

    dialog.dataset.id = Number(row.dataset.id);
    dialog.show();
}

async function remove() {
    const dialog = document.getElementById('delete-category-dialog');
    const error = dialog.querySelector('[name="error"]');

    const id = Number(dialog.dataset.id);

    try {
        await DeleteCategory(id);
        error.textContent = '';
        dialog.hide()
        render();
    } catch (err) {
        error.style.color = 'red';
        error.textContent = err;
    }
}

async function closeRemoveDialog() {
    const dialog = document.getElementById('remove-category-dialog');
    const text = dialog.querySelector('[name="text"]');
    const error = dialog.querySelector('[name="error"]');
    
    text.value = '';
    error.textContent = '';
    
    dialog.hide();
}