var xhr = new XMLHttpRequest();

function addTask() {
    var task = JSON.stringify({
        title: document.getElementById('title').value,
        description: document.getElementById('description').value,
        status: document.getElementById('status').value,
        assignedUser: document.getElementById('assignedUser').value,
    });

    xhr.open("POST", "/api/user/task/add");
    xhr.setRequestHeader("Content-type", "application/json; charset=utf-8");

    xhr.send(task);

    document.getElementById('addTaskForm').reset();
}
