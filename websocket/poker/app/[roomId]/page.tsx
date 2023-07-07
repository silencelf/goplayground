"use client";

import Button from "../components/button";
import { useEffect, useState } from "react";
import UserName from "../components/userName";

function handleSizeClick(value: string | number) {
  alert(value);
}

type action = {
  roomId: string;
  userId: string;
  type: string;
  payload: object;
}

export default function Room({ params }: { params: { roomId: string } }) {
  const [user, setUser] = useState({ id: '' ,name: ''});
  const [hasUser, setHasUser] = useState(false);

  useEffect(() => {
    try {
      const tempUser = JSON.parse(localStorage['poker_user']);
      setUser(tempUser);
    } catch(e) {
      console.log(e);
    }
  }, [params.roomId]);

  useEffect(() => {
    var conn = new WebSocket("ws://localhost:8080/ws");
    conn.onclose = function (evt) {
      console.log(evt);
      console.log("ws closed");
    };

    conn.onopen = function (evt) {
      console.log("ws connected");
    };

    conn.onmessage = function (evt) {
      var messages = evt.data.split("\n");
      for (var i = 0; i < messages.length; i++) {
        console.log(messages[i]);
      }
    };

    return function() {
      console.log('cleanup connections.');
      conn.close();
    }
  }, [params.roomId]);

  function setUserName(name: string) {
    const user = { id: '', name };
    setUser(user);
    try {
      const tempUser = JSON.stringify(user);
      localStorage['poker_user'] = tempUser;
    } catch(e) {
      console.log(e);
    }
  }

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
    const shape = '♤♧♡♢'[i % 4];
    estimations.push({ name: "name" + i, vote: i, shape: shape });
  }

  return (
    <main className="flex min-h-screen flex-col items-cente justify-normal p-4 lg:p-24">
      {
        !hasUser && <UserName name={user?.name} onChange={setUserName}></UserName>
      }

      {/* <div className="w-full py-2 text-left">{params.roomId}</div> */}
      <div
        id="actions"
        className="grid w-full min-w-fit lg:grid-cols-[2fr,1fr] lg:text-left lg:mb-4"
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
          {estimations.map(({ name, vote, shape }) => (
            <li key={name} className="px-16 py-2">
              <div className="bg-gradient-to-br from-cyan-300 to-blue-300  w-24 h-36 rounded-lg text-center inline-block shadow-md text-poker hover:border">
                <span>{shape}</span>
              </div>
              <div className="text-center">{name}</div>
            </li>
          ))}
        </ul>
      </div>
    </main>
  );
}
