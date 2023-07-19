import { cssBundleHref } from "@remix-run/css-bundle";
import type { LinksFunction } from "@remix-run/node";
import { Links, LiveReload, Meta, Outlet, Scripts, ScrollRestoration, } from "@remix-run/react";

import stylesheet from "@/tailwind.css";
import { useState } from "react";
import Sidebar, { STATE as SidebarState } from "@/components/ui/Sidebar";

export const links: LinksFunction = () => [
    {rel: "stylesheet", href: stylesheet},
    ...(cssBundleHref ? [{rel: "stylesheet", href: cssBundleHref}] : []),
    // NOTE: Architect deploys the public directory to /_static/
    {rel: "icon", href: "/_static/favicon.ico"},
];

export default function App() {
    const [sidebarState, setSidebarState] = useState(SidebarState.OPEN);

    return (
        <html lang="en" className="h-full">
        <head>
            <meta charSet="utf-8"/>
            <meta name="viewport" content="width=device-width,initial-scale=1"/>
            <Meta/>
            <Links/>
        </head>
        <body className="h-full">
        <div className="flex">
            <Sidebar
                state={sidebarState}
                toggleState={() => setSidebarState(sidebarState === SidebarState.OPEN ? SidebarState.CLOSED : SidebarState.OPEN)}
            />

            <main className="p-9">
                <Outlet/>
            </main>

        </div>
        <ScrollRestoration/>
        <Scripts/>
        <LiveReload/>
        </body>
        </html>
    );
}
