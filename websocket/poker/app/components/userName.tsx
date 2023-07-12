import { useState } from "react";

type UserNameProps = {
  onConfirmClick: (name: string) => void;
  name: string;
};


export default function UserName(props: UserNameProps) {
  const [name, setName] = useState(props.name)

  return (
    <div className="px-3 py-2">
      Please enter your name:
      <p>
        <input
          id="userName"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
        <button onClick={(e) => props.onConfirmClick(name)}>Confirm</button>
      </p>
    </div>
  );
}
