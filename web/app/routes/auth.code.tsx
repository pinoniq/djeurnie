import {LoaderArgs, json, redirect} from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import {getOAuthTokenFromCode} from "~/cognito/auth";
import invariant from "tiny-invariant";
import {commitSession, getSession} from "~/session";
import {setAccessToken} from "~/cognito/auth.session";

export async function loader({ request }: LoaderArgs) {
    const url = new URL(request.url);

    if (url.searchParams.has('error')) {
        return json({
            error: url.searchParams.get('error'),
            description: url.searchParams.get('error_description'),
        });
    }

    const code: string | null = url.searchParams.get('code');
    invariant(code, 'Missing authorization code');

    const tokens = await getOAuthTokenFromCode(code);
    invariant(tokens.access_token, 'Missing access token');

    const session = await getSession(request.headers.get('Cookie'));
    await setAccessToken(session, tokens.access_token);

    return redirect('/', {
        headers: {
            'Set-Cookie': await commitSession(session),
        },
    });
}

export default function Login() {
    const data = useLoaderData();

    return (
        <main>
            <pre>Error: {data.error}</pre>
            <pre>Description: {data.description}</pre>
        </main>
    );
}