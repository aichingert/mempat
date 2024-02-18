const GREEN = "rgb(195, 232, 141)";
const ERROR = "rgb(199, 146, 234)";

let game = document.getElementById("game");

for (let i = 0; i < 4; i++) {
    let row = document.createElement("div");
    row.classList.add("square-row");

    for (let j = 0; j < 4; j++) {
        let square = document.createElement("div");
        square.id = `${i} ${j}`;
        square.classList.add("square")

        square.addEventListener("click", function() {
            if (this.style.backgroundColor != GREEN) {
                socket.send(this.id);
            }
        });

        row.appendChild(square);
    }

    game.append(row);
}

let socket = new WebSocket("ws://" + document.location.host + "/ws");

socket.onerror = error => { console.log("Socket Error: ", error); };

socket.onmessage = function(evt) {
    message = evt.data.split(':')

    if (message.length == 1) {
        document.getElementById(evt.data).style.backgroundColor = GREEN;
        return;
    }

    switch (message[0]) {
        case "open":
            open = message[1].split(',').filter((e) => e !== "")

            for (let i = 0; i < open.length; i++) {
                document.getElementById(open[i]).style.backgroundColor = GREEN;
            }

            break;
        case "val":
        case "inv": 
            document.getElementById(message[1]).style.backgroundColor = message[0] === "val" ? GREEN : ERROR;
            break;
        default:
            console.log("message: ", message);
    }
}
