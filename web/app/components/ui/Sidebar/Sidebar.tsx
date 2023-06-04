import { ArrowRight, Binary } from "lucide-react";

import { SidebarProps } from "@/components/ui/Sidebar/types";
import { STATE } from "@/components/ui/Sidebar/constants";
import Logo from "@/components/ui/Logo";

const MenuItems = [
    {
        title: 'Ingress',
        icon: Binary,
    }
];

export default function Sidebar({
    state,
    toggleState,
}: SidebarProps) {
    const open: boolean = state === STATE.OPEN;

    return (
        <div className={`h-screen bg-bg-light p-5 pt-8 relative ${open ? 'w-72' : 'w-20'} duration-200`}>
            <ArrowRight
                className={`bg-white text-green-900 text-3xl rounded-full absolute -right-3 top-9 cursor-pointer ${open && 'rotate-180'}`}
                onClick={() => toggleState()}
            />

            <div className="inline-flex">
                <span className="text-2xl block">
                    <Logo className="text-blue-450" />
                </span>
                <h1 className={`text-blue-450 origin-left font-heading font-normal text-2xl duration-200 ${!open && 'scale-0'}`}>
                    DJEURNIE
                </h1>
            </div>

            <ul className="pt-2">
                {MenuItems.map((menuItem, index : number) => (
                    <li
                        key={index}
                        className="text-gray-900 text-sm flex items-center gap-x-4 cursor-pointer p-2 mt-2 hover:bg-blue-50 rounded-md"
                    >
                        <span className="text-2xl block"><menuItem.icon /></span>
                        <span className={`text-base font-medium flex-1 duration-200 ${!open && 'scale-0'}`}>
                            {menuItem.title}
                        </span>
                    </li>
                ))}
            </ul>
        </div>
    )
}
