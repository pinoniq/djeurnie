import { ArrowRight, Combine, FileDigit, Sunrise, Sunset } from "lucide-react";

import { SidebarProps } from "@/components/ui/Sidebar/types";
import { STATE } from "@/components/ui/Sidebar/constants";
import Logo from "@/components/ui/Logo";
import { Link } from "@remix-run/react";

const MenuItems = [
    {
        title: 'Ingress',
        icon: Sunset,
        to: '/ingress',
    },
    {
        title: 'Flows',
        icon: Combine,
        to: '/flows',
    },
    {
        title: 'Egress',
        icon: Sunrise,
        to: '/egress',
    },
    {
        title: 'Data catalog',
        icon: FileDigit,
        to: '/data-catalog',
    },
];

export default function Sidebar({
    state,
    toggleState,
}: SidebarProps) {
    const open: boolean = state === STATE.OPEN;

    return (
        <div className={`h-screen bg-bg-light p-5 pt-8 relative ${open ? 'w-72' : 'w-20'} duration-200`}>
            <ArrowRight
                className={`bg-white text-blue-450 text-3xl rounded-full absolute -right-3 top-9 cursor-pointer ${open && 'rotate-180'}`}
                onClick={toggleState}
            />

            <Link to="/">
                <div className="inline-flex">
                <span className="text-2xl block -ml-1">
                    <Logo className="text-blue-450" />
                </span>
                    <h1 className={`text-blue-450 origin-left font-heading font-normal pl-4 text-xl duration-200 ${!open ? 'scale-0' : ''}`}>
                        DJEURNIE
                    </h1>
                </div>
            </Link>

            <nav>
                <ul className="pt-8">
                    {MenuItems.map((menuItem, index : number) => (
                        <Link to={menuItem.to} key={index}>
                            <li
                                className="text-gray-900 text-sm flex items-center gap-x-4 cursor-pointer p-2 mt-4 hover:bg-blue-10 rounded-md"
                            >
                                <span className="text-base block float-left"><menuItem.icon /></span>
                                <span className={`text-base font-medium whitespace-nowrap flex-1 duration-200 ${!open ? 'hidden' : ''}`}>
                                    {menuItem.title}
                                </span>
                            </li>
                        </Link>
                    ))}
                </ul>
            </nav>
        </div>
    )
}
