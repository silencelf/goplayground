import { ReactNode, memo } from "react";

type ButtonProps = {
  onClick?: () => void;
  children: ReactNode;
};

export default memo(function Button({ onClick, children }: ButtonProps) {
  return (
    <button
      onClick={onClick}
      className="bg-blue-400 rounded-lg px-3 py-2 text-slate-700 font-medium hover:bg-slate-100 hover:text-slate-900 shadow"
    >
      {children}
    </button>
  );
});
