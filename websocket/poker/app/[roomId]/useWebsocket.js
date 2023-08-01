import { useEffect } from "react";

export default function useWebSocket({ roomId, onMessage }) {
    useEffect(() => {
        var conn = new WebSocket("ws://localhost:8080/ws");

        console.log('init room: ' + roomId);
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
          onMessage(messages);
        };
    
        return function () {
          console.log("cleanup connections.");
          conn.close();
        };
    }, [roomId, onMessage]);
}