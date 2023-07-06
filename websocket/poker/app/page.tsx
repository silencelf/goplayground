"use client";

import Image from "next/image";
import Button from "./components/button";
import { useEffect } from "react";

function handleSizeClick(value: string | number) {
  alert(value);
}

export default function Home() {
  useEffect(() => {
    var conn = new WebSocket("ws://localhost:3000" + "/api/");
    conn.onclose = function (evt) {
      console.log(evt);
      console.log('ws closed');
    };

    conn.onopen = function(evt) { console.log('ws connected')};

    conn.onmessage = function (evt) {
        var messages = evt.data.split('\n');
        for (var i = 0; i < messages.length; i++) {
          console.log(messages[i]);
        }
    };
  }, []);


  const points = [
    ["0", 0],
    ["1", 1],
    ["2", 2],
    ["3", 3],
    ["5", 5],
    ["8", 8],
    ["13", 13],
    ["21", 21],
    ["?", "?"],
  ];

  const estimations = [];
  for (let i = 0; i < 20; i++) {
    estimations.push({ name: "name" + i, vote: i });
  }

  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-4 lg:p-24">
      <div
        id="actions"
        className="grid w-full min-w-fit lg:grid-cols-[2fr,1fr] lg:text-left"
      >
        <ul className="flex bg-gradient-to-br from-blue-500 to-gray-500 rounded-lg shadow">
          {points.map(([title, val]) => (
            <li key={title} className="px-2">
              <button
                onClick={() => handleSizeClick(val)}
                className="rounded-lg px-3 py-2 text-slate-700 font-medium hover:bg-slate-100 hover:text-slate-900 transition-colors"
              >
                {title}
              </button>
            </li>
          ))}
        </ul>
        <ul className="flex justify-end my-5 lg:my-0">
          <li className="px-5">
            <Button onClick={() => alert(1)}>Show</Button>
          </li>
          <li>
            <Button>Clear</Button>
          </li>
        </ul>
      </div>

      <div id="pokers-container" className="w-full lg:mt-5 mb-32 lg:mb-0">
        <ul className="flex flex-wrap justify-center">
          {estimations.map(({ name, vote }) => (
            <li key={name} className="px-16 py-2">
              <div className="bg-gradient-to-br from-cyan-300 to-blue-300  w-24 h-36 rounded-lg text-center inline-block shadow-md text-poker hover:border">
                <span>{vote}</span>
              </div>
              <div className="text-center">{name}</div>
            </li>
          ))}
        </ul>
      </div>

    </main>
  );
}
