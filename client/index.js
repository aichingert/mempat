let game = document.getElementById("game");

for (let i = 0; i < 4; i++) {
    let row = document.createElement("div");
    row.classList.add("square-row");

    for (let j = 0; j < 4; j++) {
        let square = document.createElement("div");
        square.id = `${i} ${j}`;
        square.classList.add("square")

        square.addEventListener("click", function() {
            if (this.style.backgroundColor != "rgb(195, 232, 141)") {
                socket.send(this.id);
            }
        });

        row.appendChild(square);
    }

    game.append(row);
}

let socket = new WebSocket("ws://" + document.location.host + "/ws");

console.log("Attempting Connection...");

socket.onopen = () => {
};

socket.onerror = error => { console.log("Socket Error: ", error); };

socket.onmessage = function(evt) {
    console.log(evt)

    let div = document.createElement("div");
    let sq = document.getElementById(evt.data);
    div.innerHTML = `${evt.data} ${sq.style.backgroundColor}`;

    sq.style.backgroundColor = "rgb(195, 232, 141)";

    game.appendChild(div);
    console.log(evt.data);
}
