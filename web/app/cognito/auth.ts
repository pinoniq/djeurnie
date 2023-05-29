import invariant from "tiny-invariant";
import {query} from "express";

const {
    COGNITO_DOMAIN,
    COGNITO_CLIENT_ID,
    COGNITO_CLIENT_SECRET,
    COGNITO_SCOPES,
    COGNITO_REDIRECT_URI,
    COGNITO_RESPONSE_TYPE,
} = process.env;

function getCognitoDomain(): string {
    const cognitoDomain: string | undefined = process.env.COGNITO_DOMAIN;
    invariant(cognitoDomain, 'No Cognito Domain configured');

    return cognitoDomain;
}

function getCognitoClientId(): string {
    const cognitoClientId: string | undefined = process.env.COGNITO_CLIENT_ID;
    invariant(cognitoClientId, 'No Cognito Client ID configured');

    return cognitoClientId;
}

function getCognitoClientSecret(): string {
    const cognitoClientSecret: string | undefined = process.env.COGNITO_CLIENT_SECRET;
    invariant(cognitoClientSecret, 'No Cognito Client Secret configured');

    return cognitoClientSecret;
}

function getCognitoClientScopes(): string {
    const cognitoClientScopes: string | undefined = process.env.COGNITO_SCOPES;
    invariant(cognitoClientScopes, 'No Cognito Client Scopes configured');

    return cognitoClientScopes;
}

function getCognitoRedirectUri(): string {
    const cognitoRedirectUri: string | undefined = process.env.COGNITO_REDIRECT_URI;
    invariant(cognitoRedirectUri, 'No Cognito Client RedirectUri configured');

    return cognitoRedirectUri;
}

function getCognitoResponseType(): string {
    const cognitoResponseType: string | undefined = process.env.COGNITO_RESPONSE_TYPE;
    invariant(cognitoResponseType, 'No Cognito Response Type configured');

    return cognitoResponseType;
}

export function getOAuthAuthorizationUrl(): string {
    const queryParams = new URLSearchParams();
    queryParams.append('client_id', getCognitoClientId());
    queryParams.append('response_type', getCognitoResponseType());
    queryParams.append('redirect_uri', getCognitoRedirectUri());

    // add scopes directly into the url since Cognito required the + to be un-encoded.
    return `${COGNITO_DOMAIN}/oauth2/authorize?${queryParams.toString()}&scope=${getCognitoClientScopes()}`;
}

export async function getOAuthTokenFromCode(code: string) {
    const tokenHeaders = new Headers();
    tokenHeaders.append('Content-Type', 'application/x-www-form-urlencoded');

    const tokenBody = new URLSearchParams();
    tokenBody.append('code', code);
    tokenBody.append('grant_type', 'authorization_code');
    tokenBody.append('client_id', getCognitoClientId());
    tokenBody.append('client_secret', getCognitoClientSecret());
    tokenBody.append('redirect_uri', getCognitoRedirectUri());
    tokenBody.append('scope', getCognitoClientScopes());

    const tokenResponse = await fetch(`${COGNITO_DOMAIN}/oauth2/token`, {
        method: 'POST',
        headers: tokenHeaders,
        body: tokenBody,
    });

    return await tokenResponse.json();
}