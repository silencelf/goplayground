import { useState } from "react";

type UserNameProps = {
  onConfirmClick: (name: string) => void;
  name: string;
};

export default function UserName({ name, onConfirmClick }: UserNameProps) {
  const [nameValue, setName] = useState(name)

  return (
    <div className="px-3 py-2">
      Please enter your name:
      <p>
        <input
          id="userName"
          value={nameValue}
          onChange={(e) => setName(e.target.value)}
        />
        <button onClick={(e) => onConfirmClick(nameValue)}>Confirm</button>
      </p>
    </div>
  );
}
