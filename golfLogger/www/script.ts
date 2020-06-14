
document.getElementById("callapi").addEventListener("click", postData)

function postData(){
    var response: Promise<String> = callAPI()

    console.log(Promise.resolve(response));
}

async function callAPI(){
    
    var test = await fetch("http://localhost:9090/api/", {mode: 'cors'})
    let text = await test.text();

    return text
}