import { asJson } from "@/lib/utils";

type Ingress = {
  id: string;
  displayName: string;
};

type IngressList = {
  tenant: string;
  items: Ingress[];
};

export async function list(accessToken: string): Promise<IngressList> {
  return fetch("http://127.0.0.1:3001/ingress.json", {
    headers: {
      Authorization: `Bearer ${accessToken}`,
    },
  }).then(asJson);
}
