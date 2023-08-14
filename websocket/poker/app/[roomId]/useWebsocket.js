import { useState, useEffect } from "react";

export default function useWebSocket({ roomId, onMessage }) {
    var [conn, setConn] = useState(null);

    useEffect(() => {
        console.log('running effect, which creates a ws connection');
        conn = new WebSocket(`ws://localhost:8080/rooms/${roomId}`);

        console.log('init room: ' + roomId);
        conn.onclose = function (evt) {
          setConn(null);
          console.log(evt);
          console.log("ws closed");
        };
    
        conn.onopen = function (evt) {
          setConn(conn);
          console.log("ws connected");
        };
    
        conn.onmessage = function (evt) {
          onMessage(evt.data);
        };
    
        return function () {
          conn.close();
          console.log("cleanup connections.");
        };
    }, [roomId, onMessage]);

    return {
      name: roomId,
      send: function(m) {
        if (!conn) {
          console.log('conn is null');
          return;
        }
        conn.send(m);
      }
    }
}