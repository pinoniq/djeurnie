import { json, LoaderArgs } from "@remix-run/node";
import { requireUserSession } from "@/cognito/auth.session";
import api from "@/api";
import { useLoaderData } from "@remix-run/react";
import { sleep } from "@/lib/utils";
import Typography from "@/components/ui/Typography";
import { Card, CardHeader, CardTitle } from "@/components/ui/card";

export async function loader({ request }: LoaderArgs) {
  const [accessToken, userSession] = await requireUserSession(request);

  const ingressList = await api.ingress.list(accessToken);

  return json({
    ingressList,
  });
}

export default function Ingress() {
  const data = useLoaderData<typeof loader>();

  return (
    <div>
      <Typography.headings.Heading1>Ingress</Typography.headings.Heading1>
      <div className="flew flex-col">
        {data.ingressList.items.map((ingress) => (
          <Card key={ingress.id}>
            <CardHeader>
              <CardTitle>{ingress.displayName}</CardTitle>
            </CardHeader>
          </Card>
        ))}
      </div>
    </div>
  );
}
