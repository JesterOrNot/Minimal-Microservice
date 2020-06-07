function getTextInput() {
    element = document.getElementById("greetInput");
    value = element.value
    element.value = "";
    return value;
}

function greet() {
    fetch(`/api/hello?name=${getTextInput()}`).then(response => {
        return response.json();
    }).then(data => {
        console.log(data);
        document.getElementById("text").innerText = data.message;
    });
    return true;
}
