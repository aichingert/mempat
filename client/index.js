const GRAY  = "rgb(166, 172, 205)";
const GREEN = "rgb(195, 232, 141)";
const RED   = "rgb(240, 113, 120)";

const ROWS = 5;
const COLS = 5;

let game = document.getElementById("game");

for (let i = 0; i < ROWS; i++) {
    let row = document.createElement("div");
    row.classList.add("square-row");

    for (let j = 0; j < COLS; j++) {
        let square = document.createElement("div");
        square.id = `${i} ${j}`;
        square.style.backgroundColor = GRAY;
        square.classList.add("square")

        square.addEventListener("click", function() {
            if (!blockSend && this.style.backgroundColor === GRAY) {
                socket.send(this.id);
            }
        });

        row.appendChild(square);
    }

    game.append(row);
}

let socket = new WebSocket("ws://" + document.location.host + "/ws");
let blockSend = false;

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
                document.getElementById(open[i].substring(1)).style.backgroundColor = open[i][0] == 'y' ? GREEN : RED;
            }

            break;
        case "new":
            for (let row of game.children) {
                for (let square of row.children) {
                    square.style.backgroundColor = GRAY;
                    square.classList.remove("flip");
                }
            }

            blockSend = true;
            open = message[1].split(',').filter((e) => e !== "");

            for (let i = 0; i < open.length; i++) {
                let square = document.getElementById(open[i]);

                square.classList.add("preview");
                setTimeout(() => { 
                    square.classList.remove("preview"); 
                    blockSend = false;
                }, 1250);
            }

            break;
        case "val":
        case "inv": 
            let square = document.getElementById(message[1]);
            square.style.backgroundColor = message[0] === "val" ? GREEN : RED;
            square.classList.add("flip");
            break;
        default:
            console.log("message: ", message);
    }
}
