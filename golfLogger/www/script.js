document.getElementById("callapi").addEventListener("click", callAPI);
function callAPI() {
    fetch("http://localhost:9090/api/", { mode: 'cors' })
        .then(function (response) {
        console.log(response.text());
    });
    //.then(response =)
}
