let chatInputElm = document.getElementById("chatInputElm");
let chatContainer = document.getElementById("chatContainer");
let sendChatBtnElm = document.getElementById("sendChatBtnElm");
let sendChatIcon = document.getElementById("sendChatIcon");
let loadingIcon = document.getElementById("loadingIcon");
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
  msgElement.scrollIntoView( {behavior: 'smooth'} );
};

socket.onerror = (e) => {
  alert("Error WebSocket connection: ", e);
  console.error("Error WebSocket connection: ", e);
}

sendChatBtnElm.addEventListener("click", () => {
  if (chatInputElm.value === "") {
    return;
  }
  sendChatBtnElm.disabled = true;
  sendChatIcon.style.display = "none";
  loadingIcon.style.display = "block";
  socket.send(
    JSON.stringify({
      message: chatInputElm.value,
    })
  );
  chatInputElm.value = "";
  sendChatIcon.style.display = "block";
  loadingIcon.style.display = "none";
  sendChatBtnElm.disabled = false;
});

chatInputElm.addEventListener("keydown", (e) => {
  if (chatInputElm.value === "") {
    return;
  }
  if (e.key === "Enter") {
    sendChatBtnElm.disabled = true;
    sendChatIcon.style.display = "none";
    loadingIcon.style.display = "block";
    socket.send(
      JSON.stringify({
        message: chatInputElm.value,
      })
    );
    chatInputElm.value = "";
    sendChatIcon.style.display = "block";
    loadingIcon.style.display = "none";
    sendChatBtnElm.disabled = false;
  }
});