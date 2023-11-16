import { useState } from "react";

type UserNameProps = {
  onConfirmClick: (name: string) => void;
  userName: string;
};

export default function UserName({ userName, onConfirmClick }: UserNameProps) {
  const [name, setName] = useState(userName);

  return (
    <div className="px-3 py-2">
      Please enter your name:
      <p>
        <input
          id="userName"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
        <button className="ml-2" onClick={() => onConfirmClick(name)}>Confirm</button>
      </p>
    </div>
  );
}
