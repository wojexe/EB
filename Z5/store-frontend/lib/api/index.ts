import * as v from "valibot";

export const API_URL = "http://localhost:1323";

const productSchema = v.object({
  ID: v.number(),
  Name: v.string(),
  Price: v.pipe(v.string()),
});

export type Product = v.InferOutput<typeof productSchema>;

const getProductsSchema = v.array(productSchema);

export async function getProducts() {
  const response = await fetch(`${API_URL}/products`);
  const json = await response.json();

  const validated = await v.parseAsync(getProductsSchema, json);

  return validated;
}
