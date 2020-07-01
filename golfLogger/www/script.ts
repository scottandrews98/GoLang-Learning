
document.getElementById("callapi").addEventListener("click", postData)
document.addEventListener("DOMContentLoaded", loadEvents)

function postData(){    
    var golfTypeSelect = document.getElementById("golfType") as HTMLSelectElement;
    let golfType = golfTypeSelect.options[golfTypeSelect.selectedIndex].value;

    let totalShots = (document.getElementById("totalShots") as HTMLSelectElement).value;
    let wellHitShots = (document.getElementById("totalWellHit") as HTMLSelectElement).value;

    // Checks that total well hit shots is not larger than total hit
    if(wellHitShots > totalShots){
        alert("Your Total Well Hit Shots Cannot Be Larger Than Your Total Shots");
    }else{
        submitForm(golfType, totalShots, wellHitShots)
    }
}

function submitForm(golfType, totalShots, totalWellHit){
    var http = new XMLHttpRequest();
    var url = 'http://localhost:9090/api/addsession';
    var params = 'golftype='+golfType+'&shots='+totalShots+'&wellHit='+totalWellHit+'';
    http.open('POST', url, true);

    //Send the proper header information along with the request
    http.setRequestHeader('Content-type', 'application/x-www-form-urlencoded');

    http.onreadystatechange = function() {//Call a function when the state changes.
        if(http.readyState == 4 && http.status == 200) {
            alert(http.responseText);
            getTotalShots("getshots");
            getTotalShots("getgoodshots");  
            var sessionsDiv = document.getElementById('golfEvents');
            sessionsDiv.innerHTML = '';
            getSessions()
        }
    }
    http.send(params);
}

function loadEvents(){
    getSessions();
    getClubs();
    getTotalShots("getshots");
    getTotalShots("getgoodshots");
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
                    <div class="row" style="margin-top:8px;"><div class="col-sm-4"><p>`+obj.clubName+`</p></div><div class="col-sm-4"><p id="output`+niceClubName+`">`+obj.clubDistance+` Yards</p></div><div class="col-sm-4"><button data-toggle="collapse" href="#`+niceClubName+`" class="btn btn-primary">Edit</button></div>
                        <div class="collapse" id="`+niceClubName+`">
                            <div class="row">
                                <div class="col-sm-6">
                                    <input type="number" id="input`+niceClubName+`" class="form-control" data-id="`+obj.Id+`" value="`+obj.clubDistance+`">
                                </div>
                                <div class="col-sm-6">
                                    <button data-toggle="collapse" href="#collapseExample" data-inputid="`+niceClubName+`" class="btn btn-primary saveDistance">Save</button>
                                </div>
                            </div>
                        </div>
                    </div>`;
            }

            // Add in new event listener
            const breakdownButton = document.querySelectorAll('.saveDistance');
            breakdownButton.forEach(function(btn) {
                btn.addEventListener('click', function() {
                    let input = document.getElementById("input"+this.getAttribute('data-inputid')) as HTMLSelectElement;
                    // get id
                    let id = input.getAttribute("data-id");
                    // get distance 
                    let distance = Number(input.value);
                    // run save function
                    let response = saveDistance(distance, id);
                    // if comes back good then update front end value
                    document.getElementById("output"+this.getAttribute('data-inputid')).textContent = input.value + " Yards"        
                });
            });
        }
    }
    http.send();
}


// Sends a post request to the backend to save a new distance
function saveDistance(distance: number, id: string){
    var http = new XMLHttpRequest();
    var url = 'http://localhost:9090/api/savedistance';
    var params = 'distance='+distance+'&id='+id+'';
    http.open('POST', url, true);

    //Send the proper header information along with the request
    http.setRequestHeader('Content-type', 'application/x-www-form-urlencoded');

    http.onreadystatechange = function() {//Call a function when the state changes.
        if(http.readyState == 4 && http.status == 200) {
            if(this.responseText == "Data Updated"){
                return "updated";
            }else{
                return "error";
            }
        }
    }

    http.send(params);
}


// Add in request to get total shots carried out
function getTotalShots(shotType){
    var http = new XMLHttpRequest();
    var url = 'http://localhost:9090/api/'+shotType+'';
    http.open('GET', url, true);

    http.onreadystatechange = function() {
        if(http.readyState == 4 && http.status == 200) {
            let sessions = JSON.parse(http.responseText);

            for(var i = 0; i < sessions.length; i++) {
                var obj = sessions[i];
                
                if(shotType == "getshots"){
                    document.getElementById("golfShots").textContent = obj.TotalShots
                }else{
                    document.getElementById("wellHit").textContent = obj.TotalShots + "%"
                }
            }
        }
    }
    http.send();
}