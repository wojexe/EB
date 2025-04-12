import { type RouteConfig, index, route } from "@react-router/dev/routes";

export default [
  index("routes/home.tsx"),
  route("/cart", "routes/cart.tsx"),
  route("/checkout", "routes/checkout.tsx"),
  //
] satisfies RouteConfig;
