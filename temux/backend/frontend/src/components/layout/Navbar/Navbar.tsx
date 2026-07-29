import { Link } from "react-router-dom";

import Button from "../../ui/Button/Button";
import type { NavbarProps } from "./Navbar.types";

export default function Navbar({
  logo,
  navigation,
  cta,
}: NavbarProps) {
  return (
    <header className="sticky top-0 z-50 border-b border-slate-200 bg-white">
      <div className="mx-auto flex h-20 max-w-7xl items-center justify-between px-6">
        {/* Logo */}
        <div className="flex items-center">
          {logo}
        </div>

        {/* Navigation */}
        <nav className="hidden items-center gap-8 md:flex">
          {navigation.map((item) => (
            <Link
              key={item.label}
              to={item.href}
              className="font-medium text-slate-600 transition-colors hover:text-blue-600"
            >
              {item.label}
            </Link>
          ))}
        </nav>

        {/* CTA */}
        <div>
          {cta && (
            <Link to={cta.href}>
              <Button>
                {cta.label}
              </Button>
            </Link>
          )}
        </div>
      </div>
    </header>
  );
}