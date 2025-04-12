export function Button({
  children,
  className,
  ...props
}: {
  children: React.ReactNode;
} & React.ButtonHTMLAttributes<any>) {
  return (
    <button
      className={`flex self-end items-center gap-2 w-fit p-2 text-black font-bold bg-purple-400 border-2 border-purple-600 rounded-lg cursor-pointer ${className}`}
      {...props}
    >
      {children}
    </button>
  );
}
