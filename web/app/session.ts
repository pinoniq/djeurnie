import { createCookieSessionStorage } from "@remix-run/node";
import invariant from "tiny-invariant";

type SessionData = {
    accessToken: string;
    idToken: string;
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
                maxAge: 3600,
                path: "/",
                sameSite: "lax",
                secrets: [getCookieSecret()],
                secure: true,
            },
        }
    );

export { getSession, commitSession, destroySession };