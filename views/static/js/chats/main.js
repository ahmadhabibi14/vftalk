let chatInputElm = document.getElementById("chatInputElm");
let chatContainer = document.getElementById("chatContainer");
let sendChatBtnElm = document.getElementById("sendChatBtnElm");
let sendChatIcon = document.getElementById("sendChatIcon");
let loadingIcon = document.getElementById("loadingIcon");
let socket = new WebSocket(`wss://${window.location.host}/api/room`);

const username = localStorage.getItem("username");

function SendMessage() {
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
  return;
}

socket.onopen = () => {
  console.log("Connected");
};

socket.onmessage = (e) => {
  let data = JSON.parse(e.data);
  let rootMsg = document.createElement("chatroot")
  let msgContainer = document.createElement("chat");
  let unameElm = document.createElement("span");
  let msgElm = document.createElement("p");

  if (username === data.username) {
    rootMsg.className = "chat_root owned"
  } else {
    rootMsg.className = "chat_root"
  }
 
  msgContainer.className = "chat_item";
  unameElm.className = "chat_username";
  unameElm.textContent = data.username;
  msgElm.className = "chat_message";
  msgElm.textContent = data.message;


  chatContainer.appendChild(rootMsg);
  rootMsg.appendChild(msgContainer)
  msgContainer.appendChild(unameElm);
  msgContainer.appendChild(msgElm);
  msgContainer.scrollIntoView( {behavior: 'smooth'} );
};

socket.onerror = (e) => {
  alert("Error WebSocket connection: ", e);
  console.error("Error WebSocket connection: ", e);
}

sendChatBtnElm.addEventListener("click", () => {
  if (chatInputElm.value === "/clear") {
    chatInputElm.value = "";
    chatContainer.innerHTML = "";
    return;
  }
  SendMessage();
});

chatInputElm.addEventListener("keydown", (e) => {
  if (chatInputElm.value === "/clear") {
    chatInputElm.value = "";
    chatContainer.innerHTML = "";
    return;
  }
  if (e.key === "Enter") {
    SendMessage();
  }
});