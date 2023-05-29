import { LoaderArgs, json } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import {getOAuthAuthorizationUrl} from "~/cognito/auth";

export function loader({ request }: LoaderArgs) {
    return json({
        cognito: getOAuthAuthorizationUrl(),
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