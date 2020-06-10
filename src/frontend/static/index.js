function getTextValue(id) {
  return function () {
    element = document.getElementById(id);
    value = element.value;
    element.value = "";
    return value;
  };
}

function greet() {
  body = {
    "name": `${getTextValue("greetInput")()}`,
    "age": parseInt(`${getTextValue("ageInput")()}`),
  }
  fetch(`/api/vote`, {
    method: "POST",
    body: JSON.stringify(body),
  })
    .then(response => {
      return response.json();
    })
    .then(data => {
      console.log(data)
      if (data.message !== undefined) {
        document.getElementById("text").innerText = data.message;
      } else {
        alert(data.error);
      }
    });
}
