import Link from "next/link";

export default function Header() {
  return (
    <div className="flex justify-between font-bold text-2xl underline mb-32">
      <div className="flex gap-5">
        <Link href="/">Home</Link>
        <Link href="/users">Users</Link>
      </div>
      <div>
        <Link href="/products">Products</Link>
      </div>
    </div>
  );
}

// Dynamic USER PAGE yapacagiz
