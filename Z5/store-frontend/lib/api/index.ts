import * as v from "valibot";

export const API_URL = () => process.env.API_URL ?? "http://localhost:1323";

const productSchema = v.object({
  id: v.number(),
  name: v.string(),
  price: v.pipe(
    v.string(),
    v.transform((value) => parseFloat(value))
  ),
});

export type Product = v.InferOutput<typeof productSchema>;

const getProductsResponseSchema = v.array(productSchema);

export async function getProducts() {
  const response = await fetch(`${API_URL()}/products`);
  const json = await response.json();

  const validated = await v.parseAsync(getProductsResponseSchema, json);
  return validated;
}

const cartResponseSchema = v.object({
  id: v.number(),
  products: v.pipe(
    v.nullable(v.array(productSchema)),
    v.transform((value) => value || [])
  ),
});

export type Cart = v.InferOutput<typeof cartResponseSchema>;

export async function createCart() {
  const response = await fetch(`${API_URL()}/carts`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
  });
  const json = await response.json();

  const validated = await v.parseAsync(cartResponseSchema, json);
  return validated;
}

export async function getCart(cartId: number) {
  const response = await fetch(`${API_URL()}/carts/${cartId}`);
  const json = await response.json();

  const validated = await v.parseAsync(cartResponseSchema, json);
  return validated;
}

export async function addProductToCart(cartId: number, productId: number) {
  const response = await fetch(
    `${API_URL()}/carts/${cartId}/products/${productId}`,
    {
      method: "POST",
    }
  );
  const json = await response.json();

  const validated = await v.parseAsync(cartResponseSchema, json);
  return validated;
}

export async function deleteCart(cartId: number) {
  const response = await fetch(`${API_URL()}/carts/${cartId}`, {
    method: "DELETE",
  });
  const json = await response.json();

  const validated = await v.parseAsync(cartResponseSchema, json);
  return validated;
}

export async function checkoutCart(cartId: number, data: any) {
  const response = await fetch(`${API_URL()}/carts/${cartId}/checkout`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });

  return response;
}
