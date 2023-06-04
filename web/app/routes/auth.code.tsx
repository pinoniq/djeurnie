import {LoaderArgs, json, redirect} from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import invariant from "tiny-invariant";
import {commitSession, getSessionFromRequest} from "@/session";
import {setUserSessionFromCode} from "@/cognito/auth.session";

export async function loader({ request }: LoaderArgs) {
    const url = new URL(request.url);

    if (url.searchParams.has('error')) {
        return json({
            error: url.searchParams.get('error'),
            description: url.searchParams.get('error_description'),
        });
    }

    const session = await getSessionFromRequest(request);
    const code: string | null = url.searchParams.get('code');
    invariant(code, 'Missing authorization code');

    await setUserSessionFromCode(session, code);

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