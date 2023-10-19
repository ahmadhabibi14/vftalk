let input = document.getElementById("input");
let output = document.getElementById("output");
let socket = new WebSocket("ws://" + window.location.host + "/room");

socket.onopen = () => {
  output.innerHTML += "Status: connected\n";
};
socket.onmessage = (e) => {
  console.log("Message: ", e.data);
  output.innerHTML += "Message from server: " + e.data + "\n";
};

function send() {
  socket.send(
    JSON.stringify({
      message: input.value,
    })
  );
  input.value = "";
}
