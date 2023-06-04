import type {V2_MetaFunction} from "@remix-run/node";
import {json, LoaderArgs} from "@remix-run/node";
import {Link, useLoaderData} from "@remix-run/react";

import {getSessionFromRequest} from "@/session";
import {getAccessToken, getAccessTokenPayload, getId, SessionIdData} from "@/cognito/auth.session";
import Sidebar, {
    STATE as SidebarState
} from "@/components/ui/Sidebar";
import { useState } from "react";

export const meta: V2_MetaFunction = () => [{title: "DJEURNIE Web"}];

export async function loader({request}: LoaderArgs) {
    const session = await getSessionFromRequest(request);
    const accessToken = getAccessToken(session);
    if (!accessToken) {
        return json({
            session: {
                accessToken,
                payload: null,
            }
        });

    }

    const idPayload: SessionIdData = await getId(session);
    return json({
        session: {
            accessToken,
            payload: idPayload || null,
        }
    });
}

export default function Index() {
    const data = useLoaderData<typeof loader>();
    const [sidebarState, setSidebarState] = useState(SidebarState.OPEN);


    return (
        <div className="flex">

            <Sidebar
                state={sidebarState}
                toggleState={() => setSidebarState(sidebarState === SidebarState.OPEN ? SidebarState.CLOSED : SidebarState.OPEN)} />

            <main className="p-7">
                <ul>
                    <li><Link to="/login">Login</Link></li>
                    <li><Link to="/login">Logout</Link></li>
                </ul>
                {data.session.accessToken && (
                    <pre>{data.session.accessToken}</pre>
                )}
                {data.session.payload && (
                    <>
                        <pre>Sub: {data.session.payload.email}</pre>
                    </>
                )}
            </main>
        </div>

    );
}
