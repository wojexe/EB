import { ShoppingCart, Store } from "lucide-react";
import { Button } from "./button";
import { Link } from "react-router";
import type { Cart } from "lib/api";

export function NavigationBar({ cart }: { cart?: Cart }) {
  return (
    <nav className="sticky top-0 z-50 bg-neutral-800 border-b-2 border-neutral-700">
      <div className="flex justify-between items-center p-4 text-white">
        <Link
          to={{ pathname: "/", search: `?cartId=${cart?.id}` }}
          className="flex items-center"
        >
          <Store className="h-6 w-6 mr-2" />
          <span className="font-bold text-xl">StoreFront</span>
        </Link>

        <Link to={{ pathname: "/cart", search: `?cartId=${cart?.id}` }}>
          <Button>
            <ShoppingCart />
          </Button>
        </Link>
      </div>
    </nav>
  );
}
