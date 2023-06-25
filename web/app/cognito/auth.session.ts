import { Session } from "@remix-run/server-runtime";
import {getCognitoTokenVerifier, getOAuthTokenFromCode} from "@/cognito/auth";
import invariant from "tiny-invariant";
import { getSessionFromRequest } from "@/session";
import { redirect } from "@remix-run/node";

export type SessionIdData = {
    email: string,
};

export async function setAccessToken(session: Session<any, any>, accessToken: string) {
    const verifier = getCognitoTokenVerifier('access');

    await verifier.verify(accessToken);

    session.set('accessToken', accessToken);
}

export function getAccessToken(session: Session<any, any>): string {
    return session.get('accessToken');
}

export function getId(session: Session<any, any>): SessionIdData {
    return session.get('id');
}

export async function getAccessTokenPayload(session: Session<any, any>) {
    const accessToken = getAccessToken(session);
    const verifier = getCognitoTokenVerifier('access');

    return await verifier.verify(accessToken);
}

export async function setUserSessionFromCode(session: Session<any, any>, code: string) {
    const tokens = await getOAuthTokenFromCode(code);
    invariant(tokens.access_token, 'Missing access token');

    await setAccessToken(session, tokens.access_token);

    invariant(tokens.id_token, 'Missing id token');
    const verifier = getCognitoTokenVerifier('id');
    const payload = await verifier.verify(tokens.id_token);

    session.set('id', {
        email: payload.email,
    });
}

export type UserSession = {
    email: string,
};

export function getUserSession(session: Session<any, any>): UserSession {
    return session.get('id');
}

export async function requireUserSession(request: Request) : Promise<[string, UserSession]> {
    const session = await getSessionFromRequest(request);
    const [accessToken, userSession] = await Promise.all([
        getAccessToken(session),
        getUserSession(session)
    ]);

    if (!accessToken || !userSession) {
        throw redirect('/login');
    }

    return [accessToken, userSession];
}