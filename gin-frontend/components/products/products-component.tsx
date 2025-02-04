import { ProductsProps } from "@/types/ProductsType";
import Link from "next/link";

export default function ProductsComponent({
  product,
}: {
  product: ProductsProps;
}) {
  // sadece tarih kismini formatlamak icin
  // const formattedDate = new Date(product.created_at).toLocaleDateString(
  //   "tr-TR",
  //   {
  //     day: "2-digit",
  //     month: "2-digit",
  //     year: "numeric",
  //   },
  // );

  return (
    <Link href={`/products/${product.id}`} className="w-full">
        <div className="flex justify-between px-2 py-2 border-b border-gray-300">
            <h1>{product.name}</h1>
            <p>{product.price}$</p>
        </div>
    </Link>
  );
}
