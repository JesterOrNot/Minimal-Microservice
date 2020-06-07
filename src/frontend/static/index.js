function getTextInput() {
    element = document.getElementById("greetInput");
    value = element.value
    element.value = "";
    return value;
}

function greet() {
    fetch(`/api/hello?name=${getTextInput()}`).then(response => {
        console.log(response.json());
        return response.json();
    }).then(data => {
        document.getElementById("text").innerText = data.message;
    });
    return true;
}
