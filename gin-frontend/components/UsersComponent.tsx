import { User } from "@/types/UserType";
import Link from "next/link";

export default function UsersComponent({ user }: { user: User }) {
  return (
    <Link
      href={`/app/(users)/users/${user.id}`}
      className="flex gap-2 justify-between font-bold text-3xl"
    >
      <h1>{user.name}</h1>
      <p>{user.email}</p>
    </Link>
  );
}
