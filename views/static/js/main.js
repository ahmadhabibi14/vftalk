let chatInputElm = document.getElementById("chatInputElm");
let chatContainer = document.getElementById("chatContainer");
let sendChatBtnElm = document.getElementById("sendChatBtnElm");
let socket = new WebSocket("ws://localhost:8080/room");

socket.onopen = () => {
  console.log("Connected");
};

socket.onmessage = (e) => {
  let data = JSON.parse(e.data);
  let msgElement = document.createElement("chat");
  msgElement.className = "chat_item";
  msgElement.textContent = data.message;
  chatContainer.appendChild(msgElement);
};

sendChatBtnElm.addEventListener("click", () => {
  if (chatInputElm.value === "") {
    return;
  }
  socket.send(
    JSON.stringify({
      message: chatInputElm.value,
    })
  );
  chatInputElm.value = "";
});

chatInputElm.addEventListener("keydown", (e) => {
  if (e.key === "Enter") {
    socket.send(
      JSON.stringify({
        message: chatInputElm.value,
      })
    );
    chatInputElm.value = "";
  }
});