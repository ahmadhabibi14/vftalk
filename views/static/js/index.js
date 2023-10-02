let socket = null;

document.addEventListener("DOMContentLoaded", function () {
  const msgForm = document.getElementById("message_form");
  const msgInput = document.getElementById("msg_input");
  msgForm.addEventListener("submit", (e) => {
    e.preventDefault();
    socket.send(msgInput.value);
  });
});

(() => {
  socket = new WebSocket("ws://localhost:8080/room");
  socket.onmessage = function (e) {
    var msgElement = document.createElement("chat");
    msgElement.className = "chat";
    msgElement.textContent = e.data;
    const chatContainer = document.getElementById("chat_container");
    chatContainer.appendChild(msgElement);
  };
  socket.onclose = function () {
    console.log("closed");
  };
})();
