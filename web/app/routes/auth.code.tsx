import { LoaderArgs, json } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";

export async function loader({ request }: LoaderArgs) {
    const url = new URL(request.url);

    const {
        COGNITO_DOMAIN,
        COGNITO_CLIENT_ID,
        COGNITO_CLIENT_SECRET,
        COGNITO_REDIRECT_URI,
        COGNITO_SCOPES,
    } = process.env;

    const tokenHeaders = new Headers();
    tokenHeaders.append('Content-Type', 'application/x-www-form-urlencoded');

    const tokenBody = new URLSearchParams();
    tokenBody.append("grant_type", "authorization_code");
    tokenBody.append("client_id", COGNITO_CLIENT_ID);
    tokenBody.append("client_secret", COGNITO_CLIENT_SECRET);
    tokenBody.append("redirect_uri", COGNITO_REDIRECT_URI);
    tokenBody.append("code", url.searchParams.get('code'));
    tokenBody.append("scope", COGNITO_SCOPES);

    const tokenResponse = await fetch(`${COGNITO_DOMAIN}/oauth2/token`, {
        method: 'POST',
        headers: tokenHeaders,
        body: tokenBody,
    });

    const token = await tokenResponse.json();

    return json({
        code: url.searchParams.get('code'),
        token,
        status: tokenResponse.status,
        tokenBody: tokenBody.toString(),
    });
}

export default function Login() {
    const data = useLoaderData();

    return (
        <main>
            <pre>{data.code}</pre>
            <pre>Token: {JSON.stringify(data.token)}</pre>
            <pre>Status: {data.status}</pre>
            <pre>{data.tokenBody}</pre>
        </main>
    );
}