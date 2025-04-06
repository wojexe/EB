import type { Product } from "lib/api";

export function Product({ product }: { product: Product }) {
  return <div>{JSON.stringify(product)}</div>;
}
