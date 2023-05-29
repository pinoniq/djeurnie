import { LoaderArgs, json } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";

export function loader({ request }: LoaderArgs) {
    const {
        COGNITO_DOMAIN,
        COGNITO_CLIENT_ID,
        COGNITO_SCOPES,
        COGNITO_REDIRECT_URI,
        COGNITO_RESPONSE_TYPE,
    } = process.env;

    return json({
        cognito: `${COGNITO_DOMAIN}/oauth2/authorize?client_id=${COGNITO_CLIENT_ID}&response_type=${COGNITO_RESPONSE_TYPE}&scope=${COGNITO_SCOPES}&redirect_uri=${encodeURIComponent(COGNITO_REDIRECT_URI)}`,
    });
}

export default function Login() {
    const data = useLoaderData();

    return (
        <main>
            <pre>{data.cognito}</pre> <br />
            <a href={data.cognito}>Start login</a>
        </main>
    );
}