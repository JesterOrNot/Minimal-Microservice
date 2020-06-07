function getTextInput() {
    return document.getElementById("greetInput").value;
}

function greet() {
    fetch(`/api/hello?name=${getTextInput()}`).then(response => {
        return response.json();
    }).then(data => {
        document.getElementById("text").innerText = data.message;
    });
    return true;
}
