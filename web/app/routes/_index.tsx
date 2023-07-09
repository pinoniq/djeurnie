import type {V2_MetaFunction} from "@remix-run/node";
import {json, LoaderArgs} from "@remix-run/node";
import {Link, useLoaderData} from "@remix-run/react";

import {
    requireUserSession
} from "@/cognito/auth.session";

export const meta: V2_MetaFunction = () => [{title: "DJEURNIE Web"}];

export async function loader({request}: LoaderArgs) {
    const [
        ,
        userSession,
    ] = await requireUserSession(request);

    return json({
        session: {
            payload: userSession,
        }
    });
}

export default function Index() {
    const data = useLoaderData<typeof loader>();

    return (
        <>
            <ul>
                <li><Link to="/login">Login</Link></li>
                <li><Link to="/logout">Logout</Link></li>
            </ul>
            <pre>Sub: {data.session.payload.email}</pre>
        </>

    );
}
