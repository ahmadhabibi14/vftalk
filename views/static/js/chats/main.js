const chatInputElm = document.getElementById('chatInputElm');
const chatContainer = document.getElementById('chatContainer');
const sendChatBtnElm = document.getElementById('sendChatBtnElm');
const sendChatIcon = document.getElementById('sendChatIcon');
const loadingIcon = document.getElementById('loadingIcon');

const url = new URL(window.location.href);
let socketURL = `ws://${window.location.host}/api/room`;
if ((url.protocol).includes('s')) socketURL = `wss://${window.location.host}/api/room`;
let socket = new WebSocket(socketURL);

const username = localStorage.getItem('username');

function SendMessage() {
  if (chatInputElm.value === '') return;

  sendChatBtnElm.disabled = true;
  sendChatIcon.style.display = 'none';
  loadingIcon.style.display = 'block';

  try { socket.send( JSON.stringify({ message: chatInputElm.value }) )
  } catch (e) { notifier.showError('Error sending message: ', e) }

  chatInputElm.value = '';
  sendChatIcon.style.display = 'block';
  loadingIcon.style.display = 'none';
  sendChatBtnElm.disabled = false;
  return;
}

socket.onopen = () => console.log('Connected to the Room Chat');

socket.onmessage = (e) => {
  let data = JSON.parse(e.data);
  let rootMsg = document.createElement('chatroot')
  let msgContainer = document.createElement('chat');
  let unameElm = document.createElement('span');
  let msgElm = document.createElement('p');

  if (username === data.username) rootMsg.className = 'chat_root owned';
  else rootMsg.className = 'chat_root';
 
  msgContainer.className = 'chat_item';
  unameElm.className = 'chat_username';
  unameElm.textContent = data.username;
  msgElm.className = 'chat_message';
  msgElm.textContent = data.message;

  chatContainer.appendChild(rootMsg);
  rootMsg.appendChild(msgContainer)
  msgContainer.appendChild(unameElm);
  msgContainer.appendChild(msgElm);
  msgContainer.scrollIntoView( {behavior: 'smooth'} );
};

socket.onerror = (e) => notifier.showError('Error WebSocket connection: ', e);

socket.onclose = (e) => {
  if (e.wasClean) notifier.showWarning(`Connection closed cleanly, code=${e.code}, reason=${e.reason}`);
  else notifier.showWarning('Connection abruptly closed');

  setTimeout(() => window.location.reload(), 2000);
}

sendChatBtnElm.addEventListener('click', () => {
  if (chatInputElm.value === '/clear') {
    chatInputElm.value = '';
    chatContainer.innerHTML = '';
    return;
  }
  SendMessage();
});

chatInputElm.addEventListener('keydown', (e) => {
  if (chatInputElm.value === '/clear') {
    chatInputElm.value = '';
    chatContainer.innerHTML = '';
    return;
  }
  if (e.key === 'Enter') SendMessage();
});