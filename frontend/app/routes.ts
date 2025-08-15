import { type RouteConfig, index, route } from "@react-router/dev/routes";

export default [
  index("routes/home.tsx"),
  route("/building", "routes/common/building.tsx"),
  // Catch-all route for unmatched paths
  route("*", "routes/not-found.tsx")
] satisfies RouteConfig;
