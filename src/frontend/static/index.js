function getTextValue(id) {
    return function () {
        element = document.getElementById(id)
        value = element.value
        element.value = ""
        return value
    }
}

function greet() {
    fetch(`/api/hello?name=${getTextValue("greetInput")()}`).then(response => {
        return response.json();
    }).then(data => {
        document.getElementById("text").innerText = data.message;
    });
    fetch(`/api/vote?age=${getTextValue("ageInput")()}`).then(response => {
        return response.json();
    }).then(data => {
        oldText = document.getElementById("text").innerText
        message = ""
        if (typeof data.canVote !== "boolean") {
            alert(data.canVote)
        } else {
            message = data.canVote ? "you can vote!" : "you can't vote!"
        }
        document.getElementById("text").innerText = oldText + " " + message;
    });
    return true;
}
