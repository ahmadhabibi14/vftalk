export default function Home() {
  let webSocket = new WebSocket("ws://localhost:3000/ws/123?v=1.0");
  webSocket.onmessage = function (e) {
    console.log(e);
  };
  return <div>tess</div>;
}
