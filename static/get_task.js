document.addEventListener('DOMContentLoaded', function () {
    var table = document.querySelector('table');

    // Объявим переменную originalTitle в области видимости скрипта
    var originalTitle;

    table.addEventListener('click', function (event) {
        var target = event.target;

        if (target.tagName === 'TD' && target.classList.contains('usernameCell')) {
            var username = target.textContent;

            var xhr = new XMLHttpRequest();
            xhr.open('POST', '/api/tasks/getTask');
            xhr.setRequestHeader('Content-Type', 'application/json;charset=UTF-8');

            xhr.onload = function () {
                if (xhr.status >= 200 && xhr.status < 300) {
                    var data = JSON.parse(xhr.responseText);
                    displayTasks(data);
                } else {
                    console.error('Error fetching tasks:', xhr.statusText);
                }
            };

            xhr.onerror = function () {
                console.error('Network error while fetching tasks');
            };

            var dataToSend = JSON.stringify({ username: username });
            xhr.send(dataToSend);
        }
    });

    function displayTasks(data) {
        var tasksTable = '<table>';
        tasksTable += '<tr><th>Задача</th><th>Описание</th><th>Статус</th><th>Опции</th></tr>';

        for (var i = 0; i < data.length; i++) {
            tasksTable += '<tr>';
            tasksTable += '<td class="editable-cell" contenteditable="false">' + data[i].Title + '</td>';
            tasksTable += '<td class="editable-cell" contenteditable="false">' + data[i].Description + '</td>';
            tasksTable += '<td class="editable-cell" contenteditable="false">' + data[i].Status + '</td>';
            tasksTable += '<td class="options-cell"><div class="options-icon">&#8942;</div>' +
                '<div class="options-popup">' +
                '<button class="edit-task-button"  onclick="enableEditing(this)">Редактировать</button>' +
                '<button class="delete-task-button" onclick="DeleteTask(\'' + data[i].Title + '\')">Удалить</button>' +
                '<button class="save-task-button" style="display: none;" onclick="SaveTask(\'' + data[i].Title + '\')">Сохранить</button>' +
                '</div></td>';
            tasksTable += '</tr>';
        }

        tasksTable += '</table>';

        document.getElementById('tasksContainer').innerHTML = tasksTable;
    }

    window.enableEditing = function (editButton) {
        var row = editButton.parentNode.parentNode.parentNode;

        var titleCell = row.querySelector('.editable-cell:nth-child(1)');
        var cells = row.querySelectorAll('.editable-cell');

        // Сохраняем исходное название задания
        originalTitle = titleCell.textContent;

        cells.forEach(function (cell) {
            cell.contentEditable = 'true';
        });

        // Показываем кнопку "Сохранить" и скрываем "Редактировать"
        var editButton = row.querySelector('.edit-task-button');
        var saveButton = row.querySelector('.save-task-button');

        editButton.style.display = 'none';
        saveButton.style.display = 'inline-block';
    };

    window.SaveTask = function (title) {
        var row = document.querySelector('.editable-cell[contenteditable="true"]').parentNode;
        var updatedTitle = row.querySelector('.editable-cell:nth-child(1)').textContent;
        var updatedDescription = row.querySelector('.editable-cell:nth-child(2)').textContent;
        var updatedStatus = row.querySelector('.editable-cell:nth-child(3)').textContent;

        var updatePayload = {
            oldTitle: originalTitle,
            title: updatedTitle,
            description: updatedDescription,
            status: updatedStatus
        };

        var xhr = new XMLHttpRequest();
        xhr.open('POST', '/api/user/task/update');
        xhr.setRequestHeader('Content-Type', 'application/json;charset=UTF-8');

        xhr.onload = function () {
            if (xhr.status >= 200 && xhr.status < 300) {
                // Обработка успешного обновления
                console.log('Task updated successfully');

                // Отключаем редактирование после сохранения
                var cells = row.querySelectorAll('.editable-cell');
                cells.forEach(function (cell) {
                    cell.contentEditable = 'false';
                });

                // Показываем кнопку "Редактировать" и скрываем "Сохранить"
                var editButton = row.querySelector('.edit-task-button');
                var saveButton = row.querySelector('.save-task-button');

                editButton.style.display = 'inline-block';
                saveButton.style.display = 'none';
            } else {
                console.error('Error updating task:', xhr.statusText);
            }
        };

        xhr.onerror = function () {
            console.error('Network error while updating task');
        };

        xhr.send(JSON.stringify(updatePayload));
    };
});


// document.addEventListener('DOMContentLoaded', function () {
//     var table = document.querySelector('table');
//
//     table.addEventListener('click', function (event) {
//         var target = event.target;
//         if (target.tagName === 'TD' && target.classList.contains('usernameCell')) {
//             var username = target.textContent;
//
//             var xhr = new XMLHttpRequest();
//             xhr.open('POST', '/api/tasks/getTask');
//             xhr.setRequestHeader('Content-Type', 'application/json;charset=UTF-8');
//
//             xhr.onload = function () {
//                 if (xhr.status >= 200 && xhr.status < 300) {
//                     var data = JSON.parse(xhr.responseText);
//                     displayTasks(data);
//                 } else {
//                     console.error('Error fetching tasks:', xhr.statusText);
//                 }
//             };
//
//             xhr.onerror = function () {
//                 console.error('Network error while fetching tasks');
//             };
//
//             var dataToSend = JSON.stringify({ username: username });
//             xhr.send(dataToSend);
//         }
//     });
//
//     function displayTasks(data) {
//         var tasksTable = '<table>';
//         tasksTable += '<tr><th>Задача</th><th>Описание</th><th>Статус</th><th>Опции</th></tr>';
//
//         for (var i = 0; i < data.length; i++) {
//             tasksTable += '<tr>';
//             tasksTable += '<td class="editable-cell" contenteditable="false">' + data[i].Title + '</td>';
//             tasksTable += '<td class="editable-cell" contenteditable="false">' + data[i].Description + '</td>';
//             tasksTable += '<td class="editable-cell" contenteditable="false">' + data[i].Status + '</td>';
//             tasksTable += '<td class="options-cell"><div class="options-icon">&#8942;</div>' +
//                 '<div class="options-popup"><button class="edit-task-button"  onclick="enableEditing(this)">Редактировать</button>' +
//                 '<button class="delete-task-button" onclick="DeleteTask(\'' + data[i].Title + '\')">Удалить</button></div></td>';
//             tasksTable += '</tr>';
//         }
//
//         tasksTable += '</table>';
//
//         document.getElementById('tasksContainer').innerHTML = tasksTable;
//     }
//
//     window.enableEditing = function (editButton) {
//         var row = editButton.parentNode.parentNode.parentNode;
//
//         var cells = row.querySelectorAll('.editable-cell');
//         cells.forEach(function (cell) {
//             cell.contentEditable = 'true';
//         });
//     };
// });