
document.getElementById("callapi").addEventListener("click", postData)
document.addEventListener("DOMContentLoaded", loadEvents)

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
    var url = 'http://localhost:9090/api/addsession';
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

function loadEvents(){
    getSessions();
    getClubs();
}   

function getSessions(){
    // Gets all the golf sessions
    var http = new XMLHttpRequest();
    var url = 'http://localhost:9090/api/getsessions';
    http.open('GET', url, true);

    http.onreadystatechange = function() {
        if(http.readyState == 4 && http.status == 200) {
            let sessions = JSON.parse(http.responseText);

            for(var i = 0; i < sessions.length; i++) {
                var obj = sessions[i];
            
                var sessionsDiv = document.getElementById('golfEvents');

                sessionsDiv.innerHTML += '<p>Golf Type: '+ obj.GolfType+' Total Shots: '+obj.Value+'';
            }
        }
    }
    http.send();
}

var clubs;
function getClubs(){
    var http = new XMLHttpRequest();

    // Gets all the clubs and ranges that are used
    var url = 'http://localhost:9090/api/getclubs';
    http.open('GET', url, true);

    http.onreadystatechange = function() {
        if(http.readyState == 4 && http.status == 200) {
            clubs = JSON.parse(http.responseText);

            var clubsDiv = document.getElementById('clubDistances');

            for(var i = 0; i < clubs.length; i++) {
                var obj = clubs[i];

                var niceClubName = obj.clubName.replace(" ", "");

                clubsDiv.innerHTML += `
                    <div class="row" style="margin-top:8px;"><div class="col-sm-4"><p>`+obj.clubName+`</p></div><div class="col-sm-4"><p>`+obj.clubDistance+` Yards</p></div><div class="col-sm-4"><button data-toggle="collapse" href="#`+niceClubName+`" class="btn btn-primary">Edit</button></div>
                        <div class="collapse" id="`+niceClubName+`">
                            <div class="row">
                                <div class="col-sm-6">
                                    <input type="number" class="form-control" value="`+obj.clubDistance+`">
                                </div>
                                <div class="col-sm-6">
                                    <button data-toggle="collapse" href="#collapseExample" class="btn btn-primary">Save</button>
                                </div>
                            </div>
                        </div>
                    </div>`;
            }
        }
    }
    http.send();
}