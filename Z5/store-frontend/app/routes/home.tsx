import type { Route } from "./+types/home";
import { getProducts } from "lib/api";

export function meta(_: Route.MetaArgs) {
  return [
    { title: "New React Router App" },
    { name: "description", content: "Welcome to React Router!" },
  ];
}

export async function loader(_: Route.LoaderArgs) {
  const products = await getProducts();

  return products;
}

export async function clientLoader(_: Route.ClientLoaderArgs) {
  const products = await getProducts();

  return products;
}

export default function Home({ loaderData }: Route.ComponentProps) {
  const products = loaderData.map((product) => {
    console.log(product);
  });
}
