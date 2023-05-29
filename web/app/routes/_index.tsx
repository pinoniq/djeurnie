import type { V2_MetaFunction } from "@remix-run/node";
import { Link } from "@remix-run/react";

export const meta: V2_MetaFunction = () => [{ title: "DJEURNIE Web" }];

export default function Index() {
  return (
    <main className="relative">
      <ul>
        <li><Link to="/login">Login</Link></li>
        <li><Link to="/login">Logout</Link></li>
      </ul>
    </main>
  );
}
