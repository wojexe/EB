import { Button } from "~/components/button";
import type { Route } from "./+types/home";
import { loader as rootLoader } from "~/root";
import { Form, redirect } from "react-router";
import { checkoutCart } from "lib/api";

export function meta(_: Route.MetaArgs) {
  return [{ title: "StoreFront Checkout" }];
}

export async function loader(args: Route.LoaderArgs) {
  const loaderData = await rootLoader(args);

  if (loaderData instanceof Response) {
    return loaderData;
  }

  if (loaderData.cart?.products?.length === 0) {
    return redirect(`/?cartId=${loaderData.cart?.id}`);
  }

  return loaderData;
}

export async function clientAction({ request }: Route.ClientActionArgs) {
  const formData = await request.formData();
  const cartID = parseInt((formData.get("cartID") as string) || "");
  const name = formData.get("name");
  const email = formData.get("email");
  const phone = formData.get("phone");
  const address = formData.get("address");
  const city = formData.get("city");
  const state = formData.get("state");
  const zipCode = formData.get("zipCode");

  const data = {
    name,
    email,
    phone,
    address,
    city,
    state,
    zipCode,
  };

  const response = await checkoutCart(cartID, data);

  return redirect("/");
}

export default function Home({ loaderData }: Route.ComponentProps) {
  const cartID = loaderData.cart.id;
  const products = loaderData.cart.products;

  const list = products.map((product) => (
    <div key={product.id} className="flex justify-between">
      <span>{product.name}</span>
      <span>${product.price.toFixed(2)}</span>
    </div>
  ));

  const sum = products.reduce((acc, product) => acc + product.price, 0);

  return (
    <div className="flex flex-col w-96 gap-3 place-self-center mt-6 proportional-nums">
      <h1 className="text-center text-lg font-bold">Checkout</h1>
      <div>{list}</div>
      <div className="flex justify-between font-bold">
        <span>Final total:</span>
        <span>${sum.toFixed(2)}</span>
      </div>

      <Form method="post">
        <input type="hidden" name="cartID" value={cartID} />
        <input type="text" name="name" placeholder="Name" required />
        <input type="email" name="email" placeholder="Email" required />
        <input type="tel" name="phone" placeholder="Phone" required />
        <input type="text" name="address" placeholder="Address" required />
        <input type="text" name="city" placeholder="City" required />
        <input type="text" name="zipCode" placeholder="Zip Code" required />
        <input type="text" name="state" placeholder="State" required />

        <Button className="mt-3">Confirm and pay</Button>
      </Form>
    </div>
  );
}
