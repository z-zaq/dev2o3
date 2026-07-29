import type { ReactNode } from "react";

interface PublicLayoutProps {
  children: ReactNode;
}

export default function PublicLayout({
  children,
}: PublicLayoutProps) {
  return (
    <div className="min-h-screen flex flex-col bg-white">
      <main className="flex-1">
        {children}
      </main>
    </div>
  );
}