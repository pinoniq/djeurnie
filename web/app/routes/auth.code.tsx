import { LoaderArgs, json } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import {getOAuthTokenFromCode} from "~/cognito/auth";
import invariant from "tiny-invariant";

export async function loader({ request }: LoaderArgs) {
    const url = new URL(request.url);

    if (url.searchParams.has('error')) {
        return json({
            error: url.searchParams.get('error'),
            description: url.searchParams.get('error_description'),
        })
    }

    const code: string | null = url.searchParams.get('code');
    invariant(code, 'Missing authorization code');

    const tokens = await getOAuthTokenFromCode(code);

    return json(tokens);
}

export default function Login() {
    const data = useLoaderData();

    if (data.error) {
        return (
            <main>
                <pre>Error: {data.error}</pre>
                <pre>Description: {data.description}</pre>
            </main>
        );
    }

    return (
        <main>
            <pre>ID Token: {data.tokens.id_token}</pre>
            <pre>Access token: {data.tokens.access_token}</pre>
        </main>
    );
}