import { useState, useEffect } from "react";

export default function useWebSocket({ roomId, onMessage, onConnected }) {
    var [conn, setConn] = useState(null);
    let room = {
      name: roomId,
      send: function(m) {
        if (!conn) {
          console.log('conn is null');
          return;
        }
        conn.send(m);
      }
    };

    useEffect(() => {
        console.log('running effect, which creates a ws connection');
        conn = new WebSocket(`ws://localhost:8080/rooms/${roomId}`);

        console.log('init room: ' + roomId);
        conn.onclose = function (evt) {
          setConn(null);
          console.log(evt);
          console.log("ws closed");

          // reconnect after 5 seconds
          setTimeout(() => {
            console.log('reconnecting...');
            conn = new WebSocket(`ws://localhost:8080/rooms/${roomId}`);
            hookEvents();
          }, 5000);
        };
    
        function hookEvents() {
          conn.onopen = function (evt) {
            setConn(conn);
            onConnected();
          };

          conn.onerror = function (evt) {
            console.log(evt);
          };
      
          conn.onmessage = function (evt) {
            onMessage(evt.data);
          };
        }

        hookEvents();
    
        return function () {
          conn.close();
          console.log("cleanup connections.");
        };
    }, [roomId, onMessage]);

    return room;
}