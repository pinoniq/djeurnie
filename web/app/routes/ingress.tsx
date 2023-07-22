import { json, LoaderArgs } from "@remix-run/node";
import { requireUserSession } from "@/cognito/auth.session";
import api from "@/api";
import { useLoaderData } from "@remix-run/react";
import { sleep } from "@/lib/utils";

export async function loader({ request }: LoaderArgs) {
  const [accessToken, userSession] = await requireUserSession(request);

  await sleep(1000);

  const ingressList = await api.ingress.list(accessToken);

  return json({
    ingressList,
  });
}

export default function Ingress() {
  const data = useLoaderData<typeof loader>();

  return (
    <div>
      <h1>Ingress</h1>
      <ul>
        {data.ingressList.items.map((ingress) => (
          <li key={ingress.id}>{ingress.displayName}</li>
        ))}
      </ul>
    </div>
  );
}
