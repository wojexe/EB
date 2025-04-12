import { Product } from "~/components/product";
import type { Route } from "./+types/home";
import { addProductToCart, getProducts } from "lib/api";
import { Form } from "react-router";

import { loader as rootLoader } from "../root";

export function meta(_: Route.MetaArgs) {
  return [{ title: "StoreFront Home" }];
}

export async function loader(args: Route.LoaderArgs) {
  const loaderData = await rootLoader(args);
  return loaderData;
}

export async function clientAction({ request }: Route.ClientActionArgs) {
  const formData = await request.formData();

  const cartId = formData.get("cartId");
  const productId = formData.get("productId");

  if (productId === null) {
    throw new Error("Invalid product ID");
  }

  const cart = await addProductToCart(
    parseInt(cartId as string),
    parseInt(productId as string)
  );

  return null;
}

export default function Home({ loaderData }: Route.ComponentProps) {
  const cartId = loaderData.cart?.id;
  const products = loaderData.products.map((product) => (
    <Product key={product.id} product={product} />
  ));

  return (
    <Form method="post" className="grid gap-6 m-6 grid-cols-4">
      <input type="hidden" name="cartId" value={cartId} />
      {products}
    </Form>
  );
}
