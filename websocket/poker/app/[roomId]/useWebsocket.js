import { useEffect } from "react";

export default function useWebSocket({ roomId, onMessage }) {
    var conn;
    useEffect(() => {
        conn = new WebSocket(`ws://localhost:8080/rooms/${roomId}`);

        console.log('init room: ' + roomId);
        conn.onclose = function (evt) {
          console.log(evt);
          console.log("ws closed");
        };
    
        conn.onopen = function (evt) {
          console.log("ws connected");
        };
    
        conn.onmessage = function (evt) {
          onMessage(evt.data);
        };
    
        return function () {
          console.log("cleanup connections.");
          conn.close();
        };
    }, [roomId, onMessage]);

    return {
      name: roomId,
      send: function(m) {
        if (!conn) return;
        conn.send(m);
      }
    }
}