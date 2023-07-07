import { useState } from "react";

type UserNameProps = {
  onConfirmClick?: () => void;
  onChange: (name: string) => void;
  name: string;
};

export default function UserName(props: UserNameProps) {
  return (
    <div className="px-3 py-2">
      Please enter your name:
      <p>
        <input
          id="userName"
          value={props.name}
          onChange={(e) => props.onChange(e.target.value)}
        />
        <button onClick={props.onConfirmClick}>Confirm</button>
      </p>
    </div>
  );
}
