import type { ActionArgs } from "@remix-run/node";
import { redirect } from "@remix-run/node";

import {destroySession, getSessionFromRequest} from "@/session";

export const action = async ({ request }: ActionArgs) => destroySession(await getSessionFromRequest(request));

export const loader = async () => redirect("/");
