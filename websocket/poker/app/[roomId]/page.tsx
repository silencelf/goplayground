"use client";

import { useCallback, useEffect, useLayoutEffect, useState } from "react";
import UserName from "../components/UserName";
import Button from "../components/Button";
import useWebSocket from "./useWebsocket";

type action = {
  roomId: string;
  userId: string;
  type: string;
  payload: object;
};

export default function Room({ params }: { params: { roomId: string } }) {
  const [user, setUser] = useState({ id: '', name: '' });
  const [userLoaded, setUserLoaded] = useState(false);

  // define hub as state for displaying the votes
  const [hub, setHub] = useState({ isUnveiled: false, clients: [] });
  const [clientId, setClientId] = useState('');

  useEffect(() => {
    if (!userLoaded) {
      console.log('loading user...');
      if (localStorage["poker_user"]) {
        try {
          const user = JSON.parse(localStorage["poker_user"])
          setUser(user);
          setUserLoaded(true);
        } catch (e) {
          console.log(e);
        }
      }
    }
    return () => {
      setUserLoaded(false);
      console.log('cleanup phase 1');
    };
  }, []);


  const onMessage = useCallback(function (m: string) {
    // split the message by '\n' and then filter out the empty string
    // foreach splitted message and handle them
    m.split('\n').filter((m) => m).forEach((m) => {
      // deserialize the payload
      const payload = JSON.parse(m);
      console.log(payload)
      // if the response is a room, set the hub state
      if (payload.type === 'room') {
        // sort the clients by id to make sure the order is consistant
        payload.value.clients.sort((a: any, b: any) => {
          return a.id.localeCompare(b.id);
        });
        // assign a shape to each client and then assign to hub
        // using for loop, use the index to assign a shape
        for (let i = 0; i < payload.value.clients.length; i ++) {
          const shape = '♤♧♡♢'[i % 4];
          payload.value.clients[i].shape = shape;
        }

        setHub(payload.value);
      }

      if (payload.type === 'id') {
        setClientId(payload.value);
      }
    });

  }, [])

  const room = useWebSocket({
    roomId: params.roomId,
    onMessage: onMessage,
    onConnected: () =>  {
      console.log('ws connected.....');
    }
  });

  function handleSizeClick(value: string | number) {
    room.send(`/vote ${value}`);
  }
  
  function handleClearClick() {
    room.send('/clear');
  }

  function saveUser({ id, name }: any){
    user.name = name;
    try {
      const tempUser = JSON.stringify(user);
      localStorage["poker_user"] = tempUser;
      room.send(`/nick ${name}`);
    } catch (e) {
      console.log(e);
    }
    setUser(user);
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

  function showVotes(): void {
    // send websocket requset to set the hub as unveiled
    // And then fetch and display all the votes
    room.send('/unveil');
  }

  return (
    <main className="flex min-h-screen flex-col items-cente justify-normal p-4 lg:p-24">
      {userLoaded && (
        <UserName
          userName={user?.name}
          onConfirmClick={(name) => saveUser({ name })}
        />
      )}

      <div className="w-full py-2 text-left">
        <label htmlFor="clientId">Client Id:</label>
        <span className="ml-2" id="clientId">{clientId}</span></div>
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
            <Button onClick={() => showVotes()}>Show</Button>
          </li>
          <li>
            <Button onClick={() => handleClearClick()}>Clear</Button>
          </li>
        </ul>
      </div>

      <div id="pokers-container" className="w-full lg:mt-5 mb-32 lg:mb-0">
        <ul className="flex flex-wrap justify-center">
          {hub.clients.map(({ id, nick, vote, shape }) => (
            <li key={id} className="px-16 py-2">
              <div className="bg-gradient-to-br from-cyan-300 to-blue-300  w-24 h-36 rounded-lg text-center inline-block shadow-md text-poker hover:border">
                {
                  // display the vote in bold if the hub is unveiled, otherwise display the shape
                  hub.isUnveiled && vote.hasValue ? (
                    <span>{vote.v}</span>
                  ) : (
                    <span>{shape}</span>
                  )
                }
              </div>
              <div className="text-center">{nick}</div>
            </li>
          ))}
        </ul>
      </div>
    </main>
  );
}
