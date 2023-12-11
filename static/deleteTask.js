var xhr = new XMLHttpRequest();

function DeleteTask(Title) {
    console.log("Title is ", Title)
    var task = JSON.stringify({
        title: Title
    });

    xhr.open("POST", "/api/user/task/delete");
    xhr.setRequestHeader("Content-type", "application/json; charset=utf-8");

    xhr.onload = function () {
        if (xhr.status >= 200 && xhr.status < 300) {

            location.reload(true);

        }
    };
    xhr.send(task);
}