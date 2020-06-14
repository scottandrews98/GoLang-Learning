
document.getElementById("callapi").addEventListener("click", postData)

function postData(){
    var response: string = callAPI()

    console.log(response);
}

async function callAPI(){
    fetch("http://localhost:9090/api/", {mode: 'cors'})
    .then(response =>{
        return response.text();
    })
}