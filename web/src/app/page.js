"use client";

import { useEffect } from "react";

export default function Home() {
  let inputMsg = "";
  let socket = null;
  let messages = [];

  useEffect(() => {
    socket = new WebSocket("ws://localhost:8080/room");
    socket.onclose = function () {
      console.log("closed");
    };
    socket.onmessage = function (e) {
      messages.push(e.data);
    };
  });

  
  function handleSubmit() {
    socket.send(inputMsg);
    return;
  }

  return (
    <div className="mx-auto mt-20 bg-zinc-950 shadow-lg rounded-xl p-6 w-[450px] text-zinc-100">
      <div className="flex flex-col">
        {messages.map((msg) => (
          <p>{msg}</p>
        ))}
      </div>
      <form className="flex flex-col gap-5">
        <textarea
          value={inputMsg}
          onChange={(e) => (inputMsg = e.target.value)}
          name="post"
          id="post"
          placeholder="What's your feeling..."
          className="focus:outline-2 outline-sky-500 bg-transparent p-3 rounded-xl border border-sky-500"
        ></textarea>
        <div className="flex flex-row justify-end">
          <button
            onClick={handleSubmit}
            className="w-fit h-fit py-2 px-7 rounded-full text-white bg-sky-500 cursor-pointer hover:bg-sky-600"
          >
            Post
          </button>
        </div>
      </form>
    </div>
  );
}
