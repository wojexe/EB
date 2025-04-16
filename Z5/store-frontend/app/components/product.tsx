import type { Product } from "lib/api";
import { ShoppingCart } from "lucide-react";
import { Button } from "./button";

export function Product({
  product,
  purchasable = true,
}: {
  product: Product;
  purchasable?: boolean;
}) {
  return (
    <div
      className="flex flex-col p-3 gap-3 bg-neutral-800 rounded-md"
      data-testid="product"
    >
      <h2 className="font-bold text-center text-neutral-200 capitalize relative after:content-[''] after:absolute after:left-0 after:-bottom-3 after:w-full after:border-t after:border-dashed after:border-neutral-600">
        {product.name}
      </h2>

      {purchasable ? (
        <p className="flex flex-row mt-5 w-full justify-end">
          <Button
            type="submit"
            name="productId"
            value={product.id}
            data-testid="buy-button"
          >
            <ShoppingCart className="h-4 w-4 stroke-3" />
            <span className="proportional-nums">
              ${product.price.toFixed(2)}
            </span>
          </Button>
        </p>
      ) : (
        <p className="flex flex-row mt-5 w-full justify-end">
          <span className="text-neutral-200 proportional-nums">
            ${product.price.toFixed(2)}
          </span>
        </p>
      )}
    </div>
  );
}
