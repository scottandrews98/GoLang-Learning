
document.getElementById("callapi").addEventListener("click", postData)

function postData(){
    //var response: Promise<String> = callAPI()
    //console.log(response);
    
    var golfTypeSelect = document.getElementById("golfType") as HTMLSelectElement;
    let golfType = golfTypeSelect.options[golfTypeSelect.selectedIndex].value;

    let totalShots = document.getElementById("totalShots") as HTMLSelectElement;
    let totalShotsValue = totalShots.value;

    submitForm(golfType, totalShotsValue)
}

function submitForm(golfType, totalShots){
    var http = new XMLHttpRequest();
    var url = 'http://localhost:9090/api/';
    var params = 'golftype='+golfType+'&shots='+totalShots+'';
    http.open('POST', url, true);

    //Send the proper header information along with the request
    http.setRequestHeader('Content-type', 'application/x-www-form-urlencoded');

    http.onreadystatechange = function() {//Call a function when the state changes.
        if(http.readyState == 4 && http.status == 200) {
            alert(http.responseText);
        }
    }
    http.send(params);
}