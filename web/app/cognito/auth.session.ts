import { Session } from "@remix-run/server-runtime";
import { getCognitoTokenVerifier } from "~/cognito/auth";

export async function setAccessToken(session: Session<any, any>, accessToken: string) {
    const verifier = getCognitoTokenVerifier('access');

    await verifier.verify(accessToken);

    session.set('accessToken', accessToken);
}

export function getAccessToken(session: Session<any, any>) {
    return session.get('accessToken');
}

export async function getAccessTokenPayload(session: Session<any, any>) {
    const accessToken = getAccessToken(session);
    const verifier = getCognitoTokenVerifier('access');

    return await verifier.verify(accessToken);
}

export async function setIdToken(session: Session<any, any>, accessToken: string) {
    const verifier = getCognitoTokenVerifier('id');

    await verifier.verify(accessToken);

    session.set('idToken', accessToken);
}

export function getIdToken(session: Session<any, any>) {
    return session.get('accessToken');
}