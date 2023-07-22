import { json, LoaderArgs } from "@remix-run/node";
import { requireUserSession } from "@/cognito/auth.session";

export async function loader({ request }: LoaderArgs) {
  const [accessToken, userSession] = await requireUserSession(request);

  return json({
    session: {
      payload: userSession,
    },
  });
}

export default function Ingress() {
  return (
    <div>
      <h1>Ingress</h1>
    </div>
  );
}
