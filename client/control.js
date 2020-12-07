const url = "http://localhost:8000/api/messages";

function onget() {
    var headers = {}

    fetch(url, {
        method : "GET",
        mode: 'cors',
        headers: headers
    })
    .then((response) => {
        if (!response.ok) {
            throw new Error(response.error)
        }
        return response.json();
    })
    .then(data => {
        document.getElementById('messages').value = data.messages;
    })
    .catch(function(error) {
        document.getElementById('messages').value = error;
    });
}

function onpost() {
    fetch(url, {
        method : "POST",
        mode: 'cors',
        body: new URLSearchParams(new FormData(document.getElementById("form2"))),
    })
    .then((response) => {
        if (!response.ok) {
            throw new Error(response.error)
        }
        return response.json();
    })
    .then(data => {
        document.getElementById('messages').value = data.messages.join('\n');
    })
    .catch(function(error) {
        document.getElementById('messages').value = error;
    });
}