let socket = new WebSocket("ws://127.0.0.1:8080/ws");

console.log("Attempting Connection...");

socket.onopen = () => {
    console.log("Successfully Connected");
    socket.send("Hi From the Client!")
};

socket.onclose = event => {
    console.log("Socket Closed Connection: ", event);
    socket.send("Client Closed!")
};

socket.onerror = error => { console.log("Socket Error: ", error); };

socket.onmessage = function(evt) {
    console.log(message)
}

let game = document.getElementById("game");

for (let i = 0; i < 4; i++) {
    let row = document.createElement("div");
    row.classList.add("square-row");

    for (let j = 0; j < 4; j++) {
        let square = document.createElement("div");
        square.id = `${i} ${j}`;
        square.classList.add("square")

        square.addEventListener("click", function() {
            console.log(this.id);
            this.style.backgroundColor = "blue";
        });

        row.appendChild(square);
    }

    game.append(row);
}

