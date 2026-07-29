import type { ReactNode } from "react";

export interface NavItem {
  label: string;
  href: string;
}

export interface NavbarCTA {
  label: string;
  href: string;
}

export interface NavbarProps {
  logo?: ReactNode;
  navigation: NavItem[];
  cta?: NavbarCTA;
}