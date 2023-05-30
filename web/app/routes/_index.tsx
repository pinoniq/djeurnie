import type {V2_MetaFunction} from "@remix-run/node";
import {json, LoaderArgs} from "@remix-run/node";
import {Link, useLoaderData} from "@remix-run/react";

import {getSession} from "~/session";
import {getAccessToken, getAccessTokenPayload} from "~/cognito/auth.session";

export const meta: V2_MetaFunction = () => [{title: "DJEURNIE Web"}];

export async function loader({request}: LoaderArgs) {
    const session = await getSession(request.headers.get('Cookie'));
    const accessToken = getAccessToken(session);
    let accessTokenPayload = {}
    if (accessToken) {
        accessTokenPayload = await getAccessTokenPayload(session);
    }

    return json({
        session: {
            accessToken,
            payload: accessTokenPayload,
        }
    })
}

export default function Index() {
    const data = useLoaderData();

    return (
        <main className="relative">
            <ul>
                <li><Link to="/login">Login</Link></li>
                <li><Link to="/login">Logout</Link></li>
            </ul>
            {data.session.accessToken && (
                <pre>{data.session.accessToken}</pre>
            )}
            {data.session.payload && (
                <>
                    <pre>Sub: {data.session.payload.sub}</pre>
                    <pre>exp: {data.session.payload.exp}</pre>
                </>
            )}
        </main>
    );
}
