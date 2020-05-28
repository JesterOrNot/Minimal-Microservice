function getTextInput() {
    return document.getElementById("greetInput").nodeValue;
}

function greet() {
    fetch(`127.0.0.1:8081/api/hello?name=${getTextInput()}`).then(response => {
        return response.json();
    }).then(data => {
        document.getElementById("text").innerText = data.message;
    });
}
