import type { ActionArgs } from "@remix-run/node";
import { redirect } from "@remix-run/node";

import {destroySession, getSession} from "~/session";

export const action = async ({ request }: ActionArgs) => destroySession(await getSession(request.headers.get('Cookie')));

export const loader = async () => redirect("/");
