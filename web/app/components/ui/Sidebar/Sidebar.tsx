import { ArrowRight, Combine, FileDigit, Sunrise, Sunset } from "lucide-react";

import { SidebarProps } from "@/components/ui/Sidebar/types";
import { STATE } from "@/components/ui/Sidebar/constants";
import Logo from "@/components/ui/Logo";
import { Link, NavLink } from "@remix-run/react";
import { clsx } from "clsx";

const MenuItems = [
  {
    title: "Ingress",
    icon: Sunset,
    to: "/ingress",
  },
  {
    title: "Flows",
    icon: Combine,
    to: "/flows",
  },
  {
    title: "Egress",
    icon: Sunrise,
    to: "/egress",
  },
  {
    title: "Data catalog",
    icon: FileDigit,
    to: "/data-catalog",
  },
];

function navLinkStyle({ isActive }: { isActive: boolean }): string {
  return clsx(
    "text-sm flex items-center gap-x-4 cursor-pointer p-2 mt-4 hover:bg-blue-50 rounded-md",
    {
      "text-blue-450": isActive,
      "text-gray-900": !isActive,
    }
  );
}

export default function Sidebar({ state, toggleState }: SidebarProps) {
  const open: boolean = state === STATE.OPEN;

  return (
    <div
      className={clsx(
        "relative h-screen grow-0 bg-bg-light p-5 pt-8 duration-200",
        {
          "w-72": open,
          "w-20": !open,
        }
      )}
    >
      <ArrowRight
        className={clsx(
          "absolute -right-3 top-9 cursor-pointer rounded-full bg-white text-3xl text-blue-450",
          {
            "rotate-180": open,
          }
        )}
        onClick={toggleState}
      />
      <Link to="/">
        <div className="inline-flex">
          <span className="-ml-1 block text-2xl">
            <Logo className="text-blue-450" />
          </span>
          <h1
            className={clsx(
              "origin-left pl-4 font-heading text-xl font-normal text-blue-450 duration-200",
              {
                "scale-0": !open,
              }
            )}
          >
            DJEURNIE
          </h1>
        </div>
      </Link>

      <nav className="pt-8">
        {MenuItems.map((menuItem, index: number) => (
          <NavLink to={menuItem.to} key={index} className={navLinkStyle}>
            <span className="float-left block text-base">
              <menuItem.icon />
            </span>
            <span
              className={clsx(
                "flex-1 whitespace-nowrap text-base font-medium duration-200",
                {
                  hidden: !open,
                }
              )}
            >
              {menuItem.title}
            </span>
          </NavLink>
        ))}
      </nav>
    </div>
  );
}
