import type { Route } from "./+types/building";
import { Link } from "react-router";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "Building Page" },
    { name: "description", content: "Building component page" },
  ];
}

export default function Building() {
  return (
    <div>
      <nav style={{ padding: "1rem", borderBottom: "1px solid #ccc" }}>
        <Link to="/" style={{ marginRight: "1rem" }}>Home</Link>
        <Link to="/building">Building</Link>
      </nav>
      <div style={{ padding: "2rem" }}>
        <h1>Building Component</h1>
        <p>This is the Building component in the common routes.</p>
        <p>You have successfully navigated to the building page!</p>
      </div>
    </div>
  );
}