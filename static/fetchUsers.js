function fetchUsers() {
    var xhr = new XMLHttpRequest();
    xhr.open('GET', '/api/user/findAll');
    xhr.onload = function () {
        if (xhr.status >= 200 && xhr.status < 300) {
            var users = JSON.parse(xhr.responseText);
            var selectElement = document.getElementById('assignedUser');

            selectElement.innerHTML = '';

            users.forEach(function (user) {
                var option = document.createElement('option');
                option.textContent = user.Name;
                selectElement.appendChild(option);
            });
        } else {
            console.error('Error fetching users:', xhr.statusText);
        }
    };
    xhr.onerror = function () {
        console.error('Network error while fetching users');
    };
    xhr.send();
}
fetchUsers();