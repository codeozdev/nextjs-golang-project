import { getAllProducts } from "@/fetchData/get-product";
import DynamicTitle from "@/components/DynamicTitle";
import Link from "next/link";
import ProductsComponent from "@/components/products/products-component";
import { ProductsProps } from "@/types/ProductsType";

export default async function ProductsPage() {
  const { products, message } = await getAllProducts();
  console.log(message);
  return (
    <div>
      <DynamicTitle title="PRODUCTS PAGE" />
      <div className="flex justify-between px-2  text-xl font-bold">
        <Link href="/users-ekle" className="bg-blue-500 px-4">
          EKLE
        </Link>
      </div>
      {products.map((product: ProductsProps) => (
        <ProductsComponent key={product.id} product={product} />
      ))}
    </div>
  );
}
