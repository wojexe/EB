import type { Route } from "./+types/home";
import { loader as rootLoader } from "../root";
import { Product } from "~/components/product";
import { Button } from "~/components/button";
import { Link } from "react-router";

export function meta(_: Route.MetaArgs) {
  return [{ title: "StoreFront Cart" }];
}

export async function loader(args: Route.LoaderArgs) {
  const loaderData = await rootLoader(args);
  return loaderData;
}

export default function Home({ loaderData }: Route.ComponentProps) {
  const products = loaderData.cart?.products?.map((product) => (
    <Product key={product.id} product={product} purchasable={false} />
  ));

  return (
    <form className="flex flex-col m-6 gap-6">
      <div className="grid gap-6 grid-cols-4">{products}</div>

      <Link
        className="contents"
        to={{ pathname: "/checkout", search: `?cartId=${loaderData.cart?.id}` }}
      >
        <Button type="submit">Checkout</Button>
      </Link>
    </form>
  );
}
