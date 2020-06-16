document.getElementById("callapi").addEventListener("click", postData);
function postData() {
    //var response: Promise<String> = callAPI()
    //console.log(response);
    var golfTypeSelect = document.getElementById("golfType");
    var golfType = golfTypeSelect.options[golfTypeSelect.selectedIndex].value;
    var totalShots = document.getElementById("totalShots");
    var totalShotsValue = totalShots.value;
    submitForm(golfType, totalShotsValue);
}
function submitForm(golfType, totalShots) {
    var http = new XMLHttpRequest();
    var url = 'http://localhost:9090/api/';
    var params = 'golftype=' + golfType + '&shots=' + totalShots + '';
    http.open('POST', url, true);
    //Send the proper header information along with the request
    http.setRequestHeader('Content-type', 'application/x-www-form-urlencoded');
    http.onreadystatechange = function () {
        if (http.readyState == 4 && http.status == 200) {
            alert(http.responseText);
        }
    };
    http.send(params);
}
