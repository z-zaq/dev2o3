import type { ButtonProps } from "./Button.types";

const variantClasses = {
  primary:
    "bg-blue-600 text-white hover:bg-blue-700",

  secondary:
    "bg-slate-900 text-white hover:bg-slate-800",

  outline:
    "border border-slate-300 bg-white text-slate-900 hover:bg-slate-100",

  ghost:
    "bg-transparent text-slate-900 hover:bg-slate-100",
};

const sizeClasses = {
  sm: "px-3 py-2 text-sm",

  md: "px-5 py-3 text-base",

  lg: "px-7 py-4 text-lg",
};

export default function Button({
  children,
  variant = "primary",
  size = "md",
  className = "",
  ...props
}: ButtonProps) {
  return (
    <button
      className={`
        inline-flex
        items-center
        justify-center
        rounded-xl
        font-medium
        transition-all
        duration-200
        focus:outline-none
        focus:ring-2
        focus:ring-blue-500
        disabled:opacity-50
        disabled:cursor-not-allowed
        ${variantClasses[variant]}
        ${sizeClasses[size]}
        ${className}
      `}
      {...props}
    >
      {children}
    </button>
  );
}