let socket = null;

window.onload = function () {
  socket = new WebSocket("ws://localhost:8080/room");
  socket.onmessage = function (e) {
    console.log(e.data);
    var msgElement = document.createElement("chat");
    msgElement.className = "chat";
    msgElement.textContent = e.data;
    const chatContainer = document.getElementById("chat_container");
    chatContainer.appendChild(msgElement);
  };
  socket.onclose = function () {
    console.log("closed");
  };

  const msgForm = document.getElementById("message_form");
  const msgInput = document.getElementById("msg_input");
  msgForm.addEventListener("submit", (e) => {
    e.preventDefault();
    socket.send(msgInput.value);
  });
};
  