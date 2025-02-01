import Link from "next/link";

export default function Header() {
  return (
    <div className="flex gap-5 font-bold text-2xl underline mb-32">
      <Link href="/">Home</Link>
      <Link href="/users">Users</Link>
    </div>
  );
}

// Dynamic USER PAGE yapacagiz
