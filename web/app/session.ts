import { createCookieSessionStorage } from "@remix-run/node";
import invariant from "tiny-invariant";
import {SessionIdData} from "@/cognito/auth.session";

type SessionData = {
    accessToken: string;
    id: SessionIdData;
};

type SessionFlashData = {
    message: string;
};

function getCookieSecret(): string {
    const sessionSecret: string | undefined = process.env.SESSION_SECRET;
    invariant(sessionSecret, 'No session secret configured');

    return sessionSecret;
}

const { getSession, commitSession, destroySession } =
    createCookieSessionStorage<SessionData, SessionFlashData>(
        {
            cookie: {
                name: "__session",
                domain: "localhost",
                httpOnly: true,
                path: "/",
                sameSite: "lax",
                secrets: [getCookieSecret()],
                secure: true,
            },
        }
    );

export async function getSessionFromRequest(request: Request) {
    return await getSession(request.headers.get('Cookie'));
}

export { getSession, commitSession, destroySession };